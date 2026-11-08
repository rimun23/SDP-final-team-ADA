package main

import (
	"SDP-final-team-ADA/core"
	"SDP-final-team-ADA/core/builder"
	"fmt"
)

func main() {
	// inputs for custom spells
	customSpell := builder.NewSpellBuilder("Winds of Watoomb").
		WithIncantation("Let the winds of Watoomb clear the shadows!").
		WithRuneCircle("Twin Cyclone Mandala").
		WithElement("Astral Gale").
		WithManaCost(35).
		WithPower(80).
		WithTargetType(core.AreaOfEffect).
		RequiresSlingRing(false).
		Build()
	customSpell.Cast()
	fmt.Printf("Summary: %s\n\n", customSpell.GetDescription())
	// sample with already ready spell
	grimoire := builder.NewAncientOneGrimoire()
	crimsonBands := grimoire.CreateCrimsonBandsOfCyttorak()
	crimsonBands.Cast()
	fmt.Printf("Summary: %s\n\n", crimsonBands.GetDescription())
}
