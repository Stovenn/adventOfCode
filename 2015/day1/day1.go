package day1

import (
	"fmt"
	"log"

	"github.com/stovenn/adventOfCode/util"
)

var currentStore = 0
var firstflag = false

const inputPath = "./2015/day1/input.txt"

func Run() {
	b, err := util.GetInput(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	for i, c := range b {
		switch string(c) {
		case "(":
			currentStore++
		case ")":
			currentStore--
		}
		if currentStore == -1 && !firstflag {
			fmt.Println("First basement position:", i+1)
			firstflag = true
		}
	}

	fmt.Println("Final store:", currentStore)
}
