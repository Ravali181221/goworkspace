package main

import "fmt"

var a int
var b string = "james Bond"
var c bool
var d bool = true

type ravali int

var h1 ravali

type anjan string

var h2 anjan

func main7() {

	e := 42
	f := "Shaken not stirred"
	g := `Hi Ravali`
	h := `Q says,
	"I am learning"`

	fmt.Println(a)

	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(b, "says, ", f)
	fmt.Println(h)

	fmt.Printf("%v\t", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%b\n", e)
	fmt.Printf("%x\n", a)
	a = 911
	fmt.Printf("%#x\n", a)
	fmt.Printf("%v\n", a)

	h1 = 100

	fmt.Println("\n", h)

	h2 = "hello world"

	fmt.Printf("%T\n", h2)

	k100 := int(h1)

	fmt.Println(k100)
	fmt.Printf("%T", k100)

}
