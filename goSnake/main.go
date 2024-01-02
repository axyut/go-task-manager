package main

import "fmt"

type snake struct {
	position  int
	direction int
}

func main() {
	clear()
	width, height := termSize()
	pos := [2]int{(width / 2) - 2, 1}
	moveCursor(pos)
	draw("0oo")

	render()
	fmt.Println(width, height)

}
