import {createDOM,installDOMGlobals} from './dom';
import {nativeRender} from './native';
import {protectMath} from './protect-math';
import {createMathRenderer} from './math';
// DOMPurify and Mermaid retain DOM references after their first import.
// Reuse one DOM and serialize complete renders in index.ts.
let sharedDOM:ReturnType<typeof createDOM>|undefined;
export async function renderInProcess(source:string):Promise<string>{
const original=nativeRender(source);
const prepared=protectMath(source,original.astXML);
const html=prepared.tokens.length?nativeRender(prepared.source).html:original.html;
const window=sharedDOM??=createDOM();
const restore=installDOMGlobals(window);
const document=window.document;
let math:ReturnType<typeof createMathRenderer>|null=null;
try{
 document.body.innerHTML=html;
 // Decorative product classes, not an assertion of GitHub byte equivalence.
 for(const code of document.querySelectorAll('code'))code.classList.add('notranslate');
 for(const table of [...document.querySelectorAll('table')]){
  table.setAttribute('role','table');const wrapper=document.createElement('markdown-accessiblity-table');table.replaceWith(wrapper);wrapper.append(table);
 }
 for(const input of document.querySelectorAll('input[type="checkbox"]')){
  input.closest('li')?.classList.add('task-list-item');input.closest('ul,ol')?.classList.add('contains-task-list');
 }
 const errorBlock=(kind:string,raw:string,error:unknown,display=true)=>{
  const element=document.createElement(display?'pre':'code');element.className=kind+'-error';element.setAttribute('role','alert');element.setAttribute('data-error',String(error).slice(0,1000));element.textContent=raw;return element;
 };
 const formula=(raw:string,display:boolean)=>{
  try{math??=createMathRenderer(window);return math.render(raw,display)}catch(error){return errorBlock('math',raw,error,display)}
 };
 const tokens=new Map(prepared.tokens.map(t=>[t.id,t]));
 for(const code of [...document.querySelectorAll('code')]){
  const token=tokens.get(code.textContent??'');if(!token)continue;
  const output=formula(token.source,token.display);
  if(token.display&&code.parentElement?.tagName==='P'&&code.parentElement.childNodes.length===1)code.parentElement.replaceWith(output);else code.replaceWith(output);
 }
 for(const pre of [...document.querySelectorAll('pre[lang="math"]')])pre.replaceWith(formula(pre.textContent??'',true));
 const diagrams=[...document.querySelectorAll('pre[lang="mermaid"]')];
 if(diagrams.length){
  const {default:mermaid}=await import('mermaid');
  const digest=new Bun.CryptoHasher('sha256').update(source).digest('hex').slice(0,16);
  mermaid.initialize({startOnLoad:false,securityLevel:'strict',htmlLabels:false,flowchart:{htmlLabels:false},deterministicIds:true,deterministicIDSeed:digest,maxTextSize:50000});
  let i=0;
  for(const pre of diagrams){
   const raw=pre.textContent??'';
   try{
    if(raw.length>50000)throw Error('Diagram exceeds 50,000 characters');
    await mermaid.parse(raw);
    const id=`mermaid-${digest}-${i++}`;
    const {svg}=await mermaid.render(id,raw);
    const tree=Bun.XML.parse(svg,{compact:false});
    const figure=document.createElement('figure');figure.className='mermaid';
    const svgDoc=new window.DOMParser().parseFromString(Bun.XML.stringify(tree),'image/svg+xml');
    const svgNode=document.importNode(svgDoc.documentElement,true);
    // Keep scoped CSS alongside the SVG: happy-dom's HTML parser otherwise
    // drops its subtree when the fragment is read back by the caller.
    for(const style of [...svgNode.querySelectorAll('style')]){const css=document.createElement('style');css.textContent=style.textContent;figure.append(css);style.remove()}
    figure.append(svgNode);pre.replaceWith(figure);
   }catch(error){pre.replaceWith(errorBlock('mermaid',raw,error))}
  }
 }
 // Mermaid's temporary rendering nodes never become part of the output.
 for(const node of [...document.body.children])if(/^dmermaid-|^d?mermaid-error/.test(node.id))node.remove();
 const output=document.body.innerHTML;
 if(Buffer.byteLength(output)>32*1024*1024)throw new RangeError('Rendered HTML exceeds 32 MB');
 return output;
}finally{
 try{math?.dispose();document.body.replaceChildren();document.head.replaceChildren();await window.happyDOM.abort()}
 finally{restore()}
}
}

