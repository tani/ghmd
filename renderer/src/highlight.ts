import {all,createStarryNight} from '@wooorm/starry-night';

// Grammars are initialized once; each highlight call has its own token state.
let grammar:ReturnType<typeof createStarryNight>|undefined;
export async function highlight(document:any, pre:any):Promise<boolean>{
 const language=(pre.getAttribute('lang')||'').toLowerCase();
 if(!language || language==='math' || language==='text' || language==='plaintext') return false;
 const highlighter=await (grammar??=createStarryNight(all));
 const scope=highlighter.flagToScope(language==='geojson'||language==='topojson'?'json':language);
 if(!scope) return false;
 const source=(pre.textContent??'').replace(/\n$/,'');
 const result=highlighter.highlight(source,scope);
 const convert=(node:any):any=>{
  if(node.type==='text') return document.createTextNode(node.value);
  const element=document.createElement(node.tagName);
  if(node.properties.className) element.className=node.properties.className.join(' ');
  for(const child of node.children??[]) element.append(convert(child));
  // GitHub's Mermaid highlighter retains the nested scopes around quoted
  // node labels, while starry-night collapses the equivalent token scopes.
  if(scope==='source.mermaid' && element.className==='pl-s'){
   const children=[...element.childNodes] as any[];
   const first=children[0],last=children.at(-1);
   if(first?.className==='pl-sg' && first.textContent==='[' && last?.className==='pl-sg' && last.textContent===']' &&
      children[1]?.className==='pl-pds' && children.at(-2)?.className==='pl-pds'){
    const open=document.createElement('span');open.className='pl-pds';open.append(first);
    const content=document.createElement('span');content.className='pl-s';
    for(const child of children.slice(1,-1))content.append(child);
    const close=document.createElement('span');close.className='pl-pds';close.append(last);
    element.replaceChildren(open,content,close);
   }
  }
  return element;
 };
 const wrapper=document.createElement('div');
 wrapper.className=`highlight highlight-${scope.replaceAll('.','-')}`;
 const highlighted=document.createElement('pre');highlighted.className='notranslate';
 for(const child of result.children) highlighted.append(convert(child));
 wrapper.append(highlighted);pre.replaceWith(wrapper);
 return true;
}
