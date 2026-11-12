package facade

import (
	"SDP-final-team-ADA/core"
	"SDP-final-team-ADA/core/adapter"
	"SDP-final-team-ADA/core/bridge"
	"SDP-final-team-ADA/core/builder"
	"SDP-final-team-ADA/core/decorator"
	"SDP-final-team-ADA/core/strategy"
	"fmt"
)

type SanctumSanctorumFacade struct {
	grimoire *builder.AncientOneGrimoire
	combat   *strategy.CombatContext
	material bridge.DimensionChannel
	mirror   bridge.DimensionChannel
	astral   bridge.DimensionChannel
}

func NewSanctumSanctorumFacade() *SanctumSanctorumFacade {
	return &SanctumSanctorumFacade{
		grimoire: builder.NewAncientOneGrimoire(),
		combat:   strategy.NewCombatContext(strategy.NewEldritchAssaultStrategy()),
		material: bridge.NewMaterialDimension(),
		mirror:   bridge.NewMirrorDimension(),
		astral:   bridge.NewAstralPlane(),
	}
}

func (f *SanctumSanctorumFacade) DefendSanctumFromThreat(threat *core.Enemy) {
	fmt.Printf("Threat detected: %s (HP %d)\n", threat.Name, threat.Health)

	shield := f.grimoire.CreateShieldOfSeraphim()
	spell := decorator.WrapWithEyeOfAgamotto(decorator.WrapWithCloakOfLevitation(shield))
	fmt.Println("[Builder + Decorator]", spell.GetDescription())

	fmt.Println("[Bridge]", bridge.NewSeraphicShield(f.material).Activate())

	f.combat.SetStrategy(strategy.NewEldritchAssaultStrategy())
	fmt.Println("[Strategy]", f.combat.ExecuteStrategy(spell, threat))
}

func (f *SanctumSanctorumFacade) CastUltimateMirrorTrap(threat *core.Enemy) {
	fmt.Printf("Luring %s into the Mirror Dimension\n", threat.Name)

	bands := f.grimoire.CreateCrimsonBandsOfCyttorak()
	spell := decorator.WrapWithEyeOfAgamotto(bands)
	fmt.Println("[Builder + Decorator]", spell.GetDescription())

	fmt.Println("[Bridge]", bridge.NewEldritchWhip(f.mirror).Activate())

	f.combat.SetStrategy(strategy.NewMirrorTrapStrategy())
	fmt.Println("[Strategy]", f.combat.ExecuteStrategy(spell, threat))
}

func (f *SanctumSanctorumFacade) InvokeForbiddenInterdimensionalStrike(threat *core.Enemy) {
	fmt.Printf("Forbidden strike against %s\n", threat.Name)

	var spell core.Castable = adapter.NewDarkholdAdapter(adapter.NewDarkholdGrimoire("Interdimensional Strike"), 50)
	fmt.Println("[Adapter]", spell.GetDescription())

	spell = decorator.WrapWithDarkDimensionTaint(spell)
	fmt.Println("[Decorator]", spell.GetDescription())

	fmt.Println("[Bridge]", bridge.NewEldritchWhip(f.astral).Activate())

	f.combat.SetStrategy(strategy.NewEldritchAssaultStrategy())
	fmt.Println("[Strategy]", f.combat.ExecuteStrategy(spell, threat))
}
