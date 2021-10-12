package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

var m map[int]string

func WordCount(s string) map[string]int {
	wordMap := make(map[string]int)
	words := strings.Fields(s)

	for i := 0; i < len(words); i++ {
		wordMap[words[i]] += 1
	}
	fmt.Println(words)
	fmt.Println(wordMap)

	return wordMap
}

func main() {
	m = make(map[int]string)
	m[1] = "Manchester United"
	m[2] = "Liverpool"
	m[3] = "Arsenal"

	fmt.Println(m)

	elem, ok := m[3]

	if ok {
		fmt.Println(elem)
	} else {
		fmt.Println(ok)
	}

	m[3] = "Aston vila"
	m[4] = "Aston vila"

	fmt.Println(m)

	wc.Test(WordCount)

}
