// Not finished
package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var result *ListNode

	l1Str := reverse(l1.valuesAsString(""))
	resultl1Int, _ := strconv.Atoi(l1Str)

	l2Str := reverse(l2.valuesAsString(""))
	resultl2Int, _ := strconv.Atoi(l2Str)

	finalInt := resultl1Int + resultl2Int

	finalStr := strconv.Itoa(finalInt)

	final := reverse(finalStr)

	var last *ListNode
	for _, str := range final {
		val, _ := strconv.Atoi(string(str))

		new := ListNode{
			Val:  val,
			Next: nil,
		}

		if last == nil {
			last = &new
		} else {
			last.Next = &new
		}

	}

	return result1

}

func (node *ListNode) valuesAsString(cache string) string {

	if node != nil {
		cache = cache + strconv.Itoa(node.Val)
		return node.Next.valuesAsString(cache)
	}

	return cache
}

func reverse(str string) (result string) {
	for _, rune := range str {
		result = string(rune) + result
	}
	return
}

func main() {
	// [2,4,3]
	l1 := ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	// [5,6,4]
	l2 := ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	result1 := l1.valuesAsString("")
	result2 := reverse(result1)
	resultInt, _ := strconv.Atoi(result2)
	// result2 := l2.valuesAsString("")

	// result := addTwoNumbers(&l1, &l2)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(resultInt)
}
