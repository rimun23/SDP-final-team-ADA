package core

type TargetType string

const (
	SingleTarget    TargetType = "SINGLE_TARGET"
	AreaOfEffect    TargetType = "AREA_OF_EFFECT"
	SelfBuff        TargetType = "SELF_BUFF"
	DimensionalZone TargetType = "DIMENSIONAL_ZONE"
)
