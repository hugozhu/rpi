WiringPi-Go
============

Golang wrapped version of Gordon's Arduino-like WiringPi for the Raspberry Pi

# Installation

install WiringPi first

```
go get github.com/hugozhu/rpi
```

# GPIO naming

wiringPi   | Name     | GPIO Pin      | GPIO Pin in Broadcom 
---------- | -------- | ------------  | ------------ 
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
15         |TxD       | 8             | 14
16         |RxD       | 10            | 15


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