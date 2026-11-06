package core

type Castable interface {
	Cast()
	GetManaCost() int
	GetPower() int
	GetDescription() string
}
