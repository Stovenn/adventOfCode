package day5

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/stovenn/adventOfCode/util"
)

const (
	vowels    = "aieuo"
	inputPath = "./2015/day5/input.txt"
)

var (
	counter   = 0
	forbidden = [4]string{"ab", "cd", "pq", "xy"}
)

func Run() {
	b, err := util.GetInput(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	buffer := bytes.NewBuffer(b)

	for {
		word, err := buffer.ReadString('\n')
		if err == io.EOF {
			break
		}
		// Part 1
		// if hasThreeOrMoreVowels(word) &&
		// 	hasTwoConsecutiveIdenticalChars(word) &&
		// 	doesNotContainsForbiddenSubstring(word) {
		// 	counter++
		// }

		// Part 2
		if containsCharPairTwice(word) && isSubStrSurroundedBySameChar(word) {
			counter++
		}
	}

	fmt.Println("Nice words:", counter)
}

func isSubStrSurroundedBySameChar(str string) bool {
	for i := range str {
		if i+2 == len(str) {
			break
		}
		subString := str[i : i+3]
		if subString[0] == subString[2] {
			return true
		}
	}
	return false
}

func doesNotContainsForbiddenSubstring(str string) bool {
	for _, f := range forbidden {
		if strings.Contains(str, f) {
			return false
		}
	}
	return true
}

func hasTwoConsecutiveIdenticalChars(str string) bool {
	for i, c := range str {
		if i+1 == len(str) {
			break
		}
		if str[i+1] == byte(c) {
			return true
		}
	}
	return false
}

func hasThreeOrMoreVowels(str string) bool {
	vowelsCount := 0
	for _, c := range str {
		if strings.ContainsAny(string(c), vowels) {
			vowelsCount++
		}
	}
	return vowelsCount >= 3
}

func containsCharPairTwice(str string) bool {
	m := map[string]struct {
		exists   bool
		endIndex int
	}{}

	for i := range str {
		if i+1 == len(str) {
			break
		}
		subStr := str[i : i+2]
		if m[subStr].exists && m[subStr].endIndex != i {
			return true
		} else if !m[subStr].exists {
			m[subStr] = struct {
				exists   bool
				endIndex int
			}{exists: true, endIndex: i + 1}
		}
	}
	return false
}
