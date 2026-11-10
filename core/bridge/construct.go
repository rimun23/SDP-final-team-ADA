package bridge

type MysticConstruct interface {
	Activate() string
	SetChannel(channel DimensionChannel)
	GetChannel() DimensionChannel
	GetName() string
	GetBasePower() int
	GetEffectivePower() int
}
