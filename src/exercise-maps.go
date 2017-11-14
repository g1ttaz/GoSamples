package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

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
