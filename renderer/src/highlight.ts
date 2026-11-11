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
  return element;
 };
 const wrapper=document.createElement('div');
 wrapper.className=`highlight highlight-${scope.replaceAll('.','-')}`;
 const highlighted=document.createElement('pre');highlighted.className='notranslate';
 for(const child of result.children) highlighted.append(convert(child));
 wrapper.append(highlighted);pre.replaceWith(wrapper);
 return true;
}
