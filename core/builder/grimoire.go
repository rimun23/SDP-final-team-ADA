package builder

import "SDP-final-team-ADA/core"

type AncientOneGrimoire struct{}

func NewAncientOneGrimoire() *AncientOneGrimoire {
	return &AncientOneGrimoire{}
}

func (g *AncientOneGrimoire) CreateCrimsonBandsOfCyttorak() *Spell {
	return NewSpellBuilder("Crimson Bands of Cyttorak").
		WithIncantation("By the crimson bands of Cyttorak, bind the enemy!").
		WithRuneCircle("Bands of Cyttorak Glyphs").
		WithElement("Binding Crimson Plasma").
		WithManaCost(45).
		WithPower(120).
		WithTargetType(core.SingleTarget).
		Build()
}

func (g *AncientOneGrimoire) CreateShieldOfSeraphim() *Spell {
	return NewSpellBuilder("Shield of the Seraphim").
		WithIncantation("Seraphic golden light, shield us from destruction!").
		WithRuneCircle("Hexagonal Gold Geometry").
		WithElement("Solar Divine Barrier").
		WithManaCost(30).
		WithPower(85).
		WithTargetType(core.SelfBuff).
		Build()
}
