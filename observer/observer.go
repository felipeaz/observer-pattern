package observer

import uuid "github.com/satori/go.uuid"

type Observer interface {
	Update()
	RegisterObservable(o Observable)
	RemoveObservable()

	GetID() uuid.UUID
}
