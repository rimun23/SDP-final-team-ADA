package main

import (
	"SDP-final-team-ADA/core"
	"SDP-final-team-ADA/core/adapter"
	"SDP-final-team-ADA/core/bridge"
	"SDP-final-team-ADA/core/builder"
	"SDP-final-team-ADA/core/decorator"
	"SDP-final-team-ADA/core/facade"
	"SDP-final-team-ADA/core/strategy"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var in = bufio.NewScanner(os.Stdin)

func ask(prompt string) string {
	fmt.Print(prompt)
	if !in.Scan() {
		os.Exit(0)
	}
	return strings.TrimSpace(in.Text())
}

func askInt(prompt string, lo, hi int) int {
	for {
		n, err := strconv.Atoi(ask(prompt))
		if err == nil && n >= lo && n <= hi {
			return n
		}
		fmt.Printf("Enter a number from %d to %d\n", lo, hi)
	}
}

type battle struct {
	enemy    *core.Enemy
	spell    core.Castable
	grimoire *builder.AncientOneGrimoire
	combat   *strategy.CombatContext
	sanctum  *facade.SanctumSanctorumFacade
	whip     *bridge.EldritchWhip
	shield   *bridge.SeraphicShield
	dims     []bridge.DimensionChannel
}

func newBattle() *battle {
	g := builder.NewAncientOneGrimoire()
	dims := []bridge.DimensionChannel{
		bridge.NewMaterialDimension(),
		bridge.NewMirrorDimension(),
		bridge.NewAstralPlane(),
	}
	return &battle{
		enemy:    core.NewEnemy("Dormammu", 2000, "Cosmic"),
		spell:    g.CreateCrimsonBandsOfCyttorak(),
		grimoire: g,
		combat:   strategy.NewCombatContext(strategy.NewEldritchAssaultStrategy()),
		sanctum:  facade.NewSanctumSanctorumFacade(),
		whip:     bridge.NewEldritchWhip(dims[0]),
		shield:   bridge.NewSeraphicShield(dims[0]),
		dims:     dims,
	}
}

func (b *battle) status() {
	fmt.Printf("Enemy: %s | HP: %d | Threat: %s\n", b.enemy.Name, b.enemy.Health, b.enemy.ThreatLevel)
	fmt.Printf("Current spell: %s\n", b.spell.GetDescription())
	fmt.Printf("Power: %d | Mana: %d\n", b.spell.GetPower(), b.spell.GetManaCost())
}

func (b *battle) strike(s core.Castable) {
	s.Cast()
	b.enemy.TakeDamage(s.GetPower())
}

func (b *battle) forge() {
	fmt.Println("\n[BUILDER] Forge a spell")
	fmt.Println("1. Crimson Bands of Cyttorak (preset from AncientOneGrimoire)")
	fmt.Println("2. Shield of the Seraphim (preset from AncientOneGrimoire)")
	fmt.Println("3. Custom spell (SpellBuilder step by step)")
	switch askInt("> ", 1, 3) {
	case 1:
		b.spell = b.grimoire.CreateCrimsonBandsOfCyttorak()
	case 2:
		b.spell = b.grimoire.CreateShieldOfSeraphim()
	case 3:
		b.spell = customSpell()
	}
	b.strike(b.spell)
}

func customSpell() *builder.Spell {
	name := ask("Spell name: ")
	if name == "" {
		name = "Unnamed Spell"
	}
	incantation := ask("Incantation: ")
	power := askInt("Power (1-300): ", 1, 300)
	mana := askInt("Mana cost (1-200): ", 1, 200)
	targets := []core.TargetType{core.SingleTarget, core.AreaOfEffect, core.SelfBuff, core.DimensionalZone}
	fmt.Println("Target: 1. Single  2. Area  3. Self  4. Dimensional zone")
	t := askInt("> ", 1, 4)
	return builder.NewSpellBuilder(name).
		WithIncantation(incantation).
		WithPower(power).
		WithManaCost(mana).
		WithTargetType(targets[t-1]).
		Build()
}

func (b *battle) enchant() {
	fmt.Println("\n[DECORATOR] Enchant the current spell with a relic")
	fmt.Println("1. Eye of Agamotto (+45 power, +25 mana)")
	fmt.Println("2. Cloak of Levitation (+15 power, -5 mana)")
	fmt.Println("3. Dark Dimension Taint (x2 power, +50 mana)")
	switch askInt("> ", 1, 3) {
	case 1:
		b.spell = decorator.WrapWithEyeOfAgamotto(b.spell)
	case 2:
		b.spell = decorator.WrapWithCloakOfLevitation(b.spell)
	case 3:
		b.spell = decorator.WrapWithDarkDimensionTaint(b.spell)
	}
	fmt.Println("Pipeline:", b.spell.GetDescription())
	b.strike(b.spell)
}

func (b *battle) summon() {
	fmt.Println("\n[BRIDGE] Summon a construct in a dimension")
	fmt.Println("1. Eldritch Whip")
	fmt.Println("2. Seraphic Shield")
	c := askInt("> ", 1, 2)
	fmt.Println("Dimension: 1. Material  2. Mirror (x1.5)  3. Astral (x1.25)")
	d := askInt("> ", 1, 3)

	var construct bridge.MysticConstruct = b.whip
	if c == 2 {
		construct = b.shield
	}
	construct.SetChannel(b.dims[d-1])
	fmt.Println(construct.Activate())
	fmt.Printf("Effective power: %d (base %d)\n", construct.GetEffectivePower(), construct.GetBasePower())
	if c == 1 {
		b.enemy.TakeDamage(construct.GetEffectivePower())
	} else {
		fmt.Println("The shield only defends, no damage dealt")
	}
}

func (b *battle) tactic() {
	fmt.Println("\n[STRATEGY] Choose a combat tactic for the current spell")
	fmt.Println("1. Mirror Trap Containment")
	fmt.Println("2. Eldritch Assault (x1.2 damage)")
	fmt.Println("3. Temporal Loop (3 hits)")
	switch askInt("> ", 1, 3) {
	case 1:
		b.combat.SetStrategy(strategy.NewMirrorTrapStrategy())
	case 2:
		b.combat.SetStrategy(strategy.NewEldritchAssaultStrategy())
	case 3:
		b.combat.SetStrategy(strategy.NewTimeLoopStrategy())
	}
	fmt.Println(b.combat.ExecuteStrategy(b.spell, b.enemy))
}

func (b *battle) forbidden() {
	fmt.Println("\n[ADAPTER] Use foreign magic through the Castable interface")
	fmt.Println("1. Darkhold (drains life force)")
	fmt.Println("2. Asgardian Seidr of Loki (runes)")
	c := askInt("> ", 1, 2)
	mana := askInt("Mana to spend (10-100): ", 10, 100)

	var foreign core.Castable
	if c == 1 {
		foreign = adapter.NewDarkholdAdapter(adapter.NewDarkholdGrimoire("Interdimensional Strike"), mana)
	} else {
		foreign = adapter.NewAsgardAdapter(&adapter.LokiSeidr{}, mana)
	}
	b.spell = foreign
	fmt.Println("Now it is a regular spell:", foreign.GetDescription())
	b.strike(foreign)
}

func (b *battle) oneButton() {
	fmt.Println("\n[FACADE] Sanctum Sanctorum control panel")
	fmt.Println("1. Defend the Sanctum")
	fmt.Println("2. Ultimate Mirror Trap")
	fmt.Println("3. Forbidden Interdimensional Strike")
	switch askInt("> ", 1, 3) {
	case 1:
		b.sanctum.DefendSanctumFromThreat(b.enemy)
	case 2:
		b.sanctum.CastUltimateMirrorTrap(b.enemy)
	case 3:
		b.sanctum.InvokeForbiddenInterdimensionalStrike(b.enemy)
	}
}

func main() {
	b := newBattle()
	fmt.Println("Final project about Dr. Strange \n")
	for !b.enemy.IsDefeated() {
		b.status()
		fmt.Println("\n1. Forge a spell        (Builder)")
		fmt.Println("2. Enchant with relic   (Decorator)")
		fmt.Println("3. Summon construct     (Bridge)")
		fmt.Println("4. Choose tactic        (Strategy)")
		fmt.Println("5. Forbidden magic      (Adapter)")
		fmt.Println("6. Sanctum control      (Facade)")
		fmt.Println("0. Retreat")
		switch askInt("\nYour move: ", 0, 6) {
		case 1:
			b.forge()
		case 2:
			b.enchant()
		case 3:
			b.summon()
		case 4:
			b.tactic()
		case 5:
			b.forbidden()
		case 6:
			b.oneButton()
		case 0:
			fmt.Println("The sorcerer retreats to the Sanctum.")
			return
		}
	}
	fmt.Printf("\n%s has been defeated. Sanctum Sanctorum successfully defended!\n", b.enemy.Name)
}
