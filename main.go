package main

import (
	"log"

	"github.com/AdRoll/baker"
	"github.com/SemanticSugar/bakerplay/bakerstuff"
)

func main() {
	components := baker.Components{
		Inputs:  []baker.InputDesc{bakerstuff.StreamDesc},
		Outputs: []baker.OutputDesc{bakerstuff.EchoDesc},
	}

	if err := baker.MainCLI(components); err != nil {
		log.Fatalf("err main: %+v", err)
	}
}
