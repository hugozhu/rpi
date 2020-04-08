package rpi

/*
#cgo LDFLAGS: -lwiringPi

#include <wiringPi.h>
#include <softPwm.h>
#include <stdio.h>
#include <stdlib.h>
#define nil ((void*)0)

#define GEN_INTERRUPTER(PIN) static void interrupt_handler_##PIN() { \
	context ctxt;   \
	ctxt.pin = PIN;  \
	ctxt.ret = PIN;  \
	callback_func(goCallback, &ctxt); \
}

typedef struct context context;
struct context {
	int pin;
	int ret;
};

static void my_pinMode(int p, int m) {
    if (m == PWM_OUTPUT) {
        softPwmCreate(p,0,100);
    }
    else {
        pinMode(p,m);
    }
}

static void my_digitalWrite(int p, int m, int v) {
    if (v < 0) {
        digitalWrite(p,m);
    } else {
        softPwmWrite(p,v);
    }
}

static int my_digitalRead(int p) {
    return digitalRead(p);
}

static void(*callback_func)(void (*f)(void*), void*);

extern void goCallback(void *);

GEN_INTERRUPTER(0)
GEN_INTERRUPTER(1)
GEN_INTERRUPTER(2)
GEN_INTERRUPTER(3)
GEN_INTERRUPTER(4)
GEN_INTERRUPTER(5)
GEN_INTERRUPTER(6)
GEN_INTERRUPTER(7)
GEN_INTERRUPTER(8)
GEN_INTERRUPTER(9)
GEN_INTERRUPTER(10)
GEN_INTERRUPTER(11)
GEN_INTERRUPTER(12)
GEN_INTERRUPTER(13)
GEN_INTERRUPTER(14)
GEN_INTERRUPTER(15)
GEN_INTERRUPTER(16)
GEN_INTERRUPTER(17)
GEN_INTERRUPTER(18)
GEN_INTERRUPTER(19)
GEN_INTERRUPTER(20)

static int my_wiringPiISR(int pin, int mode) {
	switch(pin) {
		case 0: return wiringPiISR(pin, mode, &interrupt_handler_0);
		case 1: return wiringPiISR(pin, mode, &interrupt_handler_1);
		case 2: return wiringPiISR(pin, mode, &interrupt_handler_2);
		case 3: return wiringPiISR(pin, mode, &interrupt_handler_3);
		case 4: return wiringPiISR(pin, mode, &interrupt_handler_4);
		case 5: return wiringPiISR(pin, mode, &interrupt_handler_5);
		case 6: return wiringPiISR(pin, mode, &interrupt_handler_6);
		case 7: return wiringPiISR(pin, mode, &interrupt_handler_7);
		case 8: return wiringPiISR(pin, mode, &interrupt_handler_8);
		case 9: return wiringPiISR(pin, mode, &interrupt_handler_9);
		case 10: return wiringPiISR(pin, mode, &interrupt_handler_10);
		case 11: return wiringPiISR(pin, mode, &interrupt_handler_11);
		case 12: return wiringPiISR(pin, mode, &interrupt_handler_12);
		case 13: return wiringPiISR(pin, mode, &interrupt_handler_13);
		case 14: return wiringPiISR(pin, mode, &interrupt_handler_14);
		case 15: return wiringPiISR(pin, mode, &interrupt_handler_15);
		case 16: return wiringPiISR(pin, mode, &interrupt_handler_16);
		case 17: return wiringPiISR(pin, mode, &interrupt_handler_17);
		case 18: return wiringPiISR(pin, mode, &interrupt_handler_18);
		case 19: return wiringPiISR(pin, mode, &interrupt_handler_19);
		case 20: return wiringPiISR(pin, mode, &interrupt_handler_20);
	}
	return -1;
}

static void init(void *p) {
	callback_func = p;
}
*/
import "C"
import "unsafe"

import (
	"errors"
	"fmt"
	"github.com/rogpeppe/rog-go/tree/master/exp/callback"
	"sync"
)

const (
	VERSION = "0.2"
	AUTHOR  = "@hugozhu"
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

var mutex = &sync.Mutex{}

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

func WiringPiSetup() error {
	if -1 == int(C.wiringPiSetup()) {
		return errors.New("wiringPiSetup failed to call")
	}
	return nil
}

func PinMode(pin int, mode int) {
	C.my_pinMode(C.int(pin), C.int(mode))
}

func DigitalWrite(pin int, mode int, pwmval int) {
	C.my_digitalWrite(C.int(pin), C.int(mode), C.int(pwmval))
}

func DigitalRead(pin int) int {
	return int(C.my_digitalRead(C.int(pin)))
}

func Delay(ms int) {
	C.delay(C.uint(ms))
}

func DelayMicroseconds(microSec int) {
	C.delayMicroseconds(C.uint(microSec))
}

func WiringPiISR(pin int, mode int) chan int {
	mutex.Lock()
	defer mutex.Unlock()
	if interrupt_chans[pin] == nil {
		interrupt_chans[pin] = make(chan int)
	}
	C.my_wiringPiISR(C.int(pin), C.int(mode))
	return interrupt_chans[pin]
}

func init() {
	C.init(callback.Func)
}

var interrupt_chans = [64]chan int{}

//export goCallback
func goCallback(arg unsafe.Pointer) {
	ctxt := (*C.context)(arg)
	interrupt_chans[int(ctxt.pin)] <- int(ctxt.ret)
}
