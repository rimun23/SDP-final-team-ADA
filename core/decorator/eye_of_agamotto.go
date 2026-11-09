package decorator

import (
	"SDP-final-team-ADA/core"
	"fmt"
)

type EyeOfAgamottoDecorator struct {
	Spell core.Castable
}

func WrapWithEyeOfAgamotto(spell core.Castable) *EyeOfAgamottoDecorator {
	return &EyeOfAgamottoDecorator{Spell: spell}
}

func (d *EyeOfAgamottoDecorator) Cast() {
	d.Spell.Cast()
	fmt.Println(" [EYE OF AGAMOTTO]: Green time runes ignite! Time dilates and slows down enemy reflexes.")
}

func (d *EyeOfAgamottoDecorator) GetManaCost() int { return d.Spell.GetManaCost() + 25 }
func (d *EyeOfAgamottoDecorator) GetPower() int    { return d.Spell.GetPower() + 45 }
func (d *EyeOfAgamottoDecorator) GetDescription() string {
	return d.Spell.GetDescription() + " + [Eye of Agamotto (Time Dilation)]"
}
