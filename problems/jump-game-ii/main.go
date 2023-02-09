package main

import (
	"example/hello/pkg/timer"
	"fmt"
	"reflect"
)

func main() {

	timer := timer.New()

	input := []int{2, 3, 1, 1, 4}
	// input := []int{2, 2, 1, 1, 4}
	// input := []int{4, 2, 1, 1, 4}
	// input := []int{1, 2, 3}
	// input := []int{1, 1, 1, 1}
	// input := []int{1, 1, 2, 1, 1}
	// input := []int{1, 2, 1, 1, 1}

	output := jump(input)
	answer := 2

	if reflect.DeepEqual(output, answer) {
		fmt.Println("passing")
	}

	// fmt.Println(output)

	timer.LogTimeMS()
}

// check nums[i] > range nums[i] in steps
// CKM: not passing -- this one did not go well
// func jump(nums []int) int {

// 	// check for output 2:
// 	// first jump to nums[0]

// 	// n-1 = i + j1 + j2...

// 	lengthNums := len(nums)
// 	goal := lengthNums - 1

// 	if lengthNums == 1 {

// 		if nums[0] == 0 || nums[0] == 1 {
// 			return 0
// 		}

// 		return 1
// 	}

// 	if lengthNums > 0 {
// 		if nums[0] >= goal {
// 			return 1
// 		}
// 	}

// 	if lengthNums == 2 {
// 		if nums[0] == 1 {
// 			return 1
// 		}

// 		return 0
// 	}

// 	breakout := true
// 	steps := 1
// 	shortestDistance := goal

// 	for i := 0; i < lengthNums-1; i++ {
// 		next := i + 1
// 		steps = next + 1
// 		goal = lengthNums - (next)

// 		if next >= lengthNums-1 {
// 			break
// 		}

// 		// [2, 3, 1, 1, 4]
// 		// [1, 1, 2, 1, 1]
// 		fmt.Printf("for i -- next, steps, goal: %v, %v, %v, %v\n", i, next, steps, goal)

// 		if nums[next] >= goal {
// 			shortestDistance = steps
// 			return shortestDistance
// 		}

// 		if nums[next] > goal {
// 			continue
// 		}

// 		if nums[next] < (goal) {
// 			runner := next
// 			trySteps := steps
// 			for breakout {
// 				if nums[runner] < (goal) {
// 					runner++
// 					trySteps++

// 					if runner >= lengthNums-1 {
// 						breakout = false
// 						trySteps--
// 					}

// 					if nums[runner] >= (goal) {
// 						breakout = false
// 						trySteps--
// 					}
// 				}

// 			}

// 			if shortestDistance > steps {
// 				shortestDistance = steps + trySteps
// 			}
// 		}

// 	}

// 	return shortestDistance
// }

// Cheater answer with modifications made by me (ripped from js best answer.)
// take aways:
// - you can put the break condition straight in the loop
func jump(nums []int) int {
	goal, current, next, answer := len(nums)-1, -1, 0, 0

	for i := 0; next < goal; i++ {

		if i > current {
			answer++
			current = next
		}

		fmt.Printf("iteration %v\n", i)
		fmt.Printf("goal: %v\n", goal)
		fmt.Printf("current: %v\n", current)
		fmt.Printf("next: %v\n", next)
		fmt.Printf("answer: %v\n", answer)

		next = max(next, nums[i]+i)

		fmt.Printf("next: %v\n\n", next)

	}

	return answer
}

// 1, 2, 3
func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
