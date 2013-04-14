package main

import (
	. "github.com/hugozhu/rpi"
	"github.com/hugozhu/rpi/pcd8544"
	"log"
	"os/exec"
	"time"
)

const (
	DIN        = PIN_MOSI
	SCLK       = PIN_SCLK
	DC         = PIN_GPIO_2
	RST        = PIN_GPIO_0
	CS         = PIN_CE0
	PUSHBUTTON = PIN_GPIO_6
	CONTRAST   = 40 //may need tweak for each Nokia 5110 screen
)

var screen_chan chan int
var TOTAL_MODES = 3

func init() {
	WiringPiSetup()
	pcd8544.LCDInit(SCLK, DIN, DC, CS, RST, CONTRAST)
	screen_chan = make(chan int, 1)
}

func main() {
	//a goroutine to check button push event
	go func() {
		last_time := time.Now().UnixNano() / 1000000
		btn_pushed := 0
		for pin := range WiringPiISR(PUSHBUTTON, INT_EDGE_FALLING) {
			if pin > -1 {
				n := time.Now().UnixNano() / 1000000
				delta := n - last_time
				if delta > 300 { //software debouncing
					log.Println("btn pushed")
					last_time = n
					btn_pushed++
					screen_chan <- btn_pushed % TOTAL_MODES //switch the screen display
				}
			}
		}
	}()

	//a groutine to update display every 5 seconds
	go loop_update_display()

	//set screen 0 to be default
	screen_chan <- 0

	ticker := time.NewTicker(5 * time.Second)

	for {
		<-ticker.C
		screen_chan <- -1 //refresh current screen every 5 seconds
	}
}

func loop_update_display() {
	current_screen := 0
	for screen := range screen_chan {
		if screen >= 0 {
			if screen != current_screen {
				//btn pushed
				current_screen = screen
				display_loading()
			}
		}
		switch current_screen {
		case 0:
			display_screen0()
		case 1:
			display_screen1()
		case 2:
			display_screen2()
		}
	}
}

func display_loading() {
	pcd8544.LCDclear()
	pcd8544.LCDdrawstring(0, 20, "Loading ...")
	pcd8544.LCDdisplay()
}

func display_screen0() {
	out, err := exec.Command("/root/bin/screen_1.sh").CombinedOutput()
	if err != nil {
		out = []byte(err.Error())
	}

	pcd8544.LCDclear()
	pcd8544.LCDdrawstring(0, 0, string(out))
	pcd8544.LCDdisplay()
}

func display_screen1() {
	out, err := exec.Command("/root/bin/screen_1.sh").CombinedOutput()
	if err != nil {
		out = []byte(err.Error())
	}

	pcd8544.LCDclear()
	pcd8544.LCDdrawstring(0, 0, string(out))
	pcd8544.LCDdisplay()
}

func display_screen2() {
	out, err := exec.Command("/root/bin/screen_2.sh").CombinedOutput()
	if err != nil {
		out = []byte(err.Error())
	}

	pcd8544.LCDclear()
	pcd8544.LCDdrawstring(0, 0, string(out))
	pcd8544.LCDdisplay()
}
