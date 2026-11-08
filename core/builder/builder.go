package builder

import "SDP-final-team-ADA/core"

type SpellBuilder struct {
	name              string
	incantation       string
	runeCircle        string
	element           string
	manaCost          int
	power             int
	targetType        core.TargetType
	requiresSlingRing bool
}

func NewSpellBuilder(name string) *SpellBuilder {
	return &SpellBuilder{
		name:       name,
		runeCircle: "Standard Eldritch Mandala",
		element:    "Pure Eldritch Sparks",
		manaCost:   15,
		power:      25,
		targetType: core.SingleTarget,
	}
}

func (b *SpellBuilder) WithIncantation(inc string) *SpellBuilder        { b.incantation = inc; return b }
func (b *SpellBuilder) WithRuneCircle(circle string) *SpellBuilder      { b.runeCircle = circle; return b }
func (b *SpellBuilder) WithElement(element string) *SpellBuilder        { b.element = element; return b }
func (b *SpellBuilder) WithManaCost(cost int) *SpellBuilder             { b.manaCost = cost; return b }
func (b *SpellBuilder) WithPower(power int) *SpellBuilder               { b.power = power; return b }
func (b *SpellBuilder) WithTargetType(tt core.TargetType) *SpellBuilder { b.targetType = tt; return b }
func (b *SpellBuilder) RequiresSlingRing(req bool) *SpellBuilder        { b.requiresSlingRing = req; return b }

func (b *SpellBuilder) Build() *Spell {
	return &Spell{
		name:              b.name,
		incantation:       b.incantation,
		runeCircle:        b.runeCircle,
		element:           b.element,
		manaCost:          b.manaCost,
		power:             b.power,
		targetType:        b.targetType,
		requiresSlingRing: b.requiresSlingRing,
	}
}
