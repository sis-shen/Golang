package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("The World")
	var s1 string = "hello"
	var s2 string = "world"
	fmt.Println("s1 is", s1, s2)
	greateNum := 8888
	fmt.Println(greateNum)
	const name = "what"
	const (
		one = iota + 1
		two
		threa
	)

	fmt.Println(name)
	fmt.Println(one, two, threa)

	i := 8866
	is2 := strconv.Itoa(i)
	fmt.Println(is2)
	si, err := strconv.Atoi(is2)
	fmt.Println(si, err)
}
