package light

import (
	"log"
	"time"
)

const (
	RedSignal    = "RED"
	YellowSignal = "YELLOW"
	GreenSignal  = "GREEN"
)

type Signal interface {
	GetSignal() string
	SetNextSignal()
	Run(signalCh chan<- string)
}

func NewSignal() Signal {
	return &signal{
		signalMap:         getSignalMap(),
		signalDurationMap: getSignalDurationMap(),
		current:           RedSignal,
	}
}

type signal struct {
	signalMap         map[string]string
	signalDurationMap map[string]time.Duration
	current           string
}

func (s *signal) Run(signalCh chan<- string) {
	select {
	case <-time.After(s.GetSignalDuration()):
		s.SetNextSignal()
		signalCh <- s.GetSignal()
		s.Run(signalCh)
	}
}

func (s *signal) GetSignal() string {
	return s.current
}

func (s *signal) SetNextSignal() {
	nextSignal, ok := s.signalMap[s.GetSignal()]
	if !ok {
		log.Fatal("signal not found in map")
	}
	s.current = nextSignal
}

func (s *signal) GetSignalDuration() time.Duration {
	signalDuration, ok := s.signalDurationMap[s.GetSignal()]
	if !ok {
		log.Fatal("signal not mapped")
	}
	return signalDuration
}

func getSignalMap() map[string]string {
	return map[string]string{
		RedSignal:    GreenSignal,
		GreenSignal:  YellowSignal,
		YellowSignal: RedSignal,
	}
}

func getSignalDurationMap() map[string]time.Duration {
	return map[string]time.Duration{
		RedSignal:    time.Second * 5,
		GreenSignal:  time.Second * 5,
		YellowSignal: time.Second * 2,
	}
}
