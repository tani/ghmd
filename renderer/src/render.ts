import {createDOM} from './dom';
import {nativeRender} from './native';
import {parseFragment,serialize} from 'parse5';
import createDOMPurify from 'dompurify';
import {alertTitles} from './alerts';
import {highlight} from './highlight';
import type {RenderOptions} from './index';

// GitHub assigns a unique suffix to footnote anchors for each rendered fragment.
export async function renderInProcess(source:string,options:RenderOptions):Promise<string>{
 const {html}=nativeRender(source);
 const window=createDOM();
 const document=window.document;
 try {
  document.body.innerHTML=createDOMPurify(window as any).sanitize(serialize(parseFragment(html)), {ADD_ATTR:['data-footnote-ref','data-footnote-backref','data-footnote-backref-idx','data-footnotes'], ADD_TAGS:['input']});
  const comments=document.createTreeWalker(document.body, window.NodeFilter.SHOW_COMMENT);
  const removeComments=[]; while(comments.nextNode()) removeComments.push(comments.currentNode);
  for(const comment of removeComments) comment.remove();
  const token=options.runId??crypto.randomUUID().replaceAll('-','');
  // The API leaves TeX for its client-side math renderer. It does not render
  // math inside footnotes, links, code, or inline formatting elements.
  const blocked=(element:any)=>{
   for(let p=element.parentElement;p;p=p.parentElement){
    if(p.tagName==='DETAILS') return false;
    if(['B','STRONG','EM','I','A','CODE','PRE','SECTION'].includes(p.tagName)) return true;
   }
   return false;
  };
  const mathElement=(content:string, display:boolean)=>{
   const math=document.createElement('math-renderer');
   math.className=display?'js-display-math':'js-inline-math';
   math.setAttribute('style',display?'display: block':'display: inline-block');
   math.setAttribute('data-run-id',token);
   math.textContent=content;
   return math;
  };
  for(const details of document.querySelectorAll('details')){
   if(!details.querySelector(':scope > summary')){
    const summary=document.createElement('summary');summary.textContent='Details';details.prepend(summary);
   }
  }
  for(const paragraph of [...document.querySelectorAll('p')]){
   if(blocked(paragraph)) continue;
   const children=[...paragraph.childNodes];
   // GitHub reconstructs underscore emphasis in display math after Markdown
   // parsing, adding a blank line around the reconstructed expression.
   if(!paragraph.closest('b') && children.some(n=>n.nodeName==='EM') && children.every(n=>n.nodeType===3 || n.nodeName==='BR' || n.nodeName==='EM')){
    const reconstructed=children.map(n=>n.nodeName==='EM'?'_'+n.textContent+'_':n.nodeType===3?n.textContent:'').join('');
    if(/^\$\$\n[\s\S]*\n\$\$$/.test(reconstructed)){
     const inner=reconstructed.slice(3,-3);
     paragraph.replaceChildren(mathElement('$$\n\n'+inner+'\n\n$$',true));continue;
    }
   }
   // Display formulas contain only text and hard-break elements.
   if(children.every(n=>n.nodeType===3 || n.nodeName==='BR')){
    const text=paragraph.textContent??'';
    if(/^\$\$\n[\s\S]*\n\$\$$/.test(text) && !paragraph.closest('b')){
     paragraph.replaceChildren(mathElement(text.replace(/\\\n/g,'\\\\\n'),true));continue;
    }
   }
   // Dollar-backtick spans are parsed as a code element surrounded by dollars.
   for(const code of [...paragraph.querySelectorAll(':scope > code')]){
    const before=code.previousSibling,after=code.nextSibling;
    if(before?.nodeType!==3 || after?.nodeType!==3 || !before.textContent?.endsWith('$') || !after.textContent?.startsWith('$')) continue;
    const value=code.textContent??'';
    before.textContent=before.textContent.slice(0,-1);after.textContent=after.textContent.slice(1);
    // GitHub's math filter escapes an ampersand already escaped by HTML.
    code.replaceWith(mathElement('$'+value.replaceAll('&','&amp;')+'$',false));
   }
   for(const node of [...paragraph.childNodes]){
    if(node.nodeType!==3 || !node.textContent?.includes('$')) continue;
    const text=node.textContent, re=/\$(?!\$)([^$\n]+?)\$/g;
    let match:RegExpExecArray|null,last=0;const fragment=document.createDocumentFragment();
    while((match=re.exec(text))){
     if(/^\s|\s$/.test(match[1]) || match[1].includes('&')) continue;
     fragment.append(document.createTextNode(text.slice(last,match.index)),mathElement(match[0],false));
     last=match.index+match[0].length;
    }
    if(last){fragment.append(document.createTextNode(text.slice(last)));node.replaceWith(fragment)}
   }
  }
  for(const pre of [...document.querySelectorAll('pre[lang="math"]')]){
   if(blocked(pre)) continue;
   const content=(pre.textContent??'').replace(/\n$/,'').replaceAll('\n','\\\n');
   pre.replaceWith(mathElement('$$'+content+'$$',true));
  }
  for(const pre of [...document.querySelectorAll('pre[lang]')]) await highlight(document,pre);
  for(const code of document.querySelectorAll('code')) code.classList.add('notranslate');
  for(const pre of document.querySelectorAll('pre')) {
   if(pre.getAttribute('lang')==='math') pre.removeAttribute('lang');
   pre.classList.add('notranslate');
  }
  for(const table of [...document.querySelectorAll('table')]){
   table.setAttribute('role','table');
   const wrapper=document.createElement('markdown-accessiblity-table');
   table.replaceWith(wrapper); wrapper.append(table);
  }
  for(const quote of [...document.querySelectorAll('blockquote')]){
   const paragraph=quote.firstElementChild;
   if(!paragraph || paragraph.tagName!=='P' || !paragraph.firstChild || paragraph.firstChild.nodeType!==3) continue;
   const marker=paragraph.firstChild.textContent?.match(/^\[!(NOTE|TIP|IMPORTANT|WARNING|CAUTION)\]$/);
   if(!marker || paragraph.firstChild.nextSibling?.nodeName!=='BR') continue;
   const kind=marker[1].toLowerCase();
   paragraph.firstChild.nextSibling.remove(); paragraph.firstChild!.remove();
   if(paragraph.firstChild?.nodeType===3) paragraph.firstChild.textContent=paragraph.firstChild.textContent!.replace(/^\n/,'');
   const alert=document.createElement('div'); alert.className=`markdown-alert markdown-alert-${kind}`;
   const title=document.createElement('p'); title.className='markdown-alert-title'; title.innerHTML=alertTitles[kind];
   alert.append(title);
   while(quote.firstChild) alert.append(quote.firstChild);
   if(alert.childNodes[1]?.nodeType===3 && !alert.childNodes[1].textContent?.trim()) alert.childNodes[1].remove();
   quote.replaceWith(alert);
  }
  for(const input of document.querySelectorAll('input[type="checkbox"]')){
   input.closest('li')?.classList.add('task-list-item');
   input.closest('ul,ol')?.classList.add('contains-task-list');
   const checked=input.hasAttribute('checked');
   input.removeAttribute('checked'); input.removeAttribute('disabled');
   input.setAttribute('id','');
   if(checked) input.setAttribute('checked','');
   input.setAttribute('disabled','');
   input.classList.add('task-list-item-checkbox');
   input.setAttribute('aria-label',input.hasAttribute('checked')?'Completed task':'Incomplete task');
  }
  for(const link of document.querySelectorAll('a[href]')){
   const href=link.getAttribute('href')!;
   if(/^(?:https?:)?\/\//i.test(href)||/^mailto:/i.test(href)) link.setAttribute('rel','nofollow');
  }
  if(options.imageProxy){
   for(const image of [...document.querySelectorAll('img[src]')]){
    const canonical=image.getAttribute('src')!;
    if(!/^https?:\/\//i.test(canonical)) continue;
    const proxy=await options.imageProxy(canonical);
    if(!/^https?:\/\//i.test(proxy)) throw new TypeError('imageProxy must return an HTTP(S) URL');
    image.setAttribute('src',proxy);
    image.setAttribute('data-canonical-src',canonical);
    image.setAttribute('style','max-width: 100%;');
    if(!image.closest('a')){
     const link=document.createElement('a');
     link.setAttribute('target','_blank');link.setAttribute('rel','noopener noreferrer nofollow');link.setAttribute('href',proxy);
     image.replaceWith(link);link.append(image);
    }
   }
  }
  for(const sup of document.querySelectorAll('sup.footnote-ref')) sup.removeAttribute('class');
  for(const a of document.querySelectorAll('a[data-footnote-ref]')){
   const href=a.getAttribute('href');
   const id=a.getAttribute('id');
   if(href){
    const name=href.slice(4);
    if(name===':') a.removeAttribute('href');
    else a.setAttribute('href',`#user-content-fn-${name}-${token}`);
   }
   if(id) a.setAttribute('id',`user-content-${id.toLowerCase()}-${token}`);
   a.setAttribute('aria-describedby','footnote-label');
  }
  for(const a of document.querySelectorAll('a[data-footnote-backref]')){
   const href=a.getAttribute('href');
   if(href?.slice(7)===':'){a.replaceWith(document.createTextNode('↩'));continue}
   if(href) a.setAttribute('href',`#user-content-${href.slice(1)}-${token}`);
   a.removeAttribute('class');
   a.removeAttribute('data-footnote-backref-idx');
   a.setAttribute('class','data-footnote-backref');
  }
  const section=document.querySelector('section[data-footnotes]');
  if(section){
   section.removeAttribute('class');section.removeAttribute('data-footnotes');
   section.setAttribute('data-footnotes','');section.setAttribute('class','footnotes');
   for(const li of section.querySelectorAll('li[id]')) li.id=`user-content-${li.id}-${token}`;
   const title=document.createElement('h2');title.id='footnote-label';title.className='sr-only';title.textContent='Footnotes';
   section.insertBefore(title,section.firstChild);
   const next=section.nextSibling;
   if(section.parentElement!==document.body && next?.nodeType===3 && !next.textContent?.trim() && !next.nextSibling) next.remove();
  }
  const output=document.body.innerHTML.replace(/\n$/,'');
  if(Buffer.byteLength(output)>32*1024*1024)throw new RangeError('Rendered HTML exceeds 32 MB');
  return output;
 } finally {await window.happyDOM.abort()}
}
