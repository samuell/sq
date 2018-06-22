package main

import (
	"fmt"

	"github.com/samuell/sq"
)

func main() {
	kb := sq.NewKB()
	kb.AddFact("saml", "likes", "coffee")
	kb.AddFact("saml", "lives_in", "sweden")

	for res := range kb.Q("saml", "likes", "?") {
		fmt.Println(res)
	}
}
