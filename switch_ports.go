package main

import "fmt"

var _ SwitchPortInterface = (*Lamp)(nil)

type Lamp struct {
}

func NewLamp() *Lamp {
	return &Lamp{}
}

func (l *Lamp) On() error {
	fmt.Println("Lamp On")

	return nil
}

func (l *Lamp) Off() error {
	fmt.Println("Lamp Off")

	return nil
}

var _ SwitchPortInterface = (*CoffeeMaschine)(nil)

type CoffeeMaschine struct {
}

func NewCoffeeMaschine() *CoffeeMaschine {
	return &CoffeeMaschine{}
}

func (l CoffeeMaschine) On() error {
	fmt.Println("CoffeeMaschine On")

	return nil
}

func (l CoffeeMaschine) Off() error {
	fmt.Println("CoffeeMaschine Off")

	return nil
}
