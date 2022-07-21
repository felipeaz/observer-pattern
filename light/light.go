package light

import (
	uuid "github.com/satori/go.uuid"
	"log"
	"observer-pattern/observer"
)

type light struct {
	changed bool
	signal  Signal
	cars    map[uuid.UUID]observer.Observer
}

func NewObservable(signal Signal) observer.Observable {
	return &light{
		signal: signal,
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

func (l *light) GetSignal() string {
	return l.signal.GetSignal()
}

func (l *light) ClearObservers() {
	if l.signal.GetSignal() == GreenSignal {
		log.Println("Signal was green, all cars moved :)")
		l.cars = make(map[uuid.UUID]observer.Observer)
	}
}
