package main

import "fmt"

func main() {
	cs := make(Compounds)
	go func() {
		defer close(cs)
		cs <- Compound{"Caffeine", "C8H10N4O2"}
		cs <- Compound{"Water", "H2O"}
	}()
	for c := range cs.HasName("Caffeine") {
		fmt.Println(c.Formula)
	}
}

func (cs Compounds) HasName(name string) Compounds {
	outCs := make(Compounds)
	go func() {
		defer close(outCs)
		for c := range cs {
			if c.Name == name {
				outCs <- c
			}
		}
	}()
	return outCs
}

type Compounds chan Compound

type Compound struct {
	Name    string
	Formula string
}
