import {Window,Element} from 'happy-dom';
/** Create the persistent DOM used by in-process rendering. */
export function createDOM(){
 const window=new Window({settings:{enableJavaScriptEvaluation:false,disableJavaScriptFileLoading:true,disableCSSFileLoading:true,disableIframePageLoading:true,disableComputedStyleRendering:true}});
 // happy-dom's base Node.nodeName getter returns ''. DOMPurify 3.4 caches
 // and calls that getter on subclasses, as real browser implementations allow.
 // Supply the standard result rather than weakening Mermaid's sanitizer.
 const elementName=Object.getOwnPropertyDescriptor(Element.prototype,'nodeName')!.get!;
 Object.defineProperty(window.Node.prototype,'nodeName',{configurable:true,get(){
  switch(this.nodeType){case 1:return elementName.call(this);case 3:return '#text';case 8:return '#comment';case 9:return '#document';case 11:return '#document-fragment';case 10:return this.name;case 7:return this.target;default:return ''}
 }});
 // happy-dom 21 drops an SVG subtree when its HTML parser meets <style>.
 // Parse standalone SVG as XML, then import it into an HTML document. The
 // DOMPurify sanitizer still runs on that tree with its original strict rules.
 const BaseDOMParser=window.DOMParser;
 class SVGCompatibleDOMParser extends BaseDOMParser {
  parseFromString(input:string,mime:any):any {
   if(mime==='text/html'&&/^\s*<svg(?:\s|>)/i.test(String(input))){
    Bun.XML.parse(String(input),{compact:false});
    const svgDoc=super.parseFromString(input,'image/svg+xml');
    const doc=new window.HTMLDocument();doc.body.appendChild(doc.importNode(svgDoc.documentElement,true));return doc;
   }
   return super.parseFromString(input,mime);
  }
 }
 Object.defineProperty(window,'DOMParser',{value:SVGCompatibleDOMParser,configurable:true});

 return window;
}

/** Libraries expect browser globals. Restore the caller's exact descriptors,
 * including properties that were absent, once a render settles. */
export function installDOMGlobals(window:ReturnType<typeof createDOM>){
 const globals={window,document:window.document,DOMParser:window.DOMParser,XMLSerializer:window.XMLSerializer,HTMLElement:window.HTMLElement,SVGElement:window.SVGElement,CSSStyleSheet:window.CSSStyleSheet,CSSRule:window.CSSRule,Node:window.Node,getComputedStyle:window.getComputedStyle.bind(window),navigator:window.navigator};
 const previous=new Map<string,PropertyDescriptor|undefined>();
 const restore=()=>{for(const [key,descriptor] of [...previous].reverse()){if(descriptor)Object.defineProperty(globalThis,key,descriptor);else Reflect.deleteProperty(globalThis,key)}};
 try{for(const [key,value] of Object.entries(globals)){previous.set(key,Object.getOwnPropertyDescriptor(globalThis,key));Object.defineProperty(globalThis,key,{value,writable:true,enumerable:true,configurable:true})}}
 catch(error){restore();throw error}
 return restore;
}
