package main

import "fmt"

func main() {
	d := 42

	fmt.Println(d)
	//fmt.Printf("\n%d\n", g)
	//Print the number in binary
	fmt.Printf("\n%b\n", d)
	//Print the number in hexadecimal
	fmt.Printf("\n%x\n", d)

	a := 42
	b := 24

	g := a == b //a equal to b
	h := a <= b
	i := a >= b // a greater or equal to
	j := a != b // a not equal to b

	k := a > b // a greater than b
	l := a < b // a less than b

	fmt.Printf("g: %v, h: %v, i: %v, j: %v, k: %v, l: %v\n", g, h, i, j, k, l)

}
