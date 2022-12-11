package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := os.Args[1]
	f, e := os.Open(fileName)

	if e != nil {
		fmt.Println("Error occured", e)
		os.Exit(1)
	}

	bs, er := io.Copy(os.Stdout, f)

	if er != nil {
		fmt.Println("Error occured", er)
		os.Exit(1)
	}

	fmt.Println(string(bs))
}
