import {mathjax} from '@mathjax/src/js/mathjax.js';
import {TeX} from '@mathjax/src/js/input/tex.js';
import {SVG} from '@mathjax/src/js/output/svg.js';
import {HTMLAdaptor} from '@mathjax/src/js/adaptors/HTMLAdaptor.js';
import {Document, DocumentFragment, HTMLElement} from 'happy-dom';
import {RegisterHTMLHandler} from '@mathjax/src/js/handlers/html.js';
import '@mathjax/src/js/input/tex/base/BaseConfiguration.js';
import '@mathjax/src/js/input/tex/ams/AmsConfiguration.js';
import '@mathjax/src/js/input/tex/newcommand/NewcommandConfiguration.js';
import '@mathjax/src/js/input/tex/color/ColorConfiguration.js';
export function createMathRenderer(window:any){
 // happy-dom Window constructors do not share the exported base-class chain.
 const adaptorWindow=Object.create(window,{Document:{value:Document},DocumentFragment:{value:DocumentFragment},HTMLElement:{value:HTMLElement}});
 const adaptor=new HTMLAdaptor(adaptorWindow);const handler=RegisterHTMLHandler(adaptor);let counter=0;
 return {
  render(source:string,display:boolean){
   if(source.length>20000)throw Error('Formula exceeds 20,000 characters');
   const tex=new TeX({packages:['base','ams','newcommand','color'],maxMacros:1000,maxBuffer:20000,formatError(_jax:any,error:any){throw error}});
   const svg=new SVG({fontCache:'local',localID:`math-${counter++}`});
   const doc=mathjax.document('',{InputJax:tex,OutputJax:svg});
   const node=doc.convert(source,{display,em:16,ex:8,containerWidth:1280});
   const element=window.document.createElement('span');element.innerHTML=adaptor.outerHTML(node);
   if(element.querySelector('[data-mjx-error],merror'))throw Error('MathJax produced an error node');
   const tree=Bun.XML.parse(element.querySelector('svg')!.outerHTML,{compact:false});
   const holder=window.document.createElement(display?'div':'span');holder.className=display?'math math-display':'math math-inline';
   holder.setAttribute('data-tex',source);holder.innerHTML=Bun.XML.stringify(tree);return holder;
  },
  dispose(){mathjax.handlers.unregister(handler)}
 };
}
