package builder

import (
	"SDP-final-team-ADA/core"
	"fmt"
)

type Spell struct {
	name              string
	incantation       string
	runeCircle        string
	element           string
	manaCost          int
	power             int
	targetType        core.TargetType
	requiresSlingRing bool
}

func (s *Spell) Cast() {
	fmt.Printf("CASTING SPELL«%s»!\n", s.name)
	if s.incantation != "" {
		fmt.Printf("  Incantation: \"%s\"\n", s.incantation)
	}
	fmt.Printf("   Runic Circle: '%s' Element: %s\n", s.runeCircle, s.element)
	fmt.Printf("   Target: %s, Sling Ring: %t\n", s.targetType, s.requiresSlingRing)
	fmt.Printf("   Base Power: %d, Mana: %d\n", s.power, s.manaCost)
}

func (s *Spell) GetManaCost() int { return s.manaCost }
func (s *Spell) GetPower() int    { return s.power }
func (s *Spell) GetDescription() string {
	return fmt.Sprintf("%s (%s, Power: %d)", s.name, s.element, s.power)
}
func (s *Spell) GetName() string { return s.name }
