package main

import (
	"fmt"
)

func main() {
	var hello string = "Hello"
	fmt.Printf("%s world!\n", hello)
	fmt.Print("Hello world!\n")
	msg := fmt.Sprintf("Hi and %s", hello)
	fmt.Println(msg)

	//js => can be not known during compile time and run to get value
	//go => must be known in compile time, before running
	const secondsinmminute = 60
	const minutesinhour = 60
	const totalseconds = secondsinmminute * minutesinhour
	fmt.Print(totalseconds)
}

// java and go have automatic garbage collection but go uses relatives less memory
// rust has automatic thus even less memory than go is used

/*
bool => %t
string => %s
float32 float 64  => %f
complex64 complex128
rune %c  => alias for int32 => represents a Unicode code point

int  => %d
int8(upto 1-255 decimal) max number 8 bits of binary can form
int16 int32 int64

uint
uint8  uint16 uint32 uint64 uintptr

VERBS IN GO
%v => any value in natural form
%f %g %e => floating point
%q => quoted string or rune
%T => TYPE OF
%% => Percent sign

Types
Basic => string, number, boolean
Aggregate => arrays, structs
Refrerence => pointers, slices, functions and channels
Interface => a collection of method of signatures
*/
