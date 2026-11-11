import {dlopen,CString} from 'bun:ffi';
import {join} from 'node:path';
const root=join(import.meta.dir,'..');
const nativePath=join(root,'native',`libcmark-gfm-${process.platform}-${process.arch}.${process.platform==='darwin'?'dylib':'so'}`);
const lib=dlopen(nativePath,{
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
 cmark_node_free:{args:['ptr'],returns:'void'},
 cmark_parser_free:{args:['ptr'],returns:'void'},
 gfm_free:{args:['ptr'],returns:'void'},
});
const s=lib.symbols;
s.cmark_gfm_core_extensions_ensure_registered();
export const nativeVersion=String(s.cmark_version_string());
// Explicit byte length handles embedded NUL and non-ASCII input correctly.
export function nativeRender(source:string):{html:string;astXML:string}{
 const bytes=Buffer.from(source,'utf8');
 const options=(1<<17)|(1<<13)|(1<<11)|(1<<2); // raw HTML is sanitized after HTML5 parsing; footnotes, github_pre_lang, hardbreaks
 const parser=s.cmark_parser_new(options);if(!parser)throw Error('cmark allocation failed');
 let doc:any=null;
 const owned=(p:any)=>{if(!p)throw Error('cmark render allocation failed');try{return String(new CString(p))}finally{s.gfm_free(p)}};
 try{
  for(const name of ['table','autolink','tagfilter','strikethrough','tasklist']){
   const e=s.cmark_find_syntax_extension(Buffer.from(name+'\0'));if(!e)throw Error(`Missing cmark extension: ${name}`);
   if(!s.cmark_parser_attach_syntax_extension(parser,e))throw Error(`Cannot attach extension: ${name}`);
  }
  s.cmark_parser_feed(parser,bytes,bytes.byteLength);doc=s.cmark_parser_finish(parser);
  if(!doc)throw Error('cmark parse failed');
  return {html:owned(s.cmark_render_html(doc,options,s.cmark_parser_get_syntax_extensions(parser))),astXML:owned(s.cmark_render_xml(doc,1<<1))};
 }finally{if(doc)s.cmark_node_free(doc);s.cmark_parser_free(parser)}
}
