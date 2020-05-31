package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	count := make(map[string]int)
	// default value for each element in map is going to be zero
	// since it's the int's zero value :)
	words := strings.Fields(s)
	for _, v := range words {
		count[v] += 1
	}
	return count
}

func main() {
	wc.Test(WordCount)
}
