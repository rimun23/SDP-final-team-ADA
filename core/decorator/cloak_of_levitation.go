package decorator

import (
	"SDP-final-team-ADA/core"
	"fmt"
)

type CloakOfLevitationDecorator struct {
	Spell core.Castable
}

func WrapWithCloakOfLevitation(spell core.Castable) *CloakOfLevitationDecorator {
	return &CloakOfLevitationDecorator{Spell: spell}
}

func (d *CloakOfLevitationDecorator) Cast() {
	fmt.Println(" [CLOAK OF LEVITATION]: Lifting the Sorcerer into the air, dodging incoming attacks!")
	d.Spell.Cast()
}

func (d *CloakOfLevitationDecorator) GetManaCost() int {
	cost := d.Spell.GetManaCost() - 5
	if cost < 5 {
		return 5
	}
	return cost
}

func (d *CloakOfLevitationDecorator) GetPower() int { return d.Spell.GetPower() + 15 }
func (d *CloakOfLevitationDecorator) GetDescription() string {
	return d.Spell.GetDescription() + " + [Cloak of Levitation (Mobility)]"
}
