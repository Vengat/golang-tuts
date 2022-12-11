package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	// "github.com/golang/go/blob/master/src/net/http"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	io.Copy(os.Stdout, resp.Body)
}
