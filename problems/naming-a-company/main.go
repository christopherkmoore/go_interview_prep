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
	input := []string{"bzklqtbdr", "kaqvdlp", "r", "dk"}
	output := distinctNames(input)
	answer := 12

	if reflect.DeepEqual(output, answer) {
		fmt.Println("passing")
	}

	fmt.Println(output)

	timer.LogTimeS()
}

// Notes
// CKM: this is passing i think but taking too long.
// The website doesn't want to spend the $ on lambda processing time I guess...

// for redo
// create a map like this
// c: {"offee"},
// d: {"onuts"},
// t: {"ime", "offee"}
// and count those entries that way you don't need to iterate a bunch

func distinctNames(ideas []string) int64 {

	if len(ideas) == 2 {
		if ideas[0][1:] == ideas[1][1:] {
			return 0
		}

		return 2
	}

	unfiltered := make(map[string]int)
	counter := make(map[string]int)
	answer := 0
	for _, word := range ideas {
		counter[word] = -1
	}

	for index, word := range ideas {
		runner := index + 1
		trailer := index - 1

		for runner < len(ideas) {
			cut1, cut2 := swapLetters(word, ideas[runner])

			if !containsKey(counter, cut1) && !containsKey(counter, cut2) {
				new := strings.Join([]string{cut1, cut2}, " ")
				new2 := strings.Join([]string{cut2, cut1}, " ")

				if !containsKey(unfiltered, new) {
					unfiltered[new] += 1
					answer++
				}

				if !containsKey(unfiltered, new2) {
					unfiltered[new2] += 1
					answer++
				}
			}
			runner++
		}

		for trailer >= 0 {
			cut1, cut2 := swapLetters(word, ideas[trailer])

			if !containsKey(counter, cut1) && !containsKey(counter, cut2) {
				new := strings.Join([]string{cut1, cut2}, " ")
				new2 := strings.Join([]string{cut2, cut1}, " ")

				if !containsKey(unfiltered, new) {
					unfiltered[new] += 1
					answer++
				}

				if !containsKey(unfiltered, new2) {
					unfiltered[new2] += 1
					answer++
				}
			}
			trailer--

		}
	}

	return int64(answer)
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
