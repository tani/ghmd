import {describe,expect,test} from 'bun:test';
import {diagnoseCases} from './styleguide/diagnostics';
const heading=(id:string)=>`<h2>AUDITBOUNDARY-${id}</h2>`;

describe('batch case diagnostics',()=>{
 test('attributes an exact difference to the correct case',()=>{
  const expected=`${heading('one')}\n<p>same</p>\n${heading('two')}\n<p>expected</p>`;
  const actual=`${heading('one')}\n<p>same</p>\n${heading('two')}\n<p>actual</p>`;
  const results=diagnoseCases(['one','two'],expected,actual,true);
  expect(results[0]).toMatchObject({id:'one',status:'pass',scope:'batch-context'});
  expect(results[1]).toMatchObject({id:'two',status:'fail',firstDifference:{line:2,expected:'<p>expected</p>',actual:'<p>actual</p>'}});
 });
 test('reports missing, duplicated or reordered markers as not-run',()=>{
  const expected=heading('one')+heading('two');
  for(const actual of [heading('one'),heading('one')+heading('two')+heading('two'),heading('two')+heading('one')])
   expect(diagnoseCases(['one','two'],expected,actual,true).map(x=>x.status)).toEqual(['not-run','not-run']);
 });
 test('preserves preamble and final trailing content in comparison',()=>{
  const expected='before'+heading('one')+'tail';
  expect(diagnoseCases(['one'],expected,'changed'+heading('one')+'tail',true)[0].status).toBe('fail');
  expect(diagnoseCases(['one'],expected,'before'+heading('one')+'changed',true)[0].status).toBe('fail');
 });
 test('does not confuse batch slices with standalone results',()=>{
  expect(diagnoseCases(['case'],'same','same',false)[0]).toEqual({id:'case',status:'pass',scope:'whole-document'});
  expect(diagnoseCases(['a','b'],'same','same',false).map(x=>x.status)).toEqual(['not-run','not-run']);
  expect(diagnoseCases(['case'],undefined,undefined,false)[0].status).toBe('not-run');
 });
 test('normalizes only generated render identifiers',()=>{
  const a='<math-renderer data-run-id="aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa">$x$</math-renderer>';
  const b='<math-renderer data-run-id="bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb">$x$</math-renderer>';
  expect(diagnoseCases(['math'],a,b,false)[0].status).toBe('pass');
  expect(diagnoseCases(['math'],a,b.replace('$x$','$y$'),false)[0].status).toBe('fail');
 });
});
