package lint

// A manual rule describes an authoring or implementation responsibility that
// cannot be inferred reliably from Markdown syntax alone. Manual rules are
// intentionally excluded from automatic style lint recall; claiming otherwise
// would turn formatting policy into false-positive engine behavior.
var manualStyleRules = map[string]string{
	"BLOCK-003":  "A visible line break is author intent; ordinary soft-wrapped paragraphs are not a violation.",
	"TASK-002":   "A task checkbox is document data, not a portable interactive control.",
	"COMPAT-002": "Content preservation and the selected compatibility profile require author intent.",
	"COMPAT-003": "Specification behavior, a current defect, a fixed historical bug, and an untested risk require external evidence.",
	"RENDER-001": "Renderer parsing and math-node construction are implementation responsibilities.",
	"RENDER-002": "Reproducible parser and renderer configuration is a build responsibility.",
	"RENDER-003": "HTML and identifier security policy is an application responsibility.",
	"RENDER-004": "Execution limits and failure isolation are renderer operations.",
	"RENDER-005": "Source preservation and revision-safe edits are renderer operations.",
	"RENDER-006": "Full-document and resource failure coverage is a test-plan responsibility.",
	"RENDER-007": "Streaming-completion handling is a renderer state-machine responsibility.",
}

// ManualRule returns review guidance for a rule, and whether it is deliberately
// excluded from automatic linting. Rule IDs match GFM_STYLEGUIDE.md exactly.
func ManualRule(id string) (string, bool) {
	guidance, ok := manualStyleRules[id]
	return guidance, ok
}
