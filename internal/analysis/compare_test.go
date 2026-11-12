package analysis

import (
	"reflect"
	"testing"
)

func TestCompareFragmentsNormalizesWhitespaceAndAttributeOrder(t *testing.T) {
	left := `<div z="last" a="first">same
	text</div>`
	right := `<div a="first" z="last">same text</div>`
	if differences := CompareFragments(left, right); len(differences) != 0 {
		t.Fatalf("normalized fragments differ: %v", differences)
	}
}

func TestCompareFragmentsReportsTextDifferenceAtChildPath(t *testing.T) {
	got := CompareFragments(`<p>left</p>`, `<p>right</p>`)
	want := []Difference{{Path: "/0/0", Want: "left", Got: "right"}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("differences = %#v, want %#v", got, want)
	}
}

func TestCompareFragmentsReportsAttributesWithSortedRendering(t *testing.T) {
	got := CompareFragments(`<p z="same" a="left">text</p>`, `<p a="right" z="same">text</p>`)
	want := []Difference{{Path: "/0", Want: "<p a=left z=same>", Got: "<p a=right z=same>"}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("differences = %#v, want %#v", got, want)
	}
}

func TestCompareFragmentsReportsMissingChildrenOnEitherSide(t *testing.T) {
	left := `<div><span>same</span><b>extra</b></div>`
	right := `<div><span>same</span></div>`
	got := CompareFragments(left, right)
	want := []Difference{{Path: "/0/1", Want: "<b>", Got: "<missing>"}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("extra left child differences = %#v, want %#v", got, want)
	}

	got = CompareFragments(right, left)
	want = []Difference{{Path: "/0/1", Want: "<missing>", Got: "<b>"}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("extra right child differences = %#v, want %#v", got, want)
	}
}
