package rpi

/*
#cgo LDFLAGS: -lwiringPi

#include <wiringPi.h>
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
    pinMode(p,m);
}

static void my_pullUpDnControl(int p, int m) {
    pullUpDnControl(p,m);
}

static int my_digitalRead(int p) {
    return digitalRead(p);
}

static void my_digitalWrite(int p, int m) {
    digitalWrite(p,m);
}

static void my_pwmWrite(int p, int m) {
    pwmWrite(p,m);
}

static int my_analogRead(int p) {
    return analogRead(p);
}

static void my_analogWrite(int p, int m) {
    analogWrite(p,m);
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
	"sync"

	// "github.com/rogpeppe/rog-go/tree/master/exp/callback"
	"github.com/rogpeppe/rog-go/exp/callback"
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
	// PIN_GPIO_0 0
	PIN_GPIO_0 = 0
	// PIN_GPIO_1 1
	PIN_GPIO_1  = 1
	// PIN_GPIO_2 2
	PIN_GPIO_2  = 2
	// PIN_GPIO_3 3
	PIN_GPIO_3  = 3
	// PIN_GPIO_4 4
	PIN_GPIO_4  = 4
	// PIN_GPIO_5 5
	PIN_GPIO_5  = 5
	// PIN_GPIO_6 6
	PIN_GPIO_6  = 6
	// PIN_GPIO_7 7
	PIN_GPIO_7  = 7
	// PIN_SDA 8
	PIN_SDA     = 8
	// PIN_SCL 9
	PIN_SCL     = 9
	// PIN_CE0 10
	PIN_CE0     = 10
	// PIN_CE1 11
	PIN_CE1     = 11
	// PIN_MOSI 12
	PIN_MOSI    = 12
	// PIN_MOSO 13
	PIN_MOSO    = 13
	// PIN_SCLK 14
	PIN_SCLK    = 14
	// PIN_TXD 15
	PIN_TXD     = 15
	// PIN_RXD 16
	PIN_RXD     = 16
	// PIN_GPIO_8 17
	PIN_GPIO_8  = 17
	// PIN_GPIO_9 18
	PIN_GPIO_9  = 18
	// PIN_GPIO_10 19
	PIN_GPIO_10 = 19
	// PIN_GPIO_11 20
	PIN_GPIO_11 = 20

	// wiringPi modes

	// WPI_MODE_PINS 0
	WPI_MODE_PINS          = C.WPI_MODE_PINS
	// WPI_MODE_GPIO 1
	WPI_MODE_GPIO          = C.WPI_MODE_GPIO
	// WPI_MODE_GPIO_SYS 2
	WPI_MODE_GPIO_SYS      = C.WPI_MODE_GPIO_SYS
	// WPI_MODE_PHYS 3
	WPI_MODE_PHYS          = C.WPI_MODE_GPIO_SYS
	// WPI_MODE_PIFACE 4
	WPI_MODE_PIFACE        = C.WPI_MODE_PIFACE
	// WPI_MODE_UNINITIALISED -1
	WPI_MODE_UNINITIALISED = C.WPI_MODE_UNINITIALISED

	// Pin modes

	// OUTPUT 1
	OUTPUT           = C.OUTPUT
	// INPUT 0
	INPUT            = C.INPUT
	// PWM_OUTPUT 2
	PWM_OUTPUT       = C.PWM_OUTPUT
	// GPIO_CLOCK 3
	GPIO_CLOCK       = C.GPIO_CLOCK
	// SOFT_PWM_OUTPUT 4
	SOFT_PWM_OUTPUT  = 4
	// SOFT_TONE_OUTPUT 1
	SOFT_TONE_OUTPUT = 5
	// PWM_TONE_OUTPUT 6
	PWM_TONE_OUTPUT  = 6

	// LOW 0
	LOW  = C.LOW
	// HIGH 1
	HIGH = C.HIGH

	// Pull up/down/none

	// PUD_OFF 0
	PUD_OFF  = C.PUD_OFF
	// PUD_DOWN 1
	PUD_DOWN = C.PUD_DOWN
	// PUD_UP 2
	PUD_UP   = C.PUD_UP

	// PWM

	// PWM_MODE_MS 0
	PWM_MODE_MS  = C.PWM_MODE_MS
	// PWM_MODE_BAL 1
	PWM_MODE_BAL = C.PWM_MODE_BAL

	// Interrupt levels

	// INT_EDGE_SETUP 0
	INT_EDGE_SETUP   = C.INT_EDGE_SETUP
	// INT_EDGE_FALLING 1
	INT_EDGE_FALLING = C.INT_EDGE_FALLING
	// INT_EDGE_RISING 2
	INT_EDGE_RISING  = C.INT_EDGE_RISING
	// INT_EDGE_BOTH 3
	INT_EDGE_BOTH    = C.INT_EDGE_BOTH

	// Pi model types and version numbers
	// Intended for the GPIO program - USE AT YOUR OWN RISK.

	// PI_MODEL_A 0
	PI_MODEL_A = 0
	// PI_MODEL_B 1
	PI_MODEL_B = 1
	// PI_MODEL_AP 2
	PI_MODEL_AP = 2
	// PI_MODEL_BP 3
	PI_MODEL_BP = 3
	// PI_MODEL_2 4
	PI_MODEL_2 = 4
	// PI_ALPHA 5
	PI_ALPHA = 5
	// PI_MODEL_CM 6
	PI_MODEL_CM = 6
	// PI_MODEL_07 7
	PI_MODEL_07 = 7
	// PI_MODEL_3 8
	PI_MODEL_3 = 8
	// PI_MODEL_ZERO 9
	PI_MODEL_ZERO = 9
	// PI_MODEL_CM3 10
	PI_MODEL_CM3 = 10
	// PI_MODEL_ZERO_W 12
	PI_MODEL_ZERO_W = 12

	// PI_VERSION_1 0
	PI_VERSION_1 = 0
	// PI_VERSION_1_1 1
	PI_VERSION_1_1 = 1
	// PI_VERSION_1_2 2
	PI_VERSION_1_2 = 2
	// PI_VERSION_2 3
	PI_VERSION_2 = 3

	// PI_MAKER_SONY 0
	PI_MAKER_SONY = 0
	// PI_MAKER_EGOMAN 1
	PI_MAKER_EGOMAN = 1
	// PI_MAKER_EMBEST 2
	PI_MAKER_EMBEST = 2
	// PI_MAKER_UNKNOWN 3
	PI_MAKER_UNKNOWN = 3
)

var mutex = &sync.Mutex{}

// BoardToPin specifies to use RPi.GPIO's BOARD numbering
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

// WiringPiSetup initialises wiringPi and assumes that the calling program is going to be using the wiringPi pin numbering scheme. This is a simplified numbering scheme which provides a mapping from virtual pin numbers 0 through 16 to the real underlying Broadcom GPIO pin numbers. See the pins page for a table which maps the wiringPi pin number to the Broadcom GPIO pin number to the physical location on the edge connector.
// This function needs to be called with root privileges.
// One of the setup functions must be called at the start of your program or your program will fail to work correctly. You may experience symptoms from it simply not working to segfaults and timing issues.
func WiringPiSetup() error {
	if -1 == int(C.wiringPiSetup()) {
		return errors.New("wiringPiSetup failed to call")
	}
	return nil
}

// WiringPiSetupGpio is identical to WiringPiSetup(), however it allows the calling programs to use the Broadcom GPIO pin numbers directly with no re-mapping.
// As above, this function needs to be called with root privileges, and note that some pins are different from revision 1 to revision 2 boards.
// One of the setup functions must be called at the start of your program or your program will fail to work correctly. You may experience symptoms from it simply not working to segfaults and timing issues.
func WiringPiSetupGpio() error {
	if -1 == int(C.wiringPiSetupGpio()) {
		return errors.New("wiringPiSetupGpio failed to call")
	}
	return nil
}

// WiringPiSetupPhys is identical to WiringPiSetup(), however it allows the calling programs to use the physical pin numbers on the P1 connector only.
// As above, this function needs to be called with root priviliges.
// One of the setup functions must be called at the start of your program or your program will fail to work correctly. You may experience symptoms from it simply not working to segfaults and timing issues.
func WiringPiSetupPhys() error {
	if -1 == int(C.wiringPiSetupPhys()) {
		return errors.New("wiringPiSetupPhys failed to call")
	}
	return nil
}

// WiringPiSetupSys initialises wiringPi but uses the /sys/class/gpio interface rather than accessing the hardware directly. This can be called as a non-root user provided the GPIO pins have been exported before-hand using the gpio program. Pin numbering in this mode is the native Broadcom GPIO numbers - the same as wiringPiSetupGpio() above, so be aware of the differences between Rev 1 and Rev 2 boards.
// Note: In this mode you can only use the pins which have been exported via the /sys/class/gpio interface before you run your program. You can do this in a separate shell-script, or by using the system() function from inside your program to call the gpio program.
// Also note that some functions have no effect when using this mode as they’re not currently possible to action unless called with root privileges. (although you can use system() to call gpio to set/change modes if needed)
// One of the setup functions must be called at the start of your program or your program will fail to work correctly. You may experience symptoms from it simply not working to segfaults and timing issues.
func WiringPiSetupSys() error {
	if -1 == int(C.wiringPiSetupSys()) {
		return errors.New("wiringPiSetupSys failed to call")
	}
	return nil
}

// PinMode sets the mode of a pin to either INPUT, OUTPUT, PWM_OUTPUT or GPIO_CLOCK. Note that only wiringPi pin 1 (BCM_GPIO 18) supports PWM output and only wiringPi pin 7 (BCM_GPIO 4) supports CLOCK output modes.
// This function has no effect when in Sys mode. If you need to change the pin mode, then you can do it with the gpio program in a script before you start your program.
func PinMode(pin int, mode int) {
	C.my_pinMode(C.int(pin), C.int(mode))
}

// PullUpDnControl sets the pull-up or pull-down resistor mode on the given pin, which should be set as an input. Unlike the Arduino, the BCM2835 has both pull-up an down internal resistors. The parameter pud should be; PUD_OFF, (no pull up/down), PUD_DOWN (pull to ground) or PUD_UP (pull to 3.3v) The internal pull up/down resistors have a value of approximately 50KΩ on the Raspberry Pi.
// This function has no effect on the Raspberry Pi’s GPIO pins when in Sys mode. If you need to activate a pull-up/pull-down, then you can do it with the gpio program in a script before you start your program.
func PullUpDnControl(pin int, mode int) {
	C.my_pullUpDnControl(C.int(pin), C.int(mode))
}

// DigitalRead returns the value read at the given pin. It will be HIGH or LOW (1 or 0) depending on the logic level at the pin.
func DigitalRead(pin int) int {
	return int(C.my_digitalRead(C.int(pin)))
}

// DigitalWrite writes the value HIGH or LOW (1 or 0) to the given pin which must have been previously set as an output.
// WiringPi treats any non-zero number as HIGH, however 0 is the only representation of LOW.
func DigitalWrite(pin int, mode int) {
	C.my_digitalWrite(C.int(pin), C.int(mode))
}

// AnalogRead returns the value read on the supplied analog input pin. You will need to register additional analog modules to enable this function for devices such as the Gertboard, quick2Wire analog board, etc.
func AnalogRead(pin int) int {
	return int(C.my_analogRead(C.int(pin)))
}

// AnalogWrite writes the given value to the supplied analog pin. You will need to register additional analog modules to enable this function for devices such as the Gertboard.
func AnalogWrite(pin int, mode int) {
	C.my_analogWrite(C.int(pin), C.int(mode))
}

// PwmWrite writes the value to the PWM register for the given pin. The Raspberry Pi has one on-board PWM pin, pin 1 (BMC_GPIO 18, Phys 12) and the range is 0-1024. Other PWM devices may have other PWM ranges.
// This function is not able to control the Pi’s on-board PWM when in Sys mode.
func PwmWrite(pin int, mode int) {
	C.my_pwmWrite(C.int(pin), C.int(mode))
}

// Delay causes program execution to pause for at least howLong milliseconds. Due to the multi-tasking nature of Linux it could be longer. Note that the maximum delay is an unsigned 32-bit integer or approximately 49 days.
func Delay(ms int) {
	C.delay(C.uint(ms))
}

// DelayMicroseconds causes program execution to pause for at least howLong microseconds. Due to the multi-tasking nature of Linux it could be longer. Note that the maximum delay is an unsigned 32-bit integer microseconds or approximately 71 minutes.
func DelayMicroseconds(microSec int) {
	C.delayMicroseconds(C.uint(microSec))
}

// WiringPiISR registers a function to received interrupts on the specified pin. The edgeType parameter is either INT_EDGE_FALLING, INT_EDGE_RISING, INT_EDGE_BOTH or INT_EDGE_SETUP. If it is INT_EDGE_SETUP then no initialisation of the pin will happen - it’s assumed that you have already setup the pin elsewhere (e.g. with the gpio program), but if you specify one of the other types, then the pin will be exported and initialised as specified. This is accomplished via a suitable call to the gpio utility program, so it need to be available.
// The pin number is supplied in the current mode - native wiringPi, BCM_GPIO, physical or Sys modes.
// This function will work in any mode, and does not need root privileges to work.
// The function will be called when the interrupt triggers. When it is triggered, it’s cleared in the dispatcher before calling your function, so if a subsequent interrupt fires before you finish your handler, then it won’t be missed. (However it can only track one more interrupt, if more than one interrupt fires while one is being handled then they will be ignored)
// This function is run at a high priority (if the program is run using sudo, or as root) and executes concurrently with the main program. It has full access to all the global variables, open file handles and so on.
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
