package day2

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/stovenn/adventOfCode/util"
)

var ribbon int
var wrappingPaper int

const inputPath = "./2015/day2/input.txt"

func Run() {
	b, err := util.GetInput(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	buffer := bytes.NewBuffer(b)
	for {
		line, err := buffer.ReadString('\n')
		if err == io.EOF {
			break
		}
		processLine(line)
	}

	fmt.Println("Wrapping paper needed:", wrappingPaper)
	fmt.Println("Ribbon needed:", ribbon)
}

func processLine(line string) {
	//drop the new line byte and split the string to get the numbers
	args := strings.Split(line[:len(line)-1], "x")

	l, _ := strconv.Atoi(args[0])
	w, _ := strconv.Atoi(args[1])
	h, _ := strconv.Atoi(args[2])

	smallestPerimeter := findSmallestPerimeter(l, w, h)

	area1 := l * w
	area2 := w * h
	area3 := h * l
	bowArea := l * w * h

	minArea := findMinArea(area1, area2, area3)

	ribbon += smallestPerimeter + bowArea
	wrappingPaper += 2*(area1+area2+area3) + minArea
}

func findMinArea(area1, area2, area3 int) int {
	return int(
		math.Min(math.Min(float64(area1), float64(area2)),
			math.Min(float64(area2), float64(area3))))
}

func findSmallestPerimeter(l, w, h int) int {
	p1 := float64(2 * (l + w))
	p2 := float64(2 * (l + h))
	p3 := float64(2 * (w + h))
	return int(math.Min(math.Min(p1, p2), p3))
}
