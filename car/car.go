package car

import (
	uuid "github.com/satori/go.uuid"
	"log"
	"observer-pattern/observer"
)

type car struct {
	id          uuid.UUID
	LightSignal observer.Observable
}

func NewObserver() observer.Observer {
	return &car{
		id: uuid.NewV4(),
	}
}

func (c *car) Update() {
	log.Printf("Car ID %s Received the Light Signal - %s\n", c.GetID().String(), c.LightSignal.GetSignal())
}

func (c *car) RegisterObservable(lightSignal observer.Observable) {
	if c.LightSignal == lightSignal {
		// car is already listening to that signal
		return
	}
	if c.LightSignal != nil {
		c.LightSignal.RemoveObserver(c)
	}
	c.LightSignal = lightSignal
}

func (c *car) RemoveObservable() {
	if c.LightSignal == nil {
		return
	}
	c.LightSignal.RemoveObserver(c)
}

func (c *car) GetID() uuid.UUID {
	return c.id
}
