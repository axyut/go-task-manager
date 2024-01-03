package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/mattn/go-tty"
)

type snake struct {
	body      []position
	direction direction
}

type game struct {
	score int
	snake *snake
	food  position
}

type position map[string]int
type direction int

// up,down,right,left map to 0,1,2,3
const (
	up direction = iota
	down
	right
	left
)

func main() {
	game := start()
	game.handleInterrupt()

	// move in the key pressed direction
	for {
		headPos := game.snake.body[0]
		// this will update the head pos in 1 step different direction
		switch game.snake.direction {
		case up:
			headPos["height"]--
		case down:
			headPos["height"]++
		case left:
			headPos["width"]--
		case right:
			headPos["width"]++
		}

		outcome(game)

		width := headPos["width"]
		height := headPos["height"]
		// if ate food
		if width == game.food["width"] && height == game.food["height"] {
			game.score++
			// game.orderFood()
			// should it work if i made position an array instead of map?
			game.snake.body = append([]position{headPos}, game.snake.body...)
			// game.score = len(game.snake.body)
		}
		game.draw()
	}

}

func start() *game {
	x, y := termSize()
	newBody := position{"width": x / 2, "height": y / 2}
	snake := &snake{
		body:      []position{newBody},
		direction: up,
	}
	game := &game{
		score: 0,
		snake: snake,
		food:  randomPosition(),
	}
	go game.listenForKey()
	return game
}

func outcome(g *game) {
	// if hit wall
	maxX, maxY := termSize()
	width := g.snake.body[0]["width"]
	height := g.snake.body[0]["height"]
	if width > maxX || width < 1 || height > maxY || height < 1 {
		g.over()
	}

	// if hit oneself
	// for _, pos := range g.snake.body {
	// 	if width == pos["width"] || height == pos["height"] {
	// 		g.over()
	// 	}
	// }

}

func (g *game) draw() {
	clear()
	border()
	maxW, _ := termSize()
	score := "Score: " + strconv.Itoa(g.score)
	moveCursor(position{"width": (maxW / 2) - len(score), "height": 1})
	draw(score)
	moveCursor(g.food)
	draw("#")
	for i, v := range g.snake.body {
		moveCursor(v)
		if i == 0 {
			draw("0")
		} else {
			draw("o")
		}
		// moveCursor(position{"width": 1, "height": 1})
		// draw(string(i))
		// draw(strconv.Itoa(v["height"]) + "H " + strconv.Itoa(v["width"]) + "W")
	}
	render()
	if g.snake.direction == up || g.snake.direction == down {
		time.Sleep(time.Millisecond * 300)
	} else {
		time.Sleep(time.Millisecond * 200)
	}
}

func (g *game) listenForKey() {
	tty, err := tty.Open()
	if err != nil {
		panic(err)
	}
	defer tty.Close()
	// listen for keypress
	for {
		if char, err := tty.ReadRune(); err == nil {
			// str := string(char)
			// A -> up B -> down C -> right D -> left
			switch char {
			case 'A':
				g.snake.direction = up
				// fmt.Println(g.snake.direction)
			case 'B':
				g.snake.direction = down
			case 'C':
				g.snake.direction = right
			case 'D':
				g.snake.direction = left
			case 'q':
				g.over()
			}
		}
	}
}

func (g *game) handleInterrupt() {
	hideCursor()

	// handle CTRL C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			g.over()
		}
	}()
}

func (g *game) over() {
	clear()
	showCursor()
	moveCursor(position{"height": 1, "width": 1})
	draw(" Game over.\n Score: " + strconv.Itoa(g.score))
	render()
	fmt.Println(g.snake.body)
	os.Exit(0)
}

func randomPosition() map[string]int {
	x, y := termSize()
	//TODO  shouldn't be (1,1) or (x,y)
	width := rand.Intn(x - 1)
	height := rand.Intn(y - 1)
	return position{"width": width, "height": height}
}

func (g *game) orderFood() {
	for {
		newFood := randomPosition()
		for _, pos := range g.snake.body {
			if newFood["width"] == pos["width"] || newFood["width"] == pos["height"] {
				continue
			}
		}
		g.food = newFood
	}
}
