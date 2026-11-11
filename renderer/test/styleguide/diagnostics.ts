/** Diagnose case boundaries in the exact HTML from a batched document.
 * These are slices of a single rendering, not independent render results. */
export type CaseDiagnostic={id:string;status:'pass'|'fail'|'not-run';scope:'batch-context'|'whole-document';reason?:string;firstDifference?:{line:number;column:number;expected:string;actual:string}};
export function normalizeGeneratedIds(text:string){
 return text.replace(/(data-run-id="|user-content-fn(?:ref)?-[^"#\s]*?-)[0-9a-f]{32}(?=["#\s])/g,'$1TOKEN');
}
function firstDifference(expected:string,actual:string){
 const a=expected.split('\n'),b=actual.split('\n');
 for(let i=0;i<Math.max(a.length,b.length);i++)if(a[i]!==b[i]){
  const left=a[i]??'<end of output>',right=b[i]??'<end of output>';
  let column=0;while(column<left.length && column<right.length && left[column]===right[column])column++;
  return {line:i+1,column:column+1,expected:left,actual:right};
 }
}
function slices(html:string,ids:string[]):string[]|undefined{
 const markers=ids.map(id=>`<h2>AUDITBOUNDARY-${id}</h2>`);
 const positions=markers.map(marker=>html.indexOf(marker));
 if(positions.some((pos,i)=>pos<0 || html.indexOf(markers[i],pos+markers[i].length)!==-1 || (i>0 && pos<=positions[i-1])))return;
 // Preserve preamble and footer: otherwise unassigned differences can be lost.
 return positions.map((pos,i)=>html.slice(i===0?0:pos,positions[i+1]??html.length));
}
export function diagnoseCases(ids:string[],expected:string|undefined,actual:string|undefined,wrapped:boolean):CaseDiagnostic[]{
 if(expected===undefined || actual===undefined)return ids.map(id=>({id,status:'not-run',scope:wrapped?'batch-context':'whole-document',reason:'Expected or actual HTML is unavailable.'}));
 if(!wrapped){
  if(ids.length!==1)return ids.map(id=>({id,status:'not-run',scope:'whole-document',reason:'A multi-case document has no separable boundaries.'}));
  const difference=firstDifference(normalizeGeneratedIds(expected),normalizeGeneratedIds(actual));
  return [{id:ids[0],status:difference?'fail':'pass',scope:'whole-document',...(difference?{firstDifference:difference}:{})}];
 }
 const a=slices(expected,ids),b=slices(actual,ids);
 if(!a||!b)return ids.map(id=>({id,status:'not-run',scope:'batch-context',reason:'A boundary is missing, duplicated, or out of order; inspect the full document.'}));
 return ids.map((id,i)=>{
  const difference=firstDifference(normalizeGeneratedIds(a[i]),normalizeGeneratedIds(b[i]));
  return {id,status:difference?'fail':'pass',scope:'batch-context',...(difference?{firstDifference:difference}:{})};
 });
}
