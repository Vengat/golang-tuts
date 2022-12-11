package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessageWithMutex(s string, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	msg = s
	m.Unlock()
}

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func main() {
	msg = "Hello World!"

	// var mutex sync.Mutex

	wg.Add(2)
	go updateMessage("Hello Universe")
	go updateMessage("Hello cosmos")
	wg.Wait()

	fmt.Println(msg)
}
