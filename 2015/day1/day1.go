package day1

import (
	"fmt"
	"io"
	"log"
	"os"
)

var currentStore = 0
var firstflag = false

const inputPath = "./2015/day1/input.txt"

func Run() {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		if err != io.EOF {
			log.Fatal(err)
		}
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
