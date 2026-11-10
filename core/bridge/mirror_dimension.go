package bridge

import "fmt"

type MirrorDimension struct{}

func NewMirrorDimension() *MirrorDimension {
	return &MirrorDimension{}
}
func (m *MirrorDimension) GetDimensionName() string {
	return "Mirror Dimension"
}
func (m *MirrorDimension) ChannelEnergy(constructName string, power int) string {
	return fmt.Sprintf("[%s] Folding fractal reality around %s (Power: %d). Attacks echo without threatening the real world.", m.GetDimensionName(), constructName, power)
}
func (m *MirrorDimension) ModifyPower(basePower int) int {
	return int(float64(basePower) * 1.5)
}
