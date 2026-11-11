import {describe,expect,test} from 'bun:test';
import {render} from '../src/index';

describe('GitHub reference fragments',()=>{
 for(const batch of [1,2,3,4]) test(`batch ${batch}`,async()=>{
  const source=await Bun.file(`../github-markdown/batch-${batch}.md`).text();
  const expected=await Bun.file(`../github-html/batch-${batch}.html`).text();
  // GitHub signs Camo URLs with a server-side key. Supply the saved URL for
  // this fixture to test the surrounding HTML without inventing a signature.
  const camo=expected.match(/src="(https:\/\/camo\.githubusercontent\.com\/[^"]+)"[^>]*data-canonical-src="([^"]+)"/);
  const token=expected.match(/(?:data-run-id="|user-content-fn(?:ref)?-[^"#\s]*?-)([0-9a-f]{32})(?=["#\s])/)?.[1];
  expect(token).toBeDefined();
  const actual=await render(source,{runId:token,imageProxy:url=>{
   if(!camo || url!==camo[2]) throw Error(`No saved proxy URL for ${url}`);
   return camo[1];
  }});
  const tokens=[...actual.matchAll(/(?:data-run-id="|user-content-fn(?:ref)?-[^"#\s]*?-)([0-9a-f]{32})(?=["#\s])/g)].map(match=>match[1]);
  expect(new Set(tokens).size).toBe(1);
  expect(actual).toBe(expected);
 });
});
