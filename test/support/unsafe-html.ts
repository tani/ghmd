import { parseFragment } from "parse5";

/** Inspect the independently parsed output for a limited set of unsafe constructs. */
export function unsafeHtmlFindings(html: string): string[] {
  const findings: string[] = [];
  const visit = (node: any) => {
    if (
      ["script", "iframe", "object", "embed", "base", "meta", "link", "style"].includes(
        node.tagName,
      )
    )
      findings.push(`element:${node.tagName}`);
    for (const attr of node.attrs ?? []) {
      if (/^on/i.test(attr.name) || ["srcdoc", "formaction"].includes(attr.name))
        findings.push(`attribute:${attr.name}`);
      if (["href", "src", "action", "background", "poster", "cite"].includes(attr.name)) {
        const value = attr.value.replace(
          /[\u0000-\u0020\u00a0\u1680\u180e\u2000-\u2029\u205f\u3000]/g,
          "",
        );
        if (/^(?:javascript|vbscript|data):/i.test(value)) findings.push(`url:${attr.name}`);
      }
    }
    for (const child of node.childNodes ?? []) visit(child);
    if (node.content) visit(node.content);
  };
  visit(parseFragment(html));
  return findings;
}
