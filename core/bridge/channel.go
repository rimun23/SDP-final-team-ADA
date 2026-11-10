package bridge

type DimensionChannel interface {
	GetDimensionName() string
	ChannelEnergy(constructName string, power int) string
	ModifyPower(basePower int) int
}
