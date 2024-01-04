package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/term"
)

var screen = bufio.NewWriter(os.Stdout)

func hideCursor() {
	fmt.Fprint(screen, "\033[?25l")
}

func showCursor() {
	fmt.Fprint(screen, "\033[?25h")
}
func moveCursor(pos [2]int) {
	fmt.Fprintf(screen, "\033[%d;%dH", pos[1], pos[0])
}

func clear() {
	fmt.Fprint(screen, "\033[2J")
}

func draw(str string) {
	fmt.Fprint(screen, str)
}

// write all data in buffer to terminal
func render() {
	screen.Flush()
}

func termSize() (width int, height int) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(err)
	}
	return width, height
}

func border() {
	maxX, maxY := termSize()
	// ---- top ---- bottom
	for i := 1; i <= maxX; i++ {
		fmt.Fprintf(screen, "\033[1;%dH", i)
		fmt.Fprintf(screen, "_")
		fmt.Fprintf(screen, "\033[%d;%dH", maxY, i)
		fmt.Fprintf(screen, "_")
	}
	// | left | right |
	for i := 1; i <= maxY; i++ {
		fmt.Fprintf(screen, "\033[%d;1H", i)
		fmt.Fprintf(screen, "|")
		fmt.Fprintf(screen, "\033[%d;%dH", i, maxX)
		fmt.Fprintf(screen, "|")
	}
}
