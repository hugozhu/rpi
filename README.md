WiringPi-Go
============

Golang wrapped version of Gordon's Arduino-like WiringPi for the Raspberry Pi

# Installation

install WiringPi first

```
cd WiringPi\wiringPi
sudo make install

go get github.com/hugozhu/rpi
```

# GPIO numbering

wiringPi   | Name     | GPIO.BOARD    | GPIO.BCM
---------- | -------- | ------------  | --------
0          |GPIO 0    | 11            | 17 
1          |GPIO 1    | 12            | 18
2          |GPIO 2    | 13            | 21
3          |GPIO 3    | 15            | 22
4          |GPIO 4    | 16            | 23
5          |GPIO 5    | 18            | 24
6          |GPIO 6    | 22            | 25
7          |GPIO 7    | 7             | 4
8          |SDA       | 3             | 0
9          |SCL       | 5             | 1
10         |CE0       | 24            | 8
11         |CE1       | 26            | 7
12         |MOSI      | 19            | 10
13         |MOSO      | 21            | 9
14         |SCLK      | 23            | 11
15         |TXD       | 8             | 14
16         |RXD       | 10            | 15

more to read at: [http://hugozhu.myalert.info/2013/03/22/19-raspberry-pi-gpio-port-naming.html](http://hugozhu.myalert.info/2013/03/22/19-raspberry-pi-gpio-port-naming.html)

# Sample codes

## lcd.go
```
package main

import (
    . "github.com/hugozhu/rpi"
)

func main() {
    WiringPiSetup()

    //use default pin naming
    PinMode(PIN_GPIO_4, OUTPUT)
    DigitalWrite(PIN_GPIO_4, LOW)
    Delay(400)
    DigitalWrite(PIN_GPIO_4, HIGH)

    //use raspberry pi board pin numbering, similiar to RPi.GPIO.setmode(GPI.BOARD)
    Delay(400)
    DigitalWrite(BoardToPin(16), LOW)
    Delay(400)
    DigitalWrite(BoardToPin(16), HIGH)

    //use raspberry pi board pin numbering, similiar to RPi.GPIO.setmode(GPI.BCM)
    Delay(400)
    DigitalWrite(GpioToPin(23), LOW)
    Delay(400)
    DigitalWrite(GpioToPin(23), HIGH)
}
```

## Run

```
export GOPATH=`pwd`
go install github.com/hugozhu/rpi 
go run src/lcd.go 
```