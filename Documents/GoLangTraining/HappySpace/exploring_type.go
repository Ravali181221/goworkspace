package main

import "fmt"

// package scope
var z1 int

var a_str string = `Shaken, "ravali", yes`

func main4() {

	z1 = 21

	fmt.Println(z1)
	z2 := 30

	foobar_type()

	fmt.Println(a_str)

	fmt.Printf("Type of z1 %T\n", z1)
	fmt.Printf("Type of a_str %T\n", a_str)
	var z_bool bool
	fmt.Println(z1, z2, z_by, z_ru, z_bool)
	// false for boolean
	// 0 for int
	// 0.0 for floats
	// "" for strings
	// nil for
	// pointer,function,interfaces,slices,channels,maps
	// Note: use short declaration
	// fmt.Println(z1, z2)
}

func foobar_type() {
	fmt.Println(z1)
}
