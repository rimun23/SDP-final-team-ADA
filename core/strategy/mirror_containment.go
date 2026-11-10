package strategy

import (
	"SDP-final-team-ADA/core"
	"fmt"
)

type MirrorContainmentStrategy struct{}
type MirrorTrapStrategy = MirrorContainmentStrategy

func NewMirrorContainmentStrategy() *MirrorContainmentStrategy {
	return &MirrorContainmentStrategy{}
}
func NewMirrorTrapStrategy() *MirrorTrapStrategy {
	return NewMirrorContainmentStrategy()
}
func (s *MirrorContainmentStrategy) GetName() string {
	return "Mirror Trap Containment"
}
func (s *MirrorContainmentStrategy) Execute(spell core.Castable, target *core.Enemy) string {
	spell.Cast()
	damage := spell.GetPower()
	target.TakeDamage(damage)
	return fmt.Sprintf("[TACTIC: %s] Trapped %s in kaleidoscopic mirror prisms! Power applied: %d. Threat isolated.",
		s.GetName(), target.Name, damage)
}
