import { parseFragment, serialize } from "parse5";

type Node = {
  nodeName: string;
  tagName?: string;
  attrs?: { name: string; value: string }[];
  childNodes?: Node[];
  content?: Node;
  parentNode?: Node;
};

/** Undo GitHub's Camo transformation in saved API responses, not arbitrary image links. */
export function normalizeGithubImages(html: string): string {
  const root = parseFragment(html) as Node;
  const attr = (node: Node, name: string) => node.attrs?.find((item) => item.name === name)?.value;
  const camo = (node: Node) =>
    node.tagName === "img" &&
    /^https:\/\/camo\.githubusercontent\.com\//.test(attr(node, "src") ?? "") &&
    /^https?:\/\//i.test(attr(node, "data-canonical-src") ?? "");
  const visit = (node: Node) => {
    for (const child of [...(node.childNodes ?? [])]) {
      const children = child.childNodes ?? [];
      const image = children.length === 1 ? children[0] : undefined;
      const proxyWrapper =
        child.tagName === "a" &&
        image &&
        camo(image) &&
        attr(child, "href") === attr(image, "src") &&
        attr(child, "target") === "_blank" &&
        attr(child, "rel") === "noopener noreferrer nofollow" &&
        child.attrs?.length === 3;
      if (proxyWrapper) {
        const index = node.childNodes!.indexOf(child);
        node.childNodes![index] = image;
        image.parentNode = node;
        visit(image);
      } else visit(child);
    }
    if (node.content) visit(node.content);
    if (camo(node)) {
      const canonical = attr(node, "data-canonical-src")!;
      node.attrs = node.attrs!.filter(
        (item) =>
          item.name !== "data-canonical-src" &&
          !(item.name === "style" && item.value === "max-width: 100%;"),
      );
      node.attrs.find((item) => item.name === "src")!.value = canonical;
    }
  };
  visit(root);
  return serialize(root as any);
}
