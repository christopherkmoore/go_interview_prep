package main

import "fmt"

// Slices & arrays
// https://go.dev/blog/slices

type sliceHeader struct {
	Length        int
	Capacity      int
	ZerothElement *byte
}

func Extend(slice []int, element int) []int {

	n := len(slice)
	if n == cap(slice) {
		newSlice := make([]int, len(slice), 2*len(slice)+1)
		copy(newSlice, slice)
		slice = newSlice
	}

	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

// Append with per item allocation checks.
func SimpleAppend(slice []int, items ...int) []int {
	for _, item := range items {
		slice = Extend(slice, item)
	}

	return slice
}

// Append with per item allocation checks.
func Append(slice []int, elements ...int) []int {
	n := len(slice)
	total := n + len(elements)

	if total > cap(slice) {
		newSize := total*3/2 + 1
		newSlice := make([]int, total, newSize)
		copy(newSlice, slice)
		slice = newSlice
	}

	slice = slice[:total]
	copy(slice[n:], elements)
	return slice
}

// CKM Tries generic, seems to work on [] string types but regular strings become a nightmare.
func GenericAppend[T interface{}](slice []T, elements ...T) []T {
	n := len(slice)
	total := n + len(elements)

	if total > cap(slice) {
		newSize := total*3/2 + 1
		newSlice := make([]T, total, newSize)
		copy(newSlice, slice)
		slice = newSlice
	}

	slice = slice[:total]
	copy(slice[n:], elements)
	return slice
}

func main() {

	slice := []int{0, 1, 2, 3, 4}
	fmt.Println(slice)
	slice = Append(slice, 5, 6, 7, 8)
	fmt.Println(slice)

	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice), cap(slice), slice)
	fmt.Println("address of 0th element:", &slice[0])

	s1 := []int{0, 1, 2, 3, 4}
	s2 := []int{5, 6, 7, 8, 9}
	slice = Append(s1, s2...)
	fmt.Println(slice)

	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice), cap(slice), slice)
	fmt.Println("address of 0th element:", &slice[0])

	s3 := []string{"Hello", "World"}
	s4 := []string{"Goodbye", "World"}

	newSlice := GenericAppend(s3, s4...)

	fmt.Println(newSlice)
}
