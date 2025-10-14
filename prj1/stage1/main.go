package main

import (
	"errors"
	"fmt"
)

//func main() {
//	fmt.Println("The World")
//	var s1 string = "hello"
//	var s2 string = "world"
//	fmt.Println("s1 is", s1, s2)
//	greateNum := 8888
//	fmt.Println(greateNum)
//	const name = "what"
//	const (
//		one = iota + 1
//		two
//		threa
//	)
//
//	fmt.Println(name)
//	fmt.Println(one, two, threa)
//
//	i := 8866
//	is2 := strconv.Itoa(i)
//	fmt.Println(is2)
//	si, err := strconv.Atoi(is2)
//	fmt.Println(si, err)
//}

//func main() {
//	if i := 11; i > 10 {
//		fmt.Println("i is greater than 10")
//	} else if i > 5 && i <= 10 {
//		fmt.Println("i is greater than 5 , less than 10")
//	} else {
//		fmt.Println("i is less than 5")
//	}
//
//	nameAge := make(map[string]int)
//	nameAge["James"] = 21
//	fmt.Println(nameAge["James"])
//}

func sum(a, b int) int {
	return a + b
}

func unsignedSum(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("invalid arguments")
	}
	return a + b, nil
}

func sumAll(params ...int) int {
	sum := 0
	for _, param := range params {
		sum += param
	}
	return sum
}

type Age uint8

func (age Age) Show() {
	fmt.Println("the age is", age)
}

func main() {
	fmt.Println(sum(1, 2))
	res, _ := unsignedSum(1, 2)
	fmt.Println(res)

	c1 := colsure()
	fmt.Println(c1())
	fmt.Println(c1())
	fmt.Println(c1())

	c2 := colsure()
	fmt.Println(c2())
	fmt.Println(c2())
	fmt.Println(c2())
	fmt.Println(c2())

	//var a1 Age = 88
	a1 := Age(88)
	a1.Show()

}

func colsure() func() int {
	sti := 0
	return func() int {
		sti++
		return sti
	}

}
