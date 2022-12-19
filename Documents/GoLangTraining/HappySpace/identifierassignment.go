package main

import (
	"fmt"
	//"io"
	//"os"
)

// package scope
var x2 int
var y2 string
var z2 bool

type rava int

var x3 rava

func main9() {

	x2 = 42
	y2 = "James Bond"
	z2 = true

	fmt.Println(x2)
	fmt.Println(y2)
	fmt.Println(z2)

	var s = fmt.Sprintf("%v %v %v\n", x2, y2, z2)
	fmt.Print(s)

	fmt.Printf("%T\n", s)

	var x3 int64 = 10
	fmt.Println(x3)
	fmt.Printf("%T\n", x3)
	x3 = 42
	fmt.Println(x3)

	//var s = fmt.Sprintf("%s\n %d\n %t\n", x2, y2, z2)
	//io.WriteString(os.Stdout, s)
	//fmt.Printf("%T", s)

}
