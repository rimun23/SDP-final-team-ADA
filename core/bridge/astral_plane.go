package bridge

import "fmt"

type AstralPlane struct{}

func NewAstralPlane() *AstralPlane {
	return &AstralPlane{}
}
func (a *AstralPlane) GetDimensionName() string {
	return "Astral Plane"
}
func (a *AstralPlane) ChannelEnergy(constructName string, power int) string {
	return fmt.Sprintf("[%s] Projecting %s directly into the ethereal spirit realm (Power: %d). Bypasses mundane armor.", a.GetDimensionName(), constructName, power)
}
func (a *AstralPlane) ModifyPower(basePower int) int {
	return int(float64(basePower) * 1.25)
}
