import { Window } from "happy-dom";
/** Create a private DOM for a render. */
export function createDOM() {
  return new Window({
    settings: {
      enableJavaScriptEvaluation: false,
      disableJavaScriptFileLoading: true,
      disableCSSFileLoading: true,
      disableIframePageLoading: true,
      disableComputedStyleRendering: true,
    },
  });
}
