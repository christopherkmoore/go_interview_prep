// Greatest common divisor of strings
// For two strings s and t, we say "t divides s" if and only if s = t + ... + t (i.e., t is concatenated with itself one or more times).
// Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.

package main

import (
	"fmt"
)

// CKM Notes: I think the thing I struggled with was common, from looking at the discussion. People were saying the question does not make sense
// which if you measure backwards from the questions input / output, it certainly does not.
// I glanced over the expression `s = t + ... + t` in the description which I did not realize says quite abit of the problem.

// The substring to search for will ALWAYS occur from the first index of the larger string. also the substring MUST be repeating so "ABAB" and "ABABAB" seems to
// have a simple solution of "ABAB", though it is not repeating in the longer string. so ABABAB = AB + AB + AB not ABABABA = ABAB + ABAB.

func main() {
	case1Input1 := "ABAB"
	case1Input2 := "ABABAB"

	case1Output := "AB"
	result := gcdOfStrings(case1Input1, case1Input2)

	if result == case1Output {
		fmt.Println("passed")
	} else {
		fmt.Println("failed")
	}

	fmt.Print(result)
}

// TOP solution
// CKM Notes added
func gcdOfStrings(s1 string, s2 string) string {
	// if s1 = s2 + s2 + s2 then this will always be true on a repeating pattern.
	if s1+s2 != s2+s1 {
		return ""
	}

	x := gcd(len(s1), len(s2))

	return s1[:x]
}

func gcd(a, b int) int {
	// any patern will be divisible by the smaller string into the larger with modulo 0, so long as either isn't 0
	for b != 0 {
		a, b = b, a%b
	}

	// any solution will be the first occuring slice
	return a
}

// MY solution. Not working. Turns out I opted for the brute force solution.
// func gcdOfStrings(str1 string, str2 string) string {

// 	result := ""
// 	iterator := ""
// 	evaluator := ""

// 	// find shorter string
// 	if len(str1) > len(str2) {
// 		iterator = str2
// 		evaluator = str1
// 	} else {
// 		iterator = str1
// 		evaluator = str2
// 	}

// 	// take the least common denominator from the shorter string, and check to see it it occurs in the larger string.
// 	maxLength := len(iterator) / 2

// 	for i := 1; i < maxLength; i++ {
// 		resultPendingApproval := ""

// 		sliceMatch := strings.Index(evaluator, iterator[:i]) + len(iterator[:i])

// 		if sliceMatch != -1 {
// 			continueMatching := 0
// 			for continueMatching != -1 {
// 				next := i + 1

// 				if next > maxLength {
// 					break
// 				}

// 				continueMatching = strings.Index(evaluator, iterator[:next])

// 				sliceMatch := strings.Index(evaluator, resultPendingApproval) + len(resultPendingApproval)
// 				if sliceMatch != -1 && strings.Contains(evaluator[sliceMatch:], resultPendingApproval) {
// 					result = resultPendingApproval

// 				}

// 				resultPendingApproval = iterator[i:continueMatching]
// 			}

// 		}
// 	}

// 	return result
// }
