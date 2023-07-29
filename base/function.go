package main

import "fmt"

func funs() (x, y int) {
	fmt.Printf("Functions")
	return x, y // explicit return
}

func Newfuns() (int, int) {
	var x int
	var y int
	return x, y // explicit return
}

func newFunc(a, b int) (x, y int) {
	x = a
	y = b
	return // implicit return
}
