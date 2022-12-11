package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main_exe3() {
	s := fmt.Sprintf("%d\t%s\t%t\n", x, y, z)
	fmt.Println(s)
}
