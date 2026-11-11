import {renderInProcess} from './render';

// Mermaid is a singleton and DOMPurify retains its first DOM. Queue full
// renders so parallel callers cannot overwrite another request's document.
let queue:Promise<void>=Promise.resolve();

/** Markdown -> HTML fragment, in the calling Bun process. No child process.
 * Syntax errors become visible HTML error blocks; operational failures reject.
 */
export async function render(source:string):Promise<string>{
 if(typeof source!=='string')throw new TypeError('render expects a string');
 if(Buffer.byteLength(source)>1_000_000)throw new RangeError('Markdown exceeds 1 MB');
 const result=queue.then(()=>renderInProcess(source));
 // A failed render must not poison subsequent callers.
 queue=result.then(()=>{},()=>{});
 return result;
}
export default render;
