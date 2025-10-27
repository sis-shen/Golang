package main

import (
	"fmt"
)

func Fib(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}

	return Fib(n-1) + Fib(n-2)
}

func main() {
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	for i := 1; i <= 5; i++ {
		fmt.Println("i =", 100/i)
	}
}
