package main

import (
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

type position [2]int
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
			headPos[1]--
		case down:
			headPos[1]++
		case left:
			headPos[0]--
		case right:
			headPos[0]++
		}

		game.outcome(headPos)
		game.snake.body = append([]position{headPos}, game.snake.body...)
		// if ate food
		if samePosition(headPos, game.food) {
			game.score++
			game.orderFood()
		} else {
			game.snake.body = game.snake.body[:len(game.snake.body)-1]
		}
		game.draw()
	}

}

func start() *game {
	x, y := termSize()
	newBody := position{x / 2, y / 2}
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
func samePosition(a, b position) bool {
	return a[0] == b[0] && a[1] == b[1]
}
func (game *game) outcome(headPos position) {
	width := headPos[0]
	height := headPos[1]
	maxX, maxY := termSize()

	// hit wall
	if width > maxX || width < 1 || height > maxY || height < 1 {
		game.over()
	}

	// if run into one self
	for _, v := range game.snake.body {
		if width == v[0] && height == v[1] {
			game.over()
		}
	}

}

func (g *game) draw() {
	clear()
	border()
	maxW, _ := termSize()

	score := "Score: " + strconv.Itoa(g.score)
	moveCursor(position{(maxW / 2) - len(score), 1})
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
	}
	render()
	if g.snake.direction == up || g.snake.direction == down {
		time.Sleep(time.Millisecond * 90)
	} else {
		time.Sleep(time.Millisecond * 30)
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
	moveCursor(position{1, 1})
	draw(" Game over.\n Score: " + strconv.Itoa(g.score))
	render()
	// fmt.Println(g.snake.body)
	os.Exit(0)
}

func randomPosition() [2]int {
	x, y := termSize()
	//TODO  shouldn't be (1,1) or (x,y)
	width := rand.Intn(x - 1)
	height := rand.Intn(y - 1)
	return position{width, height}
}

func (g *game) orderFood() {
	for {
		newFood := randomPosition()
		maxX, maxY := termSize()
		for _, pos := range g.snake.body {
			// !spwan food on snake body
			if newFood[0] == pos[0] || newFood[1] == pos[1] {
				continue
			}
		}
		// !spwan food on border
		if newFood[0] == 1 || newFood[1] == 1 || newFood[0] == maxX || newFood[1] == maxY {
			continue
		}

		g.food = newFood
		break
	}
}
