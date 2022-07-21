# observer-pattern
Implementation of the Observer Pattern in Golang

Observer is a behavioral design pattern that lets you define a subscription mechanism to notify multiple objects about
any events that happen to the object they’re observing.

This example simulates a busy traffic where cars keep looking at the traffic light. The traffic light sends signal
to every car that is listening to it, when it is green, all car moves. When it's red, the cars have to stop and listen
to the traffic light.