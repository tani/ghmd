/** GitHub retains at most ten nested lists of each type, flattening deeper items. */
export function flattenDeepLists(document: Document): void {
  for (const list of [...document.querySelectorAll("ul, ol")]) {
    let sameTypeDepth = 1;
    for (
      let parent = list.parentElement?.closest("ul, ol");
      parent;
      parent = parent.parentElement?.closest("ul, ol") ?? null
    ) {
      if (parent.tagName === list.tagName) sameTypeDepth++;
    }
    if (sameTypeDepth <= 10) continue;
    const item = list.parentElement;
    const parentList = item?.parentElement;
    if (item?.tagName !== "LI" || !parentList || !["UL", "OL"].includes(parentList.tagName))
      continue;
    const anchor = item.nextSibling;
    for (const child of [...list.children]) {
      if (child.tagName !== "LI") continue;
      child.classList.remove("task-list-item");
      if (!child.className) child.removeAttribute("class");
      parentList.insertBefore(child, anchor);
    }
    list.remove();
  }
}
