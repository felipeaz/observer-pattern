package light

import (
	uuid "github.com/satori/go.uuid"
	"observer-pattern/observer"
)

type light struct {
	changed bool
	signal  Signal
	cars    map[uuid.UUID]observer.Observer
}

func NewObservable() observer.Observable {
	return &light{
		signal: NewSignal(),
		cars:   make(map[uuid.UUID]observer.Observer),
	}
}

func (l *light) RegisterObserver(car observer.Observer) {
	car.RegisterObservable(l)
	l.cars[car.GetID()] = car
}

func (l *light) RemoveObserver(car observer.Observer) {
	delete(l.cars, car.GetID())
}

func (l *light) SetChanged() {
	l.changed = true
}

func (l *light) NotifyObservers() {
	if !l.changed {
		return
	}
	for _, car := range l.cars {
		car.Update()
	}
	l.changed = false
}

func (l *light) GetSignal() Signal {
	return l.signal
}
