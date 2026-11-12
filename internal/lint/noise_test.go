package lint

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestFixtureDifferentialNoiseBudget guards the differential linter against
// regressions in dialect-noise suppression. The captured GitHub fixtures are
// adversarial, so raw-HTML findings are expected; expected-renderer
// regressions should stay rare.
func TestFixtureDifferentialNoiseBudget(t *testing.T) {
	directory := filepath.Join("..", "..", "test", "fixtures", "e2e", "cases")
	entries, err := os.ReadDir(directory)
	if err != nil {
		t.Skip(err)
	}
	differential, mathDifferential, mathSource, raw, style := 0, 0, 0, 0, 0
	var noisy []string
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}
		source, err := os.ReadFile(filepath.Join(directory, entry.Name()))
		if err != nil {
			t.Fatal(err)
		}
		findings, err := Source(string(source))
		if err != nil {
			t.Fatalf("%s: %v", entry.Name(), err)
		}
		for _, finding := range findings {
			switch {
			case finding.Rule == "standard/lost-math":
				mathDifferential++
			case strings.HasPrefix(finding.Rule, "math/"):
				mathSource++
			case strings.HasPrefix(finding.Rule, "standard/"):
				differential++
				noisy = append(noisy, entry.Name()+" "+finding.Rule)
			case strings.HasPrefix(finding.Rule, "style/"):
				style++
			default:
				raw++
			}
		}
	}
	if differential > 5 {
		t.Fatalf("differential lint noise regressed: %d findings (%v)", differential, noisy)
	}
	if mathDifferential > 120 {
		t.Fatalf("math differential findings regressed: %d", mathDifferential)
	}
	if mathSource > 120 {
		t.Fatalf("math source findings regressed: %d", mathSource)
	}
	if raw > 100 {
		t.Fatalf("raw-HTML lint noise regressed: %d findings", raw)
	}
	if style > 1500 {
		t.Fatalf("authoring-style lint findings regressed: %d", style)
	}
}
