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
	log.Println("Received a message from LightSignal")
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

func (c *car) GetID() uuid.UUID {
	return c.id
}
