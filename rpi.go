package main

/*
#cgo LDFLAGS: -lwiringPi

#include <wiringPi.h>
#include <stdio.h>
#include <stdlib.h>

static void my_pinMode(int p, int m) {
    pinMode(p,m);
}

static void my_digitalWrite(int p, int m) {
    digitalWrite(p,m);
}
*/
import "C"

import (
	"time"
)

const (
	WPI_MODE_PINS          = C.WPI_MODE_PINS
	WPI_MODE_GPIO          = C.WPI_MODE_GPIO
	WPI_MODE_GPIO_SYS      = C.WPI_MODE_GPIO_SYS
	WPI_MODE_PIFACE        = C.WPI_MODE_PIFACE
	WPI_MODE_UNINITIALISED = C.WPI_MODE_UNINITIALISED

	OUTPUT     = C.OUTPUT
	INPUT      = C.INPUT
	PWM_OUTPUT = C.PWM_OUTPUT
	GPIO_CLOCK = C.GPIO_CLOCK

	LOW  = C.LOW
	HIGH = C.HIGH

	PUD_OFF  = C.PUD_OFF
	PUD_DOWN = C.PUD_DOWN
	PUD_UP   = C.PUD_UP

	// PWM

	PWM_MODE_MS  = C.PWM_MODE_MS
	PWM_MODE_BAL = C.PWM_MODE_BAL

	INT_EDGE_SETUP   = C.INT_EDGE_SETUP
	INT_EDGE_FALLING = C.INT_EDGE_FALLING
	INT_EDGE_RISING  = C.INT_EDGE_RISING
	INT_EDGE_BOTH    = C.INT_EDGE_BOTH
)

func WiringPiSetup() {
	if -1 == int(C.wiringPiSetup()) {
		panic("Failed to setup Pi")
	}
}

func PinMode(port int, mode int) {
	C.my_pinMode(C.int(port), C.int(mode))
}

func DigitalWrite(port int, mode int) {
	C.my_digitalWrite(C.int(port), C.int(mode))
}

func Delay(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}
