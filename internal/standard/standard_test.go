package standard

import "testing"

func TestRender(t *testing.T) {
	out, err := Render("# hi\n\n| a | b |\n|---|---|\n| 1 | 2 |\n", true)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(out)
}
