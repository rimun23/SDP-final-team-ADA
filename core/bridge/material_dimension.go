package bridge

import "fmt"

type MaterialDimension struct{}

func NewMaterialDimension() *MaterialDimension {
	return &MaterialDimension{}
}
func (m *MaterialDimension) GetDimensionName() string {
	return "Material Dimension"
}
func (m *MaterialDimension) ChannelEnergy(constructName string, power int) string {
	return fmt.Sprintf("[%s] Materializing %s with kinetic eldritch force (Power: %d). Physical collateral applies.", m.GetDimensionName(), constructName, power)
}
func (m *MaterialDimension) ModifyPower(basePower int) int {
	return basePower
}
