package main

import (
	"log"
	"math/rand"
	"observer-pattern/car"
	"observer-pattern/light"
	"observer-pattern/observer"
	"time"
)

func main() {
	log.SetFlags(0)

	signalCh := make(chan string)
	signal := light.NewSignal()
	go signal.Run(signalCh)

	trafficLight := light.NewObservable(signal)
	log.Println("Current Traffic Light", trafficLight.GetSignal())

	// Car Traffic Simulator
	go func(trafficLight observer.Observable) {
		for {
			if trafficLight.GetSignal() == light.RedSignal {
				c := car.NewObserver()
				log.Printf("Car %s stopped at the red light", c.GetID().String())
				trafficLight.RegisterObserver(c)
				time.Sleep(time.Second * time.Duration(rand.Intn(5)))
			}
		}
	}(trafficLight)

	for current := range signalCh {
		log.Println("Traffic Light changed to", current)
		trafficLight.SetChanged()
		trafficLight.NotifyObservers()

		if current == light.GreenSignal {
			trafficLight.ClearObservers()
		}
	}
}
