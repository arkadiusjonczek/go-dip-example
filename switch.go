package main

import (
	"fmt"
	"log"
)

type SwitchPortInterface interface {
	On() error
	Off() error
}

type Switch struct {
	on   bool
	port SwitchPortInterface
}

func NewSwitch(port SwitchPortInterface) *Switch {
	return &Switch{
		on:   false,
		port: port,
	}
}

func (s *Switch) On() error {
	if s.on == true {
		log.Println("the Switch is already on")
		return fmt.Errorf("the Switch is already on")
	}

	err := s.port.On()
	if err != nil {
		return err
	}

	s.on = true

	return nil
}

func (s *Switch) Off() error {
	if s.on == false {
		log.Println("the Switch is already off")
		return fmt.Errorf("the Switch is already off")
	}

	err := s.port.Off()
	if err != nil {
		return err
	}

	s.on = false

	return nil
}
