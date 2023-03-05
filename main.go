package main

import (
	"fmt"
	"time"

	"github.com/stovenn/adventOfCode/2015/day1"
)

func main() {
	start := time.Now()
	day1.Run()
	fmt.Println(time.Since(start))
}
