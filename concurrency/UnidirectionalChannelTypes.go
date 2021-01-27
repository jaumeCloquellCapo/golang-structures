package main

import (
	"fmt"
)

/***
The firs t goro utine, counter, generates the int egers 0, 1, 2, ..., and sends them overachannel to
the secon d goro utine, squarer, which receives each value, squ ares it, and sends the result over
anot her channel to the third goroutine, pr int er, which receives the squ are d values and prints
them. For clarity of this example, we have int ent ion ally chosen ver y simple functions, thoug h
of cours e they are too computation ally trivial to war rant their own goroutines in a realist ic
prog ram.
 */

func counter(naturals chan <- int) {
	for x := 1; x < 100; x++ {
		fmt.Printf("*** Sending the value => %v \n", x)
		naturals <- x
	}
	close(naturals)
}

func squarer(naturals <- chan int, squares chan <- int) {
	for  {
		natural, ok := <- naturals
		if !ok {
			break // channel was closed and drained
		}
		fmt.Printf("--- Receive the value => %v \n", natural)
		squares <- natural*natural
	}
	close(squares)
}


func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(naturals, squares)
	for x := range squares {
		fmt.Println(x)
	}
}