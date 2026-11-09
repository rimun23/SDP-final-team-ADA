package adapter

import "fmt"

type CastableSpell interface {
	Name() string
	Cast(mana int) string
}

type DarkholdGrimoire struct {
	ritual string
}

func NewDarkholdGrimoire(ritual string) *DarkholdGrimoire {
	return &DarkholdGrimoire{ritual: ritual}
}

func (d *DarkholdGrimoire) DrainLifeForceAndSummonDemonicFlame(life int) string {
	return fmt.Sprintf("Darkhold (%s): drained %d life force, demonic flame summoned", d.ritual, life)
}

type DarkholdAdapter struct {
	grimoire *DarkholdGrimoire
}

func NewDarkholdAdapter(g *DarkholdGrimoire) *DarkholdAdapter {
	return &DarkholdAdapter{grimoire: g}
}

func (a *DarkholdAdapter) Name() string {
	return "Darkhold " + a.grimoire.ritual
}

func (a *DarkholdAdapter) Cast(mana int) string {
	life := mana / 5
	if life == 0 {
		life = 1
	}
	return a.grimoire.DrainLifeForceAndSummonDemonicFlame(life)
}

type LokiSeidr struct{}

func (l *LokiSeidr) WeaveIllusion(runes int) string {
	return fmt.Sprintf("Loki weaves an illusion from %d runes", runes)
}

type AsgardAdapter struct {
	seidr *LokiSeidr
}

func NewAsgardAdapter(s *LokiSeidr) *AsgardAdapter {
	return &AsgardAdapter{seidr: s}
}

func (a *AsgardAdapter) Name() string {
	return "Asgardian Seidr"
}

func (a *AsgardAdapter) Cast(mana int) string {
	return a.seidr.WeaveIllusion(mana / 10)
}
