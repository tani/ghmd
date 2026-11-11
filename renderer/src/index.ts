import {renderInProcess} from './render';
export {createGitHubImageProxy} from './github-image-proxy';

export interface RenderOptions {
 /** Use a specific GitHub-style 32-character lowercase hexadecimal render ID. */
 runId?:string;
 /** Resolve a canonical image URL to a signed image proxy URL. */
 imageProxy?: (url:string)=>string|Promise<string>;
}

/** Render a GitHub-style HTML fragment in the calling Bun process. */
export async function render(source:string, options:RenderOptions={}):Promise<string>{
 if(typeof source!=='string')throw new TypeError('render expects a string');
 if(options.runId!==undefined && !/^[a-f0-9]{32}$/.test(options.runId))throw new TypeError('runId must be 32 lowercase hexadecimal characters');
 if(Buffer.byteLength(source)>1_000_000)throw new RangeError('Markdown exceeds 1 MB');
 return renderInProcess(source,options);
}
export default render;
