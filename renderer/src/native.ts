import {cc,CString} from 'bun:ffi';
import {join} from 'node:path';
const root=join(import.meta.dir,'..');
const lib=cc({
 source:join(root,'src/cmark-unity.c'),
 include:['src','extensions','generated'].map(dir=>join(root,'vendor/cmark-gfm',dir)),
 symbols:{
 cmark_version_string:{args:[],returns:'cstring'},
 cmark_gfm_core_extensions_ensure_registered:{args:[],returns:'void'},
 cmark_find_syntax_extension:{args:['buffer'],returns:'ptr'},
 cmark_parser_new:{args:['i32'],returns:'ptr'},
 cmark_parser_attach_syntax_extension:{args:['ptr','ptr'],returns:'i32'},
 cmark_parser_feed:{args:['ptr','buffer','u64'],returns:'void'},
 cmark_parser_finish:{args:['ptr'],returns:'ptr'},
 cmark_parser_get_syntax_extensions:{args:['ptr'],returns:'ptr'},
 cmark_render_html:{args:['ptr','i32','ptr'],returns:'ptr'},
 cmark_render_xml:{args:['ptr','i32'],returns:'ptr'},
 cmark_node_first_child:{args:['ptr'],returns:'ptr'},
 cmark_node_next:{args:['ptr'],returns:'ptr'},
 cmark_node_get_type_string:{args:['ptr'],returns:'cstring'},
 cmark_node_get_literal:{args:['ptr'],returns:'cstring'},
 cmark_node_set_literal:{args:['ptr','buffer'],returns:'i32'},
 cmark_node_set_syntax_extension:{args:['ptr','ptr'],returns:'i32'},
 cmark_node_get_start_line:{args:['ptr'],returns:'i32'},
 cmark_node_get_start_column:{args:['ptr'],returns:'i32'},
 cmark_node_get_end_line:{args:['ptr'],returns:'i32'},
 cmark_node_get_end_column:{args:['ptr'],returns:'i32'},
 cmark_gfm_extensions_set_tasklist_item_checked:{args:['ptr','i32'],returns:'i32'},
 cmark_node_free:{args:['ptr'],returns:'void'},
 cmark_parser_free:{args:['ptr'],returns:'void'},
 gfm_free:{args:['ptr'],returns:'void'},
 },
});
const s=lib.symbols;
s.cmark_gfm_core_extensions_ensure_registered();
export const nativeVersion=String(s.cmark_version_string());
// cmark's source positions are 1-based UTF-8 byte columns.
function lineOffsets(source:string){
 const lines=source.split('\n'),starts:number[]=[];let offset=0;
 for(const line of lines){starts.push(offset);offset+=line.length+1}
 return {lines,offset(line:number,column:number){
  const text=lines[line-1]??'';
  return (starts[line-1]??source.length)+Buffer.from(text).subarray(0,Math.max(0,column-1)).toString('utf8').length;
 }};
}
const typeOf=(node:any)=>String(s.cmark_node_get_type_string(node));
function children(node:any){const result:any[]=[];for(let child=s.cmark_node_first_child(node);child;child=s.cmark_node_next(child))result.push(child);return result}
function walk(root:any,visit:(node:any)=>void){const pending=[root];while(pending.length){const node=pending.pop()!;visit(node);for(const child of children(node))pending.push(child)}}
function location(node:any,positions:ReturnType<typeof lineOffsets>){
 const a=s.cmark_node_get_start_line(node),b=s.cmark_node_get_start_column(node),c=s.cmark_node_get_end_line(node),d=s.cmark_node_get_end_column(node);
 return a>0&&b>0&&c>0&&d>0?[positions.offset(a,b),positions.offset(c,d+1)] as const:undefined;
}
// Find unresolved footnote labels containing inline images. Markdown image
// descriptions and destinations may contain balanced brackets/parentheses.
// This only identifies candidates; AST context and emitted text are checked
// before any source is rewritten.
function imageFootnoteCandidates(text:string){
 const result:{start:number;end:number;label:string;raw:string}[]=[];
 const escaped=(i:number)=>{let count=0;while(i>0&&text[--i]==='\\')count++;return count%2===1};
 for(let start=0;start<text.length-3;start++){
  if(!text.startsWith('[^',start)||escaped(start))continue;
  let image=start+2;
  while(image<text.length && text[image]!=='\n' && (text[image]!==']'||escaped(image)) && !text.startsWith('![',image)){
   // A nested opener cannot belong to this plain label; inspect it as the
   // next candidate instead of repeatedly rescanning the remaining input.
   if(text.startsWith('[^',image) && !escaped(image))break;
   image++;
  }
  if(!text.startsWith('![',image)||escaped(image))continue;
  const label=text.slice(start+2,image);
  let pos=image+2,depth=1;
  while(pos<text.length && text[pos]!=='\n' && depth){
   if(text[pos]==='\\' && pos+1<text.length && text[pos+1]!=='\n'){pos+=2;continue}
   if(text[pos]==='[')depth++;else if(text[pos]===']')depth--;
   pos++;
  }
  if(depth||text[pos]!=='(')continue;
  pos++;depth=1;let angle=false,quote='';
  while(pos<text.length && text[pos]!=='\n' && depth){
   const ch=text[pos];
   if(ch==='\\' && pos+1<text.length && text[pos+1]!=='\n'){pos+=2;continue}
   {
    if(quote){if(ch===quote)quote=''}
    else if(angle){if(ch==='>')angle=false}
    else if(ch==='<')angle=true;
    else if((ch==='"'||ch==="'") && /\s/.test(text[pos-1]??''))quote=ch;
    else if(ch==='(')depth++;
    else if(ch===')')depth--;
   }
   pos++;
  }
  if(depth||angle||quote||text[pos]!==']')continue;
  const end=pos+1;result.push({start,end,label,raw:text.slice(start,end)});start=end-1;
 }
 return result;
}
// The release loses the label of a multiline unresolved footnote. Only rewrite
// paragraphs where the parser actually produced that unresolved empty label;
// code, links, and raw HTML remain opaque. The original caller input is intact.
function recoverMultilineLabels(doc:any,source:string){
 const positions=lineOffsets(source),edits:{start:number;end:number;replacement:string}[]=[];
 walk(doc,node=>{
  if(typeOf(node)!=='paragraph')return;
  const range=location(node,positions);if(!range)return;
  let textContent='';const protectedRanges:(readonly [number,number])[]=[];
  walk(node,child=>{
   const type=typeOf(child);
   if(type==='text')textContent+=String(s.cmark_node_get_literal(child));
   if(['code','code_block','html_inline','html_block','link','image'].includes(type)){
    const r=location(child,positions);if(r)protectedRanges.push(r);
   }
  });
  const text=source.slice(range[0],range[1]);
  if(textContent.includes('[^]'))for(const m of text.matchAll(/\[\^([^\]\n]+)\r?\n[^\]\n]*\]/g)){
   const start=range[0]+m.index,end=start+m[0].length;
   if(protectedRanges.some(([a,b])=>start<b&&end>a))continue;
   // An apparent multiline definition is parsed as literal text upstream;
   // keep its colon literal rather than accidentally creating a definition.
   const definitionColon=source[end]===':';
   edits.push({start,end:end+(definitionColon?1:0),replacement:`[^${m[1]}]${definitionColon?'\\:':''}`});
  }
  // An image inside an unresolved reference can hide an otherwise valid label.
  // Rewrite only markup the unmodified parser actually emitted as text.
  for(const candidate of imageFootnoteCandidates(text)){
   if(!textContent.includes(candidate.raw))continue;
   const start=range[0]+candidate.start,end=range[0]+candidate.end;
   if(protectedRanges.some(([a,b])=>start<b&&end>a))continue;
   edits.push({start,end,replacement:`[^${candidate.label}]`});
  }
 });
 edits.sort((a,b)=>b.start-a.start);
 for(const edit of edits)source=source.slice(0,edit.start)+edit.replacement+source.slice(edit.end);
 return source;
}
function adaptNodes(doc:any,source:string){
 const positions=lineOffsets(source);
 walk(doc,node=>{
  const type=typeOf(node);
  if(type==='tasklist'){
   const line=s.cmark_node_get_start_line(node),column=s.cmark_node_get_start_column(node);
   if(line<1||column<1)return;
   const start=positions.offset(line,column);
   const marker=source.slice(start,source.indexOf('\n',start)<0?source.length:source.indexOf('\n',start)).match(/^(?: {0,3})(?:[-+*]|\d{1,9}[.)])[ \t]+\[([ xX])\]/);
   if(marker)s.cmark_gfm_extensions_set_tasklist_item_checked(node,marker[1].toLowerCase()==='x'?1:0);
  }else if(type==='item'){
   // The stock task-list extension does not always recognize nested list
   // containers. Promote only an actual list marker followed by a literal
   // task marker in the first paragraph, using the public extension API.
   const line=s.cmark_node_get_start_line(node),column=s.cmark_node_get_start_column(node);
   if(line<1||column<1)return;
   const start=positions.offset(line,column),lineEnd=source.indexOf('\n',start);
   const marker=source.slice(start,lineEnd<0?source.length:lineEnd).match(/^(?: {0,3})(?:[-+*]|\d{1,9}[.)])[ \t]+/);
   if(!marker)return;
   const paragraph=s.cmark_node_first_child(node);if(!paragraph||typeOf(paragraph)!=='paragraph')return;
   const first=s.cmark_node_first_child(paragraph);if(!first||typeOf(first)!=='text')return;
   const literal=String(s.cmark_node_get_literal(first));
   const prefix=literal.match(/^\[([ xX])\][ \t]/);if(!prefix)return;
   const ext=s.cmark_find_syntax_extension(Buffer.from('tasklist\0'));
   if(!ext||!s.cmark_node_set_syntax_extension(node,ext))throw Error('Cannot mark nested task item');
   s.cmark_gfm_extensions_set_tasklist_item_checked(node,prefix[1].toLowerCase()==='x'?1:0);
   if(!s.cmark_node_set_literal(first,Buffer.from(literal.slice(prefix[0].length)+'\0')))throw Error('Cannot update task text');
  }else if(type==='text'){
   const original=String(s.cmark_node_get_literal(node));
   // Unresolved references are emitted as text by cmark. Mirror GitHub's
   // handling of escaped punctuation and embedded images without touching code.
   let value=original.replace(/\[\^([^\]\n]*?)!\[([^\]\n]*)\]\((?:\\.|[^)\n])*\)\]/g,(_m,label)=>`[^${label}]`);
   value=value.replace(/\[\^(?:\\.|[^\]\\\n])*\]/g,ref=>ref.replace(/\\([!"#$%&'()*+,\-./:;<=>?@[\]\\^_`{|}~])/g,'$1'));
   if(value!==original)s.cmark_node_set_literal(node,Buffer.from(value+'\0'));
  }
 });
}
// Explicit byte length handles embedded NUL and non-ASCII input correctly.
export function nativeRender(source:string):{html:string;astXML:string}{
 const options=(1<<17)|(1<<13)|(1<<11)|(1<<2); // Raw HTML is sanitized later.
 const parserFor=(input:string)=>{
  const parser=s.cmark_parser_new(options);if(!parser)throw Error('cmark allocation failed');
  try{
   for(const name of ['table','autolink','tagfilter','strikethrough','tasklist']){
    const ext=s.cmark_find_syntax_extension(Buffer.from(name+'\0'));if(!ext)throw Error(`Missing cmark extension: ${name}`);
    if(!s.cmark_parser_attach_syntax_extension(parser,ext))throw Error(`Cannot attach extension: ${name}`);
   }
   const bytes=Buffer.from(input,'utf8');s.cmark_parser_feed(parser,bytes,bytes.byteLength);
   const doc=s.cmark_parser_finish(parser);if(!doc)throw Error('cmark parse failed');
   return {parser,doc};
  }catch(error){s.cmark_parser_free(parser);throw error}
 };
 const owned=(p:any)=>{if(!p)throw Error('cmark render allocation failed');try{return String(new CString(p))}finally{s.gfm_free(p)}};
 let {parser,doc}=parserFor(source);
 try{
  const recovered=recoverMultilineLabels(doc,source);
  if(recovered!==source){s.cmark_node_free(doc);s.cmark_parser_free(parser);doc=null;parser=null;({parser,doc}=parserFor(recovered));source=recovered}
  adaptNodes(doc,source);
  const html=owned(s.cmark_render_html(doc,options,s.cmark_parser_get_syntax_extensions(parser)));
  // Upstream emits invalid XML element names for the two footnote node types.
  // Repair its serialized XML without modifying the upstream parser source.
  const astXML=owned(s.cmark_render_xml(doc,1<<1))
   .replace(/<<unknown>([^>]*?)\s*\/>/g,'<footnote_reference$1 />')
   .replace(/<<unknown>/g,'<footnote_definition')
   .replace(/<\/<unknown>>/g,'</footnote_definition>');
  return {html,astXML};
 }finally{if(doc)s.cmark_node_free(doc);if(parser)s.cmark_parser_free(parser)}
}
