package main

import "testing"

func Test_updateMessage(t *testing.T) {
	msg = "Hello World"
	wg.Add(2)
	go updateMessage("x")
	go updateMessage("Goodbye world")
	wg.Wait()

	if msg != "Goodbye world" {
		t.Error("incorrect value in msg")
	}
}
