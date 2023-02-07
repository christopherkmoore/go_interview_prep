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
// CKM -- passing 94 MS
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

// without strings
// CKM - passing 94MS
// func isPalindrome(x int) bool {

// 	if x < 0 {
// 		return false
// 	}

// 	if x > 0 && x < 10 {
// 		return true
// 	}

// 	base := 1
// 	incrementor := 10
// 	check := base * incrementor

// 	last := 0
// 	iterator := getIterator(x)

// 	var reversed []int
// 	for i := 0; i < iterator; i++ {
// 		last = x % check
// 		reverse := float64(base) / (float64(check) / float64(incrementor))

// 		digit := int(float64(last) * reverse)

// 		reversed = append(reversed, digit)
// 		check = check * incrementor // 10^1, 10^2 ... 10^n
// 	}

// 	for i := 0; i < iterator; i++ {
// 		base = base * 10
// 	}

// 	var forwards []int
// 	for i := 0; i < iterator; i++ {
// 		check := x % base
// 		reverse := float64(1) / (float64(base) / float64(incrementor))

// 		digit := int(float64(check) * reverse)

// 		forwards = append(forwards, digit)
// 		base = base / incrementor // 10^1, 10^2 ... 10^n
// 	}

// 	return reflect.DeepEqual(reversed, forwards)
// }

// func getIterator(x int) int {
// 	base := 1
// 	incrementor := 10
// 	check := base * incrementor

// 	count := 0
// 	for {
// 		count++

// 		if x%check == x {
// 			break
// 		}

// 		check = check * incrementor // 10^1, 10^2 ... 10^n
// 	}

// 	return count
// }
