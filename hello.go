package main

import (
	"fmt"
	"strings"
)

func main() {
	// defer exectues in form of a stack after the surrounding function returns
	fmt.Println("1. defer: ")
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	// slices are reference to underlying array
	fmt.Println("\n2. slices: ")
	s := []int{1, 2, 3}
	fmt.Println("arra:", s[0:2])

	// maps
	WordCount("I am learning Go. Go is fun. I love Go.")

	// goroutines
	fmt.Println("\n3. goroutines: ")
	Caller()
}

func WordCount(s string) map[string]int {
	var wc = make(map[string]int)

	for _, word := range strings.Fields(s) {
		wc[word]++
	}

	return wc
}
