package sq

import (
	"fmt"
	"testing"
)

func TestNewKB(t *testing.T) {
	kb := NewKB()
	if kb.facts == nil {
		t.Error("KB.facts is not initialized")
	}
	if kb.patternRules == nil {
		t.Error("KB.patternRules is not initialized")
	}
	if kb.goRules == nil {
		t.Error("KB.goRule is not initialized")
	}
}

func TestSimpleQuery(t *testing.T) {
	kb := NewKB()
	kb.AddFact("saml", "likes", "coffee")

	i := 0
	results := kb.Q("saml", "likes", "?")
	for result := range results {
		expected := "coffee"
		if result.O != expected {
			t.Errorf("Wrong result in TestSimpleQuery. Exepcted %s but found %s", expected, result.O)
		}
		i++
	}
	if i == 0 {
		t.Errorf("No results returned from query!")
	}
}

func ExampleQuery() {
	kb := NewKB()
	kb.AddFact("saml", "likes", "coffee")
	kb.AddFact("saml", "lives_in", "sweden")
	for res := range kb.Q("saml", "likes", "?") {
		fmt.Println(res)
	}
	// Output:
	// {saml likes coffee}
}

func TestPatternRule(t *testing.T) {
	kb := NewKB()
	kb.AddFact("tea", "tastes", "soso")
	kb.AddFact("coffee", "tastes", "great")
	kb.AddFact("water", "tastes", "great")
	kb.AddFact("coffee", "contains", "caffeine")
	kb.AddFact("tea", "contains", "caffeine")

	kb.AddPatternRule(Triple{"saml", "likes", "?o"},
		Triple{"?o", "tastes", "great"},
		Triple{"?o", "contains", "caffeine"})

	expected := Triple{"saml", "likes", "coffee"}
	i := 0
	for tr := range kb.Q("saml", "likes", "?") {
		if tr != expected {
			t.Errorf("Expected %v but found %v\n", expected, tr)
		}
		i++
	}
	if i == 0 {
		t.Errorf("No results returned from query!")
	}
}
