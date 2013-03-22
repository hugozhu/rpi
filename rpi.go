package main

/*
#cgo LDFLAGS: -lwiringPi

#include <wiringPi.h>
#include <stdio.h>
#include <stdlib.h>

*/
import "C"

import (
	"log"
	"time"
)

const (
	OUTPUT = C.OUTPUT
	INPUT  = C.INPUT
	LOW    = C.LOW
	HIGH   = C.HIGH
)

func WiringPiSetup() {
	if -1 == int(C.wiringPiSetup()) {
		panic("Failed to setup Pi")
	}
}

func PinMode(port int, mode int) {
	C._pinMode(C.int(port), C.int(mode))
}

func DigitalWrite(port int, mode int) {
	C._digitalWrite(C.int(port), C.int(mode))
}

func Delay(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}
