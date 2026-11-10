package bridge

import "fmt"

type SeraphicShield struct {
	name      string
	basePower int
	channel   DimensionChannel
}

func NewSeraphicShield(channel DimensionChannel) *SeraphicShield {
	return &SeraphicShield{
		name:      "Shield of the Seraphim",
		basePower: 90,
		channel:   channel,
	}
}
func (s *SeraphicShield) SetChannel(channel DimensionChannel) {
	s.channel = channel
}
func (s *SeraphicShield) GetChannel() DimensionChannel {
	return s.channel
}
func (s *SeraphicShield) GetName() string {
	return s.name
}
func (s *SeraphicShield) GetBasePower() int {
	return s.basePower
}
func (s *SeraphicShield) GetEffectivePower() int {
	if s.channel == nil {
		return s.basePower
	}
	return s.channel.ModifyPower(s.basePower)
}
func (s *SeraphicShield) Activate() string {
	effPower := s.GetEffectivePower()
	if s.channel == nil {
		return fmt.Sprintf("Deploying unchanneled %s with base defense %d", s.name, effPower)
	}
	return s.channel.ChannelEnergy(s.name, effPower)
}
