package main

import (
	"fmt"
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wcount := make(map[string]int)
	words := strings.Fields(s)
	fmt.Println(words)
	for _, word := range words {
		wcount[word] += 1
 	}
	return wcount
}

func main() {
	wc.Test(WordCount)
}
