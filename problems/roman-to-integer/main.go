package main

import (
	"example/hello/pkg/timer"
	"fmt"
	"reflect"
)

func main() {

	timer := timer.New()

	input := "III"

	output := romanToInt(input)
	answer := 3

	if reflect.DeepEqual(output, answer) {
		fmt.Println("passing")
	}

	fmt.Printf("output: %v", output)

	timer.LogTimeMS()
}

// CKM: Passing with good numbers. O(n) constant space && time complexity.
func romanToInt(s string) int {

	lookup := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	answer := 0
	for index := 0; index <= len(s)-1; index++ {

		runner := index + 1

		if runner > len(s)-1 {
			answer += lookup[rune(s[index])]
			return answer
		}

		k1, k2 := rune(s[index]), rune(s[runner])
		v1, v2 := lookup[rune(s[index])], lookup[rune(s[runner])]

		if (k2 == 'V' && k1 == 'I') || (k2 == 'D' && k1 == 'C') || (k2 == 'X' && k1 == 'I') || (k2 == 'L' && k1 == 'X') || (k2 == 'M' && k1 == 'C') || (k2 == 'C' && k1 == 'X') {
			answer += v2 - v1
			index++
		} else {
			answer += v1
		}
	}

	return answer
}
