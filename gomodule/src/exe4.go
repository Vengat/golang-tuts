package main

import "fmt"

type customnumber int

var x1 customnumber
var y1 int

func main() {
	fmt.Println(x1)
	fmt.Printf("%T\n", x1)
	y1 = int(x1)
	fmt.Println(y1)
	//one-twenty-eight(2^7) sixty-fours(2^6) thirty-twos(2^5) sixteens(2^4)  eights(2^3) fours(2^2)    twos(2^1)  ones(2^0)
	//       0                    0                1                   0               1         0             1            0
	//x2 := 00101010
	fmt.Printf("%T\n", 42)
	fmt.Printf("42 in binary is %b\n", 42)
	//Hexadecimal 911
	//                 16^6      16^5     16^4    16^3    16^2    16^1     16^0
	//
	//                                                    3      8        F
	fmt.Printf("%#X\n", 911)
}
