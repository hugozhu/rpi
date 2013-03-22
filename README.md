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

# Example

```
package main

import (
    . "github.com/hugozhu/rpi"
)

func main() {
    WiringPiSetup()
    PinMode(4, OUTPUT)
    DigitalWrite(4, LOW)
    Delay(400)
    DigitalWrite(4, HIGH)
}
```