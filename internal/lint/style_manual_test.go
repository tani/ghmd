package lint

import "testing"

func TestManualRulesAreExplicitlyExcluded(t *testing.T) {
	for _, id := range []string{"BLOCK-003", "TASK-002", "COMPAT-002", "COMPAT-003", "RENDER-001"} {
		guidance, ok := ManualRule(id)
		if !ok || guidance == "" {
			t.Fatalf("expected manual review guidance for %s", id)
		}
	}
	if _, ok := ManualRule("BLOCK-001"); ok {
		t.Fatal("duplicate top-level title is detectable and must not be excluded")
	}
}
