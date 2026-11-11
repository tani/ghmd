/** Held-out malformed inputs; expectations come from the established renderer. */
export function htmlScopeCases(): { id: string; source: string }[] {
  const cases: { id: string; source: string }[] = [];
  const formatting = ["b", "i", "strong", "em", "code", "del"];
  for (const outer of formatting)
    for (const inner of formatting) {
      if (outer === inner) continue;
      for (const closed of [false, true]) {
        cases.push({
          id: `misnested-${outer}-${inner}-${closed}`,
          source: `<${outer}><${inner}>one</${outer}> two${closed ? `</${inner}>` : ""}\n\n$x$\n\nlast`,
        });
      }
    }
  for (const prefix of ["", "<b>", "<em>", "<details><summary>More</summary>"]) {
    for (const href of ["https://example.com", "javascript:alert(1)"]) {
      for (const ending of ["\n\n$x$", "\n\nplain", "\n\n- $x$", "<div>block</div> $x$"]) {
        cases.push({
          id: `anchor-${cases.length}`,
          source: `${prefix}<a href="${href}">first${ending}`,
        });
      }
    }
  }
  return cases;
}
