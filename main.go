package main

import (
	"observer-pattern/car"
	"observer-pattern/light"
)

func main() {
	trafficLight := light.NewObservable()
	car1 := car.NewObserver()

	trafficLight.SetChanged()
	trafficLight.RegisterObserver(car1)
	trafficLight.NotifyObservers()
}
