package strategy

import (
	"SDP-final-team-ADA/core"
	"fmt"
)

type EldritchAssaultStrategy struct{}

func NewEldritchAssaultStrategy() *EldritchAssaultStrategy {
	return &EldritchAssaultStrategy{}
}
func (s *EldritchAssaultStrategy) GetName() string {
	return "Eldritch Assault"
}
func (s *EldritchAssaultStrategy) Execute(spell core.Castable, target *core.Enemy) string {
	spell.Cast()
	damage := int(float64(spell.GetPower()) * 1.2)
	target.TakeDamage(damage)
	return fmt.Sprintf("[TACTIC:%s] Overwhelming offensive strike with '%s' dealt %d damage to %s!",
		s.GetName(), spell.GetDescription(), damage, target.Name)
}
