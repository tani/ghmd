import {describe,test,expect} from 'bun:test';
import {Window} from 'happy-dom';
import {render} from '../src/index';
import {nativeRender,nativeVersion} from '../src/native';
async function inspect(source:string){const html=await render(source);const w=new Window();w.document.body.innerHTML=html;return {html,d:w.document,close:()=>w.happyDOM.abort()}}

describe('render :: String -> Promise<String>',()=>{
 test('async string and native GFM extensions',async()=>{
  expect(nativeVersion).toBe('0.29.0.gfm.13');
  const r=await inspect('# 日本語\n\n| A | B |\n|---|---|\n| 1 | 2 |\n\n- [x] done\n\n~~old~~\n\nhttps://example.com');
  try{expect(typeof r.html).toBe('string');expect(r.d.querySelector('h1')?.textContent).toBe('日本語');expect(r.d.querySelectorAll('td').length).toBe(2);expect(r.d.querySelector('input')?.checked).toBe(true);expect(r.d.querySelector('del')?.textContent).toBe('old');expect(r.d.querySelector('a')?.href).toBe('https://example.com/')}finally{await r.close()}
 });
 test('real FFI emits ordered AST XML with UTF-8 source positions',()=>{
  const {html,astXML}=nativeRender('あ **太字**\n\n[^a]\n\n[^a]: note');
  expect(html).toContain('<strong>太字</strong>');expect(html).toContain('data-footnote-ref');
  const tree=Bun.XML.parse(astXML,{compact:false});expect(tree.name).toBe('document');expect(tree.attributes.xmlns).toBe('http://commonmark.org/xml/1.0');
 });
 test('MathJax renders inline, protected, display and fenced math into SVG',async()=>{
  const r=await inspect('あ $x^2$ と $`\\{a\\,b\\}`$\n\n$$\n\\frac{1}{2}\n$$\n\n```math\n\\begin{pmatrix}1 & 2\\\\3 & 4\\end{pmatrix}\n```');
  try{expect(r.d.querySelectorAll('.math svg').length).toBe(4);expect(r.d.querySelectorAll('.math-display').length).toBe(2);expect(r.d.querySelector('.math-error')).toBeNull();expect([...r.d.querySelectorAll('.math')].map(e=>e.getAttribute('data-tex'))).toContain('\\{a\\,b\\}');for(const svg of r.d.querySelectorAll('.math svg'))expect(Bun.XML.parse(svg.outerHTML,{compact:false}).name).toBe('svg')}
  finally{await r.close()}
 });
 test('ordinary code and link destinations stay opaque',async()=>{
  const r=await inspect('`$x$`\n\n```text\n$y$\n```\n\n[price](https://example.com/$z$)\n\nEscaped \\$a\\$');
  try{expect(r.d.querySelector('.math')).toBeNull();expect(r.d.querySelector('a')?.getAttribute('href')).toContain('$z$');expect(r.d.body.textContent).toContain('$x$')}finally{await r.close()}
 });
 test('invalid TeX is an escaped visible error, without leaking macros',async()=>{
  const r=await inspect('$`\\newcommand{\\private}{x}\\private`$ and $`\\private`$ and $`\\notacommand`$');
  try{expect(r.d.querySelectorAll('.math svg').length).toBe(1);expect(r.d.querySelectorAll('.math-error').length).toBe(2);expect(r.d.querySelector('.math-error')?.textContent).toBe('\\private')}finally{await r.close()}
 });
 test('Mermaid actually generates sanitized SVG and retains both node labels',async()=>{
  const r=await inspect('```mermaid\nflowchart TD\n A[Input] --> B[Output]\n```');
  try{expect(r.d.querySelector('figure.mermaid svg')).not.toBeNull();expect(r.d.querySelector('.mermaid-error')).toBeNull();expect(r.d.querySelector('figure.mermaid')?.textContent).toContain('Input');expect(r.d.querySelector('figure.mermaid')?.textContent).toContain('Output');expect(Bun.XML.parse(r.d.querySelector('figure.mermaid svg')!.outerHTML,{compact:false}).name).toBe('svg')}
  finally{await r.close()}
 });
 test('invalid Mermaid is visible source rather than false success',async()=>{
  const r=await inspect('```mermaid\nthis is not mermaid\n```');
  try{expect(r.d.querySelector('figure.mermaid')).toBeNull();expect(r.d.querySelector('.mermaid-error')?.textContent).toContain('this is not mermaid');expect(r.d.querySelector('.mermaid-error')?.getAttribute('role')).toBe('alert')}finally{await r.close()}
 });
 test('raw HTML and dangerous Markdown URLs do not execute or survive',async()=>{
  const r=await inspect('<script>globalThis.pwned=true</script>\n\n[x](javascript:alert(1))\n\n<img src=x onerror=alert(1)>');
  try{expect(r.d.querySelector('script,img')).toBeNull();expect(r.d.querySelector('a')?.getAttribute('href')).toBe('');expect(r.html).not.toContain('onerror=')}finally{await r.close()}
 });
 test('source reaching Mermaid cannot emit script or event attributes',async()=>{
  const r=await inspect('```mermaid\nflowchart TD\n A["<img src=x onerror=alert(1)>"] --> B[safe]\n```');
  try{expect(r.d.querySelector('script,[onerror],[onload]')).toBeNull();expect(r.html).not.toContain('href="javascript:')}finally{await r.close()}
 });
 test('identical Markdown has deterministic output and no caller global pollution',async()=>{
  const source='Math $`x`$\n\n```mermaid\nflowchart LR\n A --> B\n```';
  const before=(globalThis as any).document;expect(await render(source)).toBe(await render(source));expect((globalThis as any).document).toBe(before);
 });
 test('bounds and input types reject asynchronously',async()=>{
  await expect(render(null as any)).rejects.toBeInstanceOf(TypeError);await expect(render('x'.repeat(1000001))).rejects.toBeInstanceOf(RangeError);
 });
 test('render returns a Promise without calling process creation APIs',async()=>{
  const {spyOn}=await import('bun:test');
  const spawn=spyOn(Bun,'spawn').mockImplementation(()=>{throw Error('Forbidden child process')});
  const spawnSync=spyOn(Bun,'spawnSync').mockImplementation(()=>{throw Error('Forbidden child process')});
  try{
   const result=render('**hello**\n\n```mermaid\nflowchart TD\n A --> B\n```');
   expect(result).toBeInstanceOf(Promise);
   const html=await result;expect(html).toContain('<strong>hello</strong>');expect(html).toContain('class="mermaid"');
   expect(spawn).not.toHaveBeenCalled();expect(spawnSync).not.toHaveBeenCalled();
  }finally{spawn.mockRestore();spawnSync.mockRestore()}
 });
 test('parallel calls keep labels and configuration separate',async()=>{
  const sources=['Alpha','Beta','Gamma'].map(label=>'# '+label+'\n\n$`x`$\n\n```mermaid\nflowchart TD\n A['+label+'] --> B[End]\n```');
  const htmls=await Promise.all(sources.map(render));
  for(let i=0;i<htmls.length;i++){
   expect(htmls[i]).toContain('<h1>'+['Alpha','Beta','Gamma'][i]+'</h1>');
   expect(htmls[i]).toContain('class="mermaid"');
   for(let j=0;j<htmls.length;j++)if(i!==j)expect(htmls[i]).not.toContain(['Alpha','Beta','Gamma'][j]);
  }
 });
 test('global descriptors are restored and operational rejection does not poison the queue',async()=>{
  const keys=['window','document','DOMParser','XMLSerializer','HTMLElement','SVGElement','CSSStyleSheet','CSSRule','Node','getComputedStyle','navigator'];
  const before=keys.map(k=>Object.getOwnPropertyDescriptor(globalThis,k));
  // A raw XML-forbidden control character makes cmark AST XML invalid.
  await expect(render('bad \x01')).rejects.toBeInstanceOf(SyntaxError);
  expect(await render('recovered')).toContain('<p>recovered</p>');
  expect(keys.map(k=>Object.getOwnPropertyDescriptor(globalThis,k))).toEqual(before);
 });
});
