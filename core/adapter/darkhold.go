package adapter

import (
	"SDP-final-team-ADA/core"
	"fmt"
)

var _ core.Castable = (*DarkholdAdapter)(nil)
var _ core.Castable = (*AsgardAdapter)(nil)

type DarkholdGrimoire struct {
	ritual string
}

func NewDarkholdGrimoire(ritual string) *DarkholdGrimoire {
	return &DarkholdGrimoire{ritual: ritual}
}

func (d *DarkholdGrimoire) DrainLifeForceAndSummonDemonicFlame(lifeForce int) string {
	return fmt.Sprintf("Darkhold '%s' drained %d life force and summoned demonic flame", d.ritual, lifeForce)
}

func (d *DarkholdGrimoire) FlameIntensity(lifeForce int) int {
	return lifeForce * 12
}

type DarkholdAdapter struct {
	grimoire *DarkholdGrimoire
	mana     int
}

func NewDarkholdAdapter(g *DarkholdGrimoire, mana int) *DarkholdAdapter {
	return &DarkholdAdapter{grimoire: g, mana: mana}
}

func (a *DarkholdAdapter) lifeForce() int {
	life := a.mana / 5
	if life < 1 {
		return 1
	}
	return life
}

func (a *DarkholdAdapter) Cast() {
	fmt.Println(" [DARKHOLD]:", a.grimoire.DrainLifeForceAndSummonDemonicFlame(a.lifeForce()))
}

func (a *DarkholdAdapter) GetManaCost() int { return a.mana }
func (a *DarkholdAdapter) GetPower() int {
	return a.grimoire.FlameIntensity(a.lifeForce())
}
func (a *DarkholdAdapter) GetDescription() string {
	return fmt.Sprintf("Darkhold: %s (Power: %d)", a.grimoire.ritual, a.GetPower())
}

type LokiSeidr struct{}

func (l *LokiSeidr) WeaveIllusion(runes int) string {
	return fmt.Sprintf("Loki weaves an illusion from %d runes", runes)
}

type AsgardAdapter struct {
	seidr *LokiSeidr
	mana  int
}

func NewAsgardAdapter(s *LokiSeidr, mana int) *AsgardAdapter {
	return &AsgardAdapter{seidr: s, mana: mana}
}

func (a *AsgardAdapter) Cast() {
	fmt.Println(" [ASGARD]:", a.seidr.WeaveIllusion(a.mana/10))
}

func (a *AsgardAdapter) GetManaCost() int { return a.mana }
func (a *AsgardAdapter) GetPower() int    { return a.mana / 10 * 6 }
func (a *AsgardAdapter) GetDescription() string {
	return fmt.Sprintf("Asgardian Seidr (Power: %d)", a.GetPower())
}
