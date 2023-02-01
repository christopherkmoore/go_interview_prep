package main

import "fmt"


// simplest solution n^2
func twoSums(nums []int, target int) []int { 
	max := len(nums) -1 
    var solution []int
	
    for index, _ := range nums { 
        for i := 0; i < max; i++ { 
            if i == index { 
                continue
            }

            if nums[index] + nums[i] == target { 
                
				if i > index { 
					solution = append(solution, index, i)
				} else { 
					solution = append(solution, i, index)
				}
                
				return solution
            }
        }
    }

	return solution
}



func main() {
	input := []int {3, 2, 4}
	target := 6
    result := twoSums(input, target)

	fmt.Println(result)

	input = []int{2,7,11,15}
	target = 9

	result = twoSums(input, target)
	fmt.Print("Result 2: ")
	fmt.Println(result)
}