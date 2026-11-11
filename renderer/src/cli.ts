import {render} from './index';
const path=process.argv[2];
try{const source=path&&path!=='-'?await Bun.file(path).text():await Bun.stdin.text();process.stdout.write(await render(source))}
catch(error){console.error(String(error));process.exitCode=1}
