package day3

import (
	"bytes"
	"fmt"
	"io"
	"log"

	"github.com/stovenn/adventOfCode/util"
)

const inputPath = "./2015/day3/input.txt"

type Point struct {
	x, y int
}

var visitedHouses = 1
var turnFlag bool = true
var santaPos = Point{0, 0}
var robotSantaPos = Point{0, 0}
var visitedHousesPos = []Point{{0, 0}}

func Run() {
	b, err := util.GetInput(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	buffer := bytes.NewBuffer(b)

	for {
		r, _, err := buffer.ReadRune()
		if err == io.EOF {
			break
		}

		if turnFlag {
			move(r, &santaPos)
		} else {
			move(r, &robotSantaPos)
		}
	}

	fmt.Println("Visited houses:", visitedHouses)
}

func move(r rune, position *Point) {
	doMove(r, position)
	if !alreadyVisited(*position) {
		visitedHouses++
		visitedHousesPos = append(visitedHousesPos, *position)
	}
	turnFlag = !turnFlag
}

func doMove(movement rune, pos *Point) {
	switch movement {
	case 'v':
		pos.y--
	case '^':
		pos.y++
	case '<':
		pos.x--
	case '>':
		pos.x++
	}
}

func alreadyVisited(currentPos Point) bool {
	for _, pos := range visitedHousesPos {
		if currentPos == pos {
			return true
		}
	}
	return false
}
