package sq

import (
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
	kb.AddFact(Triple{"saml", "likes", "coffee"})

	results := kb.Q("saml", "likes", "?")
	for result := range results {
		expected := "coffee"
		if result.O != expected {
			t.Errorf("Wrong result in TestSimpleQuery. Exepcted %s but found %s", expected, result.O)
		}
	}
}
