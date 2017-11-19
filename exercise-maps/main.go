package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

// WordCount returns a map containing for every word the number of occurances in the text
func WordCount(s string) map[string]int {
	parts := strings.Fields(s)
	m := make(map[string]int)
	for _, p := range parts {
		m[p]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
