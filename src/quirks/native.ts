import { renderMarkdown as renderNative, renderAstXml as renderXml } from "./native-core";
import { removeRawInputs } from "./raw-html";
export * from "./native-core";
export const renderMarkdown = (source: string): string => renderNative(source, removeRawInputs);
export const renderAstXml = (source: string): string => renderXml(source, removeRawInputs);
