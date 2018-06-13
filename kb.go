package sq

type Triple struct {
	S string
	P string
	O string
}

type TripleChan chan Triple

func NewKB() *KB {
	return &KB{
		map[Triple]Triple{},
		map[Triple][]Triple{},
		map[Triple]func(kb *KB, tr Triple) TripleChan{},
	}
}

type KB struct {
	facts        map[Triple]Triple
	patternRules map[Triple][]Triple
	goRules      map[Triple]func(kb *KB, tr Triple) TripleChan
}

func (kb *KB) AddFact(s string, p string, o string) {
	tr := Triple{s, p, o}
	kb.facts[tr] = tr
}

func (kb *KB) AddPatternRule(triplePattern Triple, trs ...Triple) {
	kb.patternRules[triplePattern] = trs
}

func (kb *KB) AddGoRule(triplePattern Triple, rule func(kb *KB, tr Triple) TripleChan) {
	kb.goRules[triplePattern] = rule
}

func (kb *KB) Q(s string, p string, o string) TripleChan {
	out := make(TripleChan)
	go func() {
		defer close(out)
		// Check explicit facts
		for tr := range kb.facts {
			if (s == "?" && p == tr.P && o == tr.O) ||
				(s == "?" && p == "?" && o == tr.O) ||
				(s == "?" && p == tr.P && o == "?") ||
				(s == "?" && p == "?" && o == "?") ||
				(s == tr.S && p == "?" && o == tr.O) ||
				(s == tr.S && p == tr.P && o == "?") ||
				(s == tr.S && p == "?" && o == "?") {
				out <- tr
			}
		}
		// TODO: Check rules here
	}()
	return out
}
