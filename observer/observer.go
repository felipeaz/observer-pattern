package observer

import uuid "github.com/satori/go.uuid"

type Observer interface {
	Update()
	RegisterObservable(o Observable)

	GetID() uuid.UUID
}
