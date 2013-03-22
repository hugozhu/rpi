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
	C.my_pinMode(C.int(port), C.int(mode))
}

func DigitalWrite(port int, mode int) {
	C.my_digitalWrite(C.int(port), C.int(mode))
}

func Delay(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}
