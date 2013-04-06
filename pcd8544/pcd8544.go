package pcd8544

/*
#cgo LDFLAGS: -L . -lPCD8544 -lwiringPi

#include <wiringPi.h>
#include <PCD8544.h>
#include <stdio.h>
#include <stdlib.h>


void Test(char *c, char *c2) {
	LCDclear();
	LCDdrawstring(0, 24, c);
	LCDdrawstring(0, 34, c2);
	LCDdisplay();
}
*/
import "C"

const (
	BLACK = 1
	WHITE = 0
)

func LCDclear() {
	C.LCDclear()
}

func Test(s1 string, s2 string) {
	C.Test(C.CString(s1), C.CString(s2))
}

func LCDInit(_sclk int, _din int, _dc int, _cs int, _rst int, contrast int) {
	C.LCDInit(C.uint8_t(_sclk), C.uint8_t(_din), C.uint8_t(_dc), C.uint8_t(_cs), C.uint8_t(_rst), C.uint8_t(contrast))
}

func LCDcommand(c int) {
	C.LCDcommand(C.uint8_t(c))
}

func LCDdata(c int8) {
	C.LCDdata(C.uint8_t(c))
}

func LCDdrawstring(x int8, y int8, str string) {
	C.LCDdrawstring(C.uint8_t(x), C.uint8_t(y), C.CString(str))
}

func LCDdisplay() {
	C.LCDdisplay()
}

func LCDdrawline(x0 int, y0 int, x1 int, y1 int, color int) {
	C.LCDdrawline(C.uint8_t(x0), C.uint8_t(y0), C.uint8_t(x1), C.uint8_t(y1), C.uint8_t(color))
}
