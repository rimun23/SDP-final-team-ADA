package strategy

import "SDP-final-team-ADA/core"

type CombatStrategy interface {
	GetName() string
	Execute(spell core.Castable, target *core.Enemy) string
}
type CombatContext struct {
	strategy CombatStrategy
}

func NewCombatContext(initialStrategy CombatStrategy) *CombatContext {
	return &CombatContext{strategy: initialStrategy}
}
func (c *CombatContext) SetStrategy(s CombatStrategy) {
	c.strategy = s
}
func (c *CombatContext) GetStrategy() CombatStrategy {
	return c.strategy
}
func (c *CombatContext) ExecuteStrategy(spell core.Castable, target *core.Enemy) string {
	if c.strategy == nil {
		return "No combat strategy selected!"
	}
	return c.strategy.Execute(spell, target)
}
