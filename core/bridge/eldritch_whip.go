package bridge

import "fmt"

type EldritchWhip struct {
	name      string
	basePower int
	channel   DimensionChannel
}

func NewEldritchWhip(channel DimensionChannel) *EldritchWhip {
	return &EldritchWhip{
		name:      "Eldritch Whip",
		basePower: 70,
		channel:   channel,
	}
}
func (w *EldritchWhip) SetChannel(channel DimensionChannel) {
	w.channel = channel
}
func (w *EldritchWhip) GetChannel() DimensionChannel {
	return w.channel
}
func (w *EldritchWhip) GetName() string {
	return w.name
}
func (w *EldritchWhip) GetBasePower() int {
	return w.basePower
}
func (w *EldritchWhip) GetEffectivePower() int {
	if w.channel == nil {
		return w.basePower
	}
	return w.channel.ModifyPower(w.basePower)
}
func (w *EldritchWhip) Activate() string {
	effPower := w.GetEffectivePower()
	if w.channel == nil {
		return fmt.Sprintf("Whipping unchanneled %s with raw power %d", w.name, effPower)
	}
	return w.channel.ChannelEnergy(w.name, effPower)
}
