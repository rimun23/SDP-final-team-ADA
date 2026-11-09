package facade

import (
	"fmt"

	"SDP-final-team-ADA/core/adapter"
)

type Enemy struct {
	Name string
	HP   int
}

type SpellFactory interface {
	BuildSpell(name string) adapter.CastableSpell
}

type Enhancer interface {
	Enhance(s adapter.CastableSpell, relics ...string) adapter.CastableSpell
}

type DimensionProjector interface {
	Project(s adapter.CastableSpell, dimension string) string
}

type CombatStrategy interface {
	ExecuteTactics(s adapter.CastableSpell, target *Enemy) string
}

type SanctumSanctorumFacade struct {
	factory    SpellFactory
	enhancer   Enhancer
	projector  DimensionProjector
	aggressive CombatStrategy
	mirrorTrap CombatStrategy
	darkhold   adapter.CastableSpell
}

func NewSanctumSanctorumFacade(f SpellFactory, e Enhancer, p DimensionProjector, aggressive, mirrorTrap CombatStrategy) *SanctumSanctorumFacade {
	book := adapter.NewDarkholdGrimoire("Interdimensional Strike")
	return &SanctumSanctorumFacade{
		factory:    f,
		enhancer:   e,
		projector:  p,
		aggressive: aggressive,
		mirrorTrap: mirrorTrap,
		darkhold:   adapter.NewDarkholdAdapter(book),
	}
}

func (s *SanctumSanctorumFacade) DefendSanctumFromThreat(threat *Enemy) {
	fmt.Printf("Threat detected: %s (HP %d)\n", threat.Name, threat.HP)

	spell := s.factory.BuildSpell("Shield of the Seraphim")
	fmt.Println("[1] built:", spell.Name())

	spell = s.enhancer.Enhance(spell, "EyeOfAgamotto", "CloakOfLevitation")
	fmt.Println("[2] enhanced:", spell.Name())

	fmt.Println("[3]", s.projector.Project(spell, "real"))
	fmt.Println("[4]", s.aggressive.ExecuteTactics(spell, threat))
}

func (s *SanctumSanctorumFacade) CastUltimateMirrorTrap(threat *Enemy) {
	fmt.Printf("Luring %s into the Mirror Dimension\n", threat.Name)

	spell := s.factory.BuildSpell("Mirror Wall")
	fmt.Println("[1] built:", spell.Name())

	spell = s.enhancer.Enhance(spell, "EyeOfAgamotto")
	fmt.Println("[2] enhanced:", spell.Name())

	fmt.Println("[3]", s.projector.Project(spell, "mirror"))
	fmt.Println("[4]", s.mirrorTrap.ExecuteTactics(spell, threat))
}

func (s *SanctumSanctorumFacade) InvokeForbiddenInterdimensionalStrike(threat *Enemy) {
	fmt.Printf("Forbidden strike against %s\n", threat.Name)

	spell := s.darkhold
	fmt.Println("[1]", spell.Cast(50))

	spell = s.enhancer.Enhance(spell, "DarkDimensionTaint")
	fmt.Println("[2] enhanced:", spell.Name())

	fmt.Println("[3]", s.projector.Project(spell, "astral"))
	fmt.Println("[4]", s.aggressive.ExecuteTactics(spell, threat))
}
