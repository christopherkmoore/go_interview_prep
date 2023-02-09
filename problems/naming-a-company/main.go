package main

import (
	"example/hello/pkg/timer"
	"fmt"
	"reflect"
	"strings"
)

func main() {

	timer := timer.New()

	// input := []string{"coffee", "donuts", "time", "toffee"}
	// input := []string{"lack", "back"}
	// input := []string{"aaa", "baa", "caa", "bbb", "cbb", "dbb"}
	// input := []string{"bzklqtbdr", "kaqvdlp", "r", "dk"}
	output := distinctNames(input)
	answer := 12

	if reflect.DeepEqual(output, answer) {
		fmt.Println("passing")
	}

	fmt.Println(output)

	timer.LogTimeS()
}

func distinctNames(ideas []string) int64 {

	if len(ideas) == 2 {
		if ideas[0][1:] == ideas[1][1:] {
			return 0
		}

		return 2
	}

	unfiltered := make(map[string]int)
	counter := make(map[string]int)
	for _, word := range ideas {
		counter[word] = -1
	}

	for index, word := range ideas {
		runner := index + 1
		trailer := index - 1

		for runner < len(ideas) {
			cut1, cut2 := swapLetters(word, ideas[runner])

			if !containsKeyWithValue(counter, cut1, -1) && !containsKeyWithValue(counter, cut2, -1) {
				new := strings.Join([]string{cut1, cut2}, " ")
				new2 := strings.Join([]string{cut2, cut1}, " ")

				if !containsKey(unfiltered, new) {
					unfiltered[new] += 1
				}

				if !containsKey(unfiltered, new2) {
					unfiltered[new2] += 1
				}
			}
			runner++
		}

		for trailer >= 0 {
			cut1, cut2 := swapLetters(word, ideas[trailer])

			if !containsKeyWithValue(counter, cut1, -1) && !containsKeyWithValue(counter, cut2, -1) {
				new := strings.Join([]string{cut1, cut2}, " ")
				new2 := strings.Join([]string{cut2, cut1}, " ")

				if !containsKey(unfiltered, new) {
					unfiltered[new] += 1
				}

				if !containsKey(unfiltered, new2) {
					unfiltered[new2] += 1
				}
			}
			trailer--

		}
	}

	return int64(len(unfiltered))
}

func swapLetters(word1, word2 string) (string, string) {
	return word1[:1] + word2[1:], word2[:1] + word1[1:]
}

func containsKey(lookup map[string]int, key string) bool {
	if _, ok := lookup[key]; ok {
		return true
	}

	return false
}

func containsKeyWithValue(lookup map[string]int, key string, value int) bool {
	if val, ok := lookup[key]; ok && val == value {
		return true
	}

	return false
}
