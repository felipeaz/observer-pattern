package light

import "time"

type Signal interface {
	GetColor() string
	SwitchTime() time.Duration
}

type signal struct {
	color      string
	switchTime time.Duration
}

func NewSignal() Signal {
	return nil
}
