package main

import (
	"example/hello/pkg/timer"
	"fmt"
	"reflect"
)

func main() {

	timer := timer.New()

	input := []int{3, 3, 3, 1, 2, 1, 1}
	// input := []int{0, 1, 2, 2}
	// input := []int{0, 1, 1}
	// input := make([]int, 10*10*10*10*10)

	output := totalFruit(input)
	// answer := 10 * 10 * 10 * 10 * 10
	answer := 4

	if reflect.DeepEqual(output, answer) {
		fmt.Println("passing")
	}

	fmt.Print(output)

	timer.LogTimeMS()
}

// pre notes:
// iterate fruits
// start a counter and save the integer
// if last != next counter++ trees++
// if counter > 2 trees = 0 counter =0 restart
// look for highest trees

// post notes:
// need a runner

// might need a runner
// CKM passing - when n == 10^5 306MS
func totalFruit(fruits []int) int {
	fruitsLength := len(fruits)
	if fruitsLength <= 2 {
		return fruitsLength
	}

	if fruitsLength == 3 {

		if fruits[0] == fruits[1] || fruits[1] == fruits[2] || fruits[0] == fruits[2] {
			return 3
		}

		return 2
	}

	tree1 := -1
	tree2 := -1

	streak := 2
	longestStreak := 0
	runner := 0
	breakout := true

	for i := 0; i < fruitsLength-1; i++ {
		runner = i + 1

		if runner >= fruitsLength {
			return longestStreak
		}

		if fruits[i] != fruits[runner] {
			tree1, tree2 = fruits[i], fruits[runner]
		} else {
			tree1 = fruits[i]
		}

		runner++

		if runner >= fruitsLength {
			return longestStreak
		}

		for breakout {
			if runner >= fruitsLength {
				breakout = false
				continue
			}

			tryTree := fruits[runner]

			if tryTree != tree1 && tree2 == -1 {
				tree2 = tryTree
			}

			if tree2 != -1 && (tryTree != tree1 && tryTree != tree2) {
				breakout = false
			}

			if tryTree == tree1 || tryTree == tree2 {
				streak++
			}

			runner++

		}

		if streak >= longestStreak {
			longestStreak = streak
		}

		streak, tree1, tree2, breakout = 2, -1, -1, true

		if (i + longestStreak) >= fruitsLength {
			break
		}

	}

	return longestStreak
}
