package rpi

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
	"fmt"
	"time"
)

var (
	board2pin = []int{
		-1,
		-1,
		-1,
		8,
		-1,
		9,
		-1,
		7,
		15,
		-1,
		16,
		0,
		1,
		2,
		-1,
		-1,
		4,
		-1,
		5,
		12,
		-1,
		13,
		6,
		14,
		10,
		-1,
		11,
	}
	gpio2pin = []int{
		8,
		9,
		-1,
		-1,
		7,
		-1,
		-1,
		11,
		10,
		13,
		12,
		14,
		-1,
		-1,
		15,
		16,
		-1,
		0,
		1,
		-1,
		-1,
		2,
		3,
		4,
		5,
		6,
		-1,
		-1,
		17,
		18,
		19,
		20,
	}
)

const (
	PIN_GPIO_0  = 0
	PIN_GPIO_1  = 1
	PIN_GPIO_2  = 2
	PIN_GPIO_3  = 3
	PIN_GPIO_4  = 4
	PIN_GPIO_5  = 5
	PIN_GPIO_6  = 6
	PIN_GPIO_7  = 7
	PIN_SDA     = 8
	PIN_SCL     = 9
	PIN_CE0     = 10
	PIN_CE1     = 11
	PIN_MOSI    = 12
	PIN_MOSO    = 13
	PIN_SCLK    = 14
	PIN_TXD     = 15
	PIN_RXD     = 16
	PIN_GPIO_8  = 17
	PIN_GPIO_9  = 18
	PIN_GPIO_10 = 19
	PIN_GPIO_11 = 20

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

//use RPi.GPIO's BOARD numbering
func BoardToPin(pin int) int {
	if pin < 1 || pin >= len(board2pin) {
		panic(fmt.Sprintf("Invalid board pin number: %d", pin))
	}
	return board2pin[pin]
}

func GpioToPin(pin int) int {
	if pin < 0 || pin >= len(gpio2pin) {
		panic(fmt.Sprintf("Invalid bcm gpio number: %d", pin))
	}
	return gpio2pin[pin]
}

func PinToGpio(pin int) int {
	return int(C.wpiPinToGpio(C.int(pin)))
}

func WiringPiSetup() {
	if -1 == int(C.wiringPiSetup()) {
		panic("Failed to setup Pi")
	}
}

// Better to stick to one GPIO numbering, not use other setup methods for now

// func WiringPiSetupGpio() {
// 	if -1 == int(C.wiringPiSetupSys()) {
// 		panic("Failed to setup Pi")
// 	}
// }

// func WiringPiSetupSys() {
// 	if -1 == int(C.wiringPiSetupSys()) {
// 		panic("Failed to setup Pi")
// 	}
// }

// func WiringPiSetupPiFace() {
// 	if -1 == int(C.wiringPiSetupPiFace()) {
// 		panic("Failed to setup Pi")
// 	}
// }

func PinMode(pin int, mode int) {
	C.my_pinMode(C.int(pin), C.int(mode))
}

func DigitalWrite(pin int, mode int) {
	C.my_digitalWrite(C.int(pin), C.int(mode))
}

func Delay(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}
