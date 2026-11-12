package analysis

import (
	"strings"
	"testing"
)

func requireFeature(t *testing.T, features FeatureSet, kind FeatureKind, value string) Feature {
	t.Helper()
	feature, exists := features[string(kind)+"\x00"+value]
	if !exists {
		t.Fatalf("missing %s feature %q; features: %#v", kind, value, features)
	}
	return feature
}

func TestFeatureSetAddMergesCountsAndRetainsDescription(t *testing.T) {
	features := make(FeatureSet)
	features.Add(Feature{Kind: "kind", Value: "value", Search: "first", Message: "first message"})
	features.Add(Feature{Kind: "kind", Value: "value", Search: "second", Message: "second message"})
	features.Add(Feature{Kind: "other", Value: "value", Search: "other", Message: "other message"})

	if len(features) != 2 {
		t.Fatalf("feature count = %d, want 2", len(features))
	}
	feature := requireFeature(t, features, "kind", "value")
	if feature.Count != 2 || feature.Search != "first" || feature.Message != "first message" {
		t.Fatalf("merged feature = %#v", feature)
	}
}

func TestSemanticFeaturesNormalizesInlineAndDisplayMath(t *testing.T) {
	features, err := SemanticFeatures(`<math-renderer>$` + "`x+y`" + `$</math-renderer><x-equation>$$
x+y
$$</x-equation><x-equation>z^2</x-equation>`)
	if err != nil {
		t.Fatal(err)
	}
	if feature := requireFeature(t, features, featureMath, "x+y"); feature.Count != 2 || feature.Search != "x+y" {
		t.Fatalf("normalized math feature = %#v", feature)
	}
	if feature := requireFeature(t, features, featureMath, "z^2"); feature.Count != 1 {
		t.Fatalf("second math feature = %#v", feature)
	}
}

func TestSemanticFeaturesSeparatesEmbeddedMathTextTokens(t *testing.T) {
	features, err := SemanticFeatures(`<p>before $x$ after</p>`)
	if err != nil {
		t.Fatal(err)
	}
	requireFeature(t, features, featureText, "before")
	requireFeature(t, features, featureText, "after")
	if _, exists := features[string(featureText)+"\x00before $x$ after"]; exists {
		t.Fatal("embedded math was retained as one visible text feature")
	}
}

func TestSemanticFeaturesRetainsCurrencyText(t *testing.T) {
	features, err := SemanticFeatures(`<p>cost $5 and $6 total</p>`)
	if err != nil {
		t.Fatal(err)
	}
	requireFeature(t, features, featureText, "cost $5 and $6 total")
}

func TestSemanticFeaturesExtractsTaskStates(t *testing.T) {
	features, err := SemanticFeatures(`<ul><li><input type="checkbox"> open</li><li><input checked type="CHECKBOX"> done</li></ul>`)
	if err != nil {
		t.Fatal(err)
	}
	unchecked := requireFeature(t, features, featureTaskState, "unchecked")
	if unchecked.Search != "[ ]" || unchecked.Count != 1 {
		t.Fatalf("unchecked task feature = %#v", unchecked)
	}
	checked := requireFeature(t, features, featureTaskState, "checked")
	if checked.Search != "[x]" || checked.Count != 1 || !strings.Contains(checked.Message, "checked task state") {
		t.Fatalf("checked task feature = %#v", checked)
	}
}

func TestSemanticFeaturesExtractsTableDimensions(t *testing.T) {
	features, err := SemanticFeatures(`<table><thead><tr><th>Name</th><th>Value</th></tr></thead><tbody><tr><td>first</td><td>2</td><td>extra</td></tr></tbody></table>`)
	if err != nil {
		t.Fatal(err)
	}
	table := requireFeature(t, features, featureStructure, "table:2:3")
	if table.Search != "Name" || !strings.Contains(table.Message, "2 rows x 3 columns") {
		t.Fatalf("table feature = %#v", table)
	}
}

func TestSemanticFeaturesExtractsGenericContent(t *testing.T) {
	features, err := SemanticFeatures(`<p>visible <strong>bold words</strong></p><pre><code>block()</code></pre><code>inline()</code><a href="https://example.com">label</a><img src="pic.png" alt="picture">`)
	if err != nil {
		t.Fatal(err)
	}
	requireFeature(t, features, featureText, "visible bold words")
	requireFeature(t, features, featureFormatting, "strong")
	requireFeature(t, features, featureCode, "preformatted\x00block()")
	requireFeature(t, features, featureCode, "inline\x00inline()")
	requireFeature(t, features, featureLink, "https://example.com")
	requireFeature(t, features, featureImage, "pic.png\x00picture")
}
