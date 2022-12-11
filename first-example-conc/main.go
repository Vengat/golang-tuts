package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {

	var wg sync.WaitGroup
	// wg.Add(9)
	// go printSomething("First print", &wg)

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	wg.Add(len(words))

	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
	}

	// bad idea to use time sleep
	// time.Sleep(1 * time.Second)

	wg.Wait()
	wg.Add(1)
	printSomething("Second print", &wg)
}
