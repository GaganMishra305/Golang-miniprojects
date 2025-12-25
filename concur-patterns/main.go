package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// generator
	fmt.Println("Generator begins")
	ch := boring("boring")
	for i := 0; i < 5; i++ {
		fmt.Printf("%s", <-ch)
	}
	fmt.Println("Generator ends")


	// fanin
	fmt.Println("fanin begins")
	c := fanIn(boring("Ann"), boring("Jones"))
	for i := 0; i < 5; i++ {
		fmt.Printf("%s", <-c)
	}
	fmt.Println("Generator begins")

}

// generator --- function that returns a channel
func boring(word string) <-chan string {
	ch := make(chan string)
	go func(){
		for i:= 0; ;i++ {
			ch <- fmt.Sprintln(word, i)
			time.Sleep(time.Duration(rand.Intn(1e3))*time.Millisecond)
		}
	}()
	return ch
}


// fan-in --- function that takes input from two streams and multiplezes them together
func fanIn(input1, input2 <-chan string) chan string {
	c := make(chan string)

	go func() { for { c <- <-input1 } }()
	go func() { for { c <- <-input2 } }()
	return c
}


// select statement --- use for timeout and early quitting


// (slow swquential failure-sensitive) ---> (fast concurrent robust replicated)