package main

import (
	"example/hello/pkg/timer"
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	timer := timer.New()

	input := 1211

	output := isPalindrome(input)
	answer := false

	if reflect.DeepEqual(output, answer) {
		fmt.Println("passing")
	}

	fmt.Print(output)

	timer.LogTimeMS()
}

// simple solution, conver the number to a string and evaluate the characters

// without strings: ?
// []int split, reverse, join == x?
func isPalindrome(x int) bool {
	// input < 10 == true && input > 0 false
	if x < 0 {
		return false
	}

	if x > 0 && x < 10 {
		return true
	}

	converted := strconv.Itoa(x)
	split := []rune(converted)

	var result string
	for i, _ := range split {
		backwards := len(converted) - 1 - i

		result += string(split[backwards])
	}

	final, ok := strconv.Atoi(result)

	if ok == nil && final == x {
		return true
	}

	return false
}

// without strings: x % 10
