package main

import (
	"fmt"
	"reflect"
)

func main() {

	n := 3
	input := []int{2, 5, 1, 3, 4, 7}

	output := shuffle(input, n)
	answer := []int{2, 3, 5, 4, 1, 7}

	if reflect.DeepEqual(output, answer) {
		fmt.Println("passing")
	}

	fmt.Print(output)

}

func shuffle(nums []int, n int) []int {
	// array is always 2n, therefore always divisible by 2.
	// Slice it in half to get x and y

	x, y := nums[:n], nums[n:]

	// fmt.Printf("x array: %v", x)
	// fmt.Printf("y array: %v", y)

	var result []int
	for i := 0; i < n; i++ {
		result = append(result, x[i])
		result = append(result, y[i])
	}

	return result
}
