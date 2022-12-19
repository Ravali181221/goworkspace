package main

import (
	"fmt"
	"math"
)

func main2() {
	fmt.Println("hello team")
	fmt.Println("Hello, Every one")
	fmt.Println("Everyone, this is the most awesome class in the entire world for having fun")
	fmt.Println("Having fun and learning go programming language")
	fmt.Println("\n\n\n-----SWEET-----\n")
	//What is control flow?
	//Ans: 3 steps i.e., Sequential Conditional iterative
	foo()
	bar()

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println("VL", i)
		}
	}
	fmt.Println(isPrime(10))
	//exitMethod()

	//fmt.Println("exit")
	exitMethod()
}

func foo() {
	fmt.Println("I am foo...")
}

func bar() {
	fmt.Println("I am bar")
}

func exitMethod() {
	fmt.Println("I am exiting from the code....:(")
}

func isPrime(n int) bool {
	if n <= 1 {
		//fmt.Println("Number must be greater than 2")
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			fmt.Println("Non Prime Number")
			return false
		}
	}
	fmt.Println("Prime Number")
	return true
}

//fmt.Println("ISPRIME", n*10)
//return true

//////////////////////
//
//Prime no -- self and 1 divisibility, no negative no is prime number, 0 and 1 is not a prime number .
// write a function pass a value in that func to check if it is prime or not
// to test it with all conditions -- code coverage must be >90%
//write the logic for prime no, and test the prime method with 90% coverage
////////////////////
