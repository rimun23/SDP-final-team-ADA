package core

import "fmt"

type Enemy struct {
	Name        string
	Health      int
	ThreatLevel string
}

func NewEnemy(name string, health int, threatLevel string) *Enemy {
	return &Enemy{Name: name, Health: health, ThreatLevel: threatLevel}
}

func (e *Enemy) TakeDamage(damage int) {
	e.Health -= damage
	if e.Health < 0 {
		e.Health = 0
	}
	fmt.Printf("   Target '%s' took %d dmg! Remaining HP: %d\n", e.Name, damage, e.Health)
}
func (e *Enemy) IsDefeated() bool {
	return e.Health <= 0
}
