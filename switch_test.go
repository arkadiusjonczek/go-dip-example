package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Switch_On_Successful(t *testing.T) {
	s := NewSwitch(NewLamp())
	err := s.On()

	require.NoError(t, err)
}

func Test_Switch_Off_Successful(t *testing.T) {
	s := NewSwitch(NewLamp())
	err := s.On()

	require.NoError(t, err)

	err = s.Off()
	require.NoError(t, err)
}

func Test_Switch_On_ReturnsErrorOnActiveSwitch(t *testing.T) {
	s := NewSwitch(NewLamp())
	err := s.On()

	require.NoError(t, err)

	err = s.On()
	require.Error(t, err)
}

func Test_Switch_Off_ReturnsErrorOnInactiveSwitch(t *testing.T) {
	s := NewSwitch(NewLamp())
	err := s.Off()

	require.Error(t, err)
}
