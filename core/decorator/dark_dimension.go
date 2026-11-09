package decorator

import (
	"SDP-final-team-ADA/core"
	"fmt"
)

type DarkDimensionTaintDecorator struct {
	Spell core.Castable
}

func WrapWithDarkDimensionTaint(spell core.Castable) *DarkDimensionTaintDecorator {
	return &DarkDimensionTaintDecorator{Spell: spell}
}

func (d *DarkDimensionTaintDecorator) Cast() {
	d.Spell.Cast()
	fmt.Println(" [DARK DIMENSION TAINT]: Dark purple necrotic flame bursts! Power is doubled.")
}

func (d *DarkDimensionTaintDecorator) GetManaCost() int { return d.Spell.GetManaCost() + 50 }
func (d *DarkDimensionTaintDecorator) GetPower() int    { return d.Spell.GetPower() * 2 }
func (d *DarkDimensionTaintDecorator) GetDescription() string {
	return d.Spell.GetDescription() + " + [CORRUPTED: Dark Dimension (x2 Power)]"
}
