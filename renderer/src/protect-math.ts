export type MathToken={id:string;source:string;display:boolean};
type XmlNode={name:string;attributes:Record<string,string>;children:(XmlNode|string|object)[]};
/** Only a '$' in a cmark text node may open math. Code, HTML, and link
 * destinations remain opaque. cmark source columns count UTF-8 bytes. */
export function protectMath(source:string,astXML:string){
 const tree=Bun.XML.parse(astXML,{compact:false}) as XmlNode;
 const lines=source.split('\n'),starts:number[]=[];let at=0;
 for(const line of lines){starts.push(at);at+=line.length+1}
 const offset=(line:number,column:number)=>{
  const value=lines[line-1]??'';return (starts[line-1]??source.length)+Buffer.from(value).subarray(0,Math.max(0,column-1)).toString('utf8').length;
 };
 const text=new Uint8Array(source.length);
 function visit(n:any){
  if(!n||typeof n!=='object')return;
  const pos=n.attributes?.sourcepos?.match(/^(\d+):(\d+)-(\d+):(\d+)$/);
  if(n.name==='text'&&pos){const [l,c,el,ec]=pos.slice(1).map(Number);if(l&&c&&el&&ec)text.fill(1,offset(l,c),offset(el,ec+1))}
  for(const ch of n.children??[])visit(ch);
 }
 visit(tree);
 const escaped=(i:number)=>{let n=0;while(i>0&&source[--i]==='\\')n++;return n%2===1};
 let prefix='GFMMATH'+new Bun.CryptoHasher('sha256').update(source).digest('hex').slice(0,16).toUpperCase();
 while(source.includes(prefix))prefix+='X';
 const tokens:MathToken[]=[];let output='',last=0;
 for(let i=0;i<source.length;i++){
  if(!text[i]||source[i]!=='$'||escaped(i))continue;
  const protectedInline=source.startsWith('$`',i),display=!protectedInline&&source.startsWith('$$',i);
  const open=protectedInline?'$`':display?'$$':'$',close=protectedInline?'`$':open;
  if(!protectedInline&&!display&&/\s/.test(source[i+1]??' '))continue;
  let end=i+open.length;
  for(;;){
   end=source.indexOf(close,end);
   if(end<0)break;
   if(!escaped(end)&&(protectedInline||text[end])&&(display||protectedInline||(!/\s/.test(source[end-1]??' ')&&!/\d/.test(source[end+1]??''))))break;
   end+=close.length;
  }
  if(end<0||end===i+open.length)continue;
  const body=source.slice(i+open.length,end);
  if(body.length>20000||(!display&&body.includes('\n\n')))continue;
  const id=prefix+'N'+tokens.length;tokens.push({id,source:body,display});
  output+=source.slice(last,i)+'`'+id+'`';i=end+close.length-1;last=i+1;
 }
 return {source:output+source.slice(last),tokens};
}
