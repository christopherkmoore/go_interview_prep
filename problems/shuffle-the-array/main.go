package main

import (
	"example/hello/pkg/timer"
	"fmt"
	"reflect"
)

func main() {

	timer := timer.New()

	n := 3
	input := []int{2, 5, 1, 3, 4, 7}

	output := shuffle(input, n)
	answer := []int{2, 3, 5, 4, 1, 7}

	if reflect.DeepEqual(output, answer) {
		fmt.Println("passing")
	}

	fmt.Print(output)

	timer.LogTimeMS()
}

// array is always 2n, therefore always divisible by 2.
// Slice it in half to get x and y
// CKM: my answer -- passing ~ 49MS.
func shuffle(nums []int, n int) []int {

	x, y := nums[:n], nums[n:]

	var result []int
	for i := 0; i < n; i++ {
		result = append(result, x[i], y[i])
	}

	return result
}

/** Most memory efficient answer: ~ 42MS

func shuffle(nums []int, n int) []int {
    newNums := make([]int, n*2)
    for i, j := 0, 0; i+n < len(nums); i++ {
        newNums[j] = nums[i]
        j++
        newNums[j] = nums[i+n]
        j++
    }
    return newNums
}
*/
