package main

import (
	"SDP-final-team-ADA/core"
	"SDP-final-team-ADA/core/builder"
	"SDP-final-team-ADA/core/decorator"
	"fmt"
)

func main() {
	fmt.Println("-***- Builder pattern -***-")
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

	fmt.Println("-***- Decorator pattern -***-")
	fmt.Println("Example 1: Single Decorator: Eye of Agamotto")
	timeEnchantedBands := decorator.WrapWithEyeOfAgamotto(crimsonBands)
	timeEnchantedBands.Cast()
	fmt.Printf("Result Pipeline: %s\n", timeEnchantedBands.GetDescription())
	fmt.Printf("Stats: Power = %d | Mana = %d\n\n", timeEnchantedBands.GetPower(), timeEnchantedBands.GetManaCost())

	fmt.Println("Example 2: Stacked Decorators (Cloak + Eye + Dark Dimension)")
	ultimateSpell := decorator.WrapWithDarkDimensionTaint(
		decorator.WrapWithEyeOfAgamotto(
			decorator.WrapWithCloakOfLevitation(customSpell),
		),
	)
	fmt.Printf("Full Pipeline: %s\n", ultimateSpell.GetDescription())
	fmt.Printf("Calculated: Total Power = %d | Total Mana = %d\n", ultimateSpell.GetPower(), ultimateSpell.GetManaCost())
	ultimateSpell.Cast()
	fmt.Println("\nCombat with enemy")
	target := core.NewEnemy("Kaecilius Zealot", 250, "High")
	target.TakeDamage(ultimateSpell.GetPower())

}
