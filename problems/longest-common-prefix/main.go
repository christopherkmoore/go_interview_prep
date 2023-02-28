package main

import (
	"example/hello/pkg/timer"
	"fmt"
	"reflect"
)

// input.length <= 200 && input[i].length <= 200. input[i] is always lowercase letters.
// find the longest common prefix and return it as a string. If there is no prefix return an empty string.
func main() {

	timer := timer.New()

	input := []string{"flower", "flow", "flight"}
	// input := []string{"dog", "racecar", "car"}
	// input := []string{"ab", "a"}
	// input := []string{"reflower", "flow", "flight"}
	// input := []string{"flower", "flower", ""}
	// input := []string{"cir", "car"}

	output := longestCommonPrefix(input)
	answer := "fl"
	// answer := "fl"
	// answer := "a"
	// answer := "c"
	// answer := ""

	if reflect.DeepEqual(output, answer) {
		fmt.Println("passing")
	}

	fmt.Printf("output: %v", output)

	timer.LogTimeMS()
}

// CKM: Taking too long, go next.
func longestCommonPrefix(strs []string) string {

	return "Hello World"
}

func sort(strs []string) string {

	shortest := strs[0]
	for _, word := range strs {
		if len(word) < len(shortest) {
			shortest = word
		}
	}

	return shortest
}
