package main

import (
	"SDP-final-team-ADA/core"
	"SDP-final-team-ADA/core/bridge"
	"SDP-final-team-ADA/core/builder"
	"SDP-final-team-ADA/core/decorator"
	"SDP-final-team-ADA/core/strategy"
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
	target := core.NewEnemy("Kaecilius Zealot", 1200, "High")
	target.TakeDamage(ultimateSpell.GetPower())
	fmt.Println()
	fmt.Println("bridge pattern")
	materialDim := bridge.NewMaterialDimension()
	mirrorDim := bridge.NewMirrorDimension()
	astralDim := bridge.NewAstralPlane()
	whip := bridge.NewEldritchWhip(materialDim)
	fmt.Println(">> 1. Eldritch Whip in Material Dimension:")
	fmt.Println("  ", whip.Activate())
	fmt.Printf("   Effective Power: %d (Base: %d)\n\n", whip.GetEffectivePower(), whip.GetBasePower())
	fmt.Println(">> 2. Switching Eldritch Whip channel to Mirror Dimension:")
	whip.SetChannel(mirrorDim)
	fmt.Println("  ", whip.Activate())
	fmt.Printf("   Effective Power: %d (Amplified by Mirror Geometry x1.5)\n\n", whip.GetEffectivePower())
	fmt.Println(">> 3. Switching Eldritch Whip channel to Astral Plane:")
	whip.SetChannel(astralDim)
	fmt.Println("  ", whip.Activate())
	fmt.Printf("   Effective Power: %d (Soul Damage x1.25)\n\n", whip.GetEffectivePower())
	shield := bridge.NewSeraphicShield(mirrorDim)
	fmt.Println(">> 4. Seraphic Shield in Mirror Dimension:")
	fmt.Println("  ", shield.Activate())
	fmt.Printf("   Effective Defense: %d\n\n", shield.GetEffectivePower())
	fmt.Println("strategy pattern")
	fmt.Println(">> Tactic 1: Isolating enemy with Mirror Trap Containment Strategy")
	combatCtx := strategy.NewCombatContext(strategy.NewMirrorTrapStrategy())
	res1 := combatCtx.ExecuteStrategy(crimsonBands, target)
	fmt.Println("  Combat Result:", res1)
	fmt.Printf("   -> Enemy defeated: %t (HP remaining: %d)\n\n", target.IsDefeated(), target.Health)
	fmt.Println(">> Tactic 2: Switching to aggressive Eldritch Assault Strategy")
	combatCtx.SetStrategy(strategy.NewEldritchAssaultStrategy())
	res2 := combatCtx.ExecuteStrategy(crimsonBands, target)
	fmt.Println("  Combat Result:", res2)
	fmt.Printf("   -> Enemy defeated: %t (HP remaining: %d)\n\n", target.IsDefeated(), target.Health)
	fmt.Println(">> Tactic 3: Switching to Temporal Loop (Time Loop) Strategy")
	combatCtx.SetStrategy(strategy.NewTimeLoopStrategy())
	res3 := combatCtx.ExecuteStrategy(crimsonBands, target)
	fmt.Println("  Combat Result:", res3)
	fmt.Printf("   -> Enemy defeated: %t (HP remaining: %d)\n\n", target.IsDefeated(), target.Health)
	if !target.IsDefeated() {
		fmt.Println(">> Finishing blow with Temporal Loop:")
		combatCtx.SetStrategy(strategy.NewEldritchAssaultStrategy())
		round := 1
		for !target.IsDefeated() {
			fmt.Printf("\nFinishing Round%d\n", round)
			result := combatCtx.ExecuteStrategy(crimsonBands, target)
			fmt.Println(" ", result)
			round++
		}
		fmt.Printf("   -> Final status: Enemy defeated = %t (HP remaining: %d)\n", target.IsDefeated(), target.Health)
	}
	fmt.Println("Sanctum Sanctorum successfully defended!")
}
