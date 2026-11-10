package strategy

import (
	"SDP-final-team-ADA/core"
	"fmt"
)

type TemporalLoopStrategy struct{}
type TimeLoopStrategy = TemporalLoopStrategy

func NewTemporalLoopStrategy() *TemporalLoopStrategy {
	return &TemporalLoopStrategy{}
}
func NewTimeLoopStrategy() *TimeLoopStrategy {
	return NewTemporalLoopStrategy()
}
func (s *TemporalLoopStrategy) GetName() string {
	return "Temporal Loop (Time Stone Bargain)"
}
func (s *TemporalLoopStrategy) Execute(spell core.Castable, target *core.Enemy) string {
	spell.Cast()
	totalDamage := 0
	iterations := 3
	damagePerLoop := spell.GetPower() / 2
	if damagePerLoop <= 0 {
		damagePerLoop = 10
	}
	for i := 1; i <= iterations; i++ {
		target.TakeDamage(damagePerLoop)
		totalDamage += damagePerLoop
	}
	return fmt.Sprintf("[TACTIC: %s] 'Dormammu, I've come to bargain!' %s trapped across %d time loops, taking %d total damage.",
		s.GetName(), target.Name, iterations, totalDamage)
}
