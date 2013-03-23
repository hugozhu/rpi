package rpi

import (
	"testing"
)

func Test_hello(t *testing.T) {
	t.Log("hello world")
}

func Test_Setup(t *testing.T) {
	WiringPiSetup()
}

func Test_PinToGpio(t *testing.T) {
	for i := 0; i <= 26; i++ {
		t.Log("PinToGpio", i, PinToGpio(i))
	}
}

func Test_BoardToPin(t *testing.T) {
	for i := 1; i <= 26; i++ {
		t.Log("BoardToPin", i, BoardToPin(i))
	}
}

func Test_GpioToPin(t *testing.T) {
	for i := 0; i <= 31; i++ {
		t.Log("Test_GpioToPin", i, GpioToPin(i))
	}
}
