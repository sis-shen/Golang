package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

type Person struct {
	name string
	age  int
}

type Shower interface {
	showAge() int
}

func (person Person) showAge() int {
	fmt.Println(person.age)
	return person.age
}

type commonError struct {
	message string
	code    int
}

func (ce *commonError) Error() string {
	return ce.message
}

func add(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, &commonError{message: "<UNK>", code: -1}
	}
	return a + b, nil
}

func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return io.ReadAll(f)
}

func main() {
	p1 := Person{"Joe", 87}
	p2 := Person{age: 30, name: "pupu"}
	p1.showAge()
	func_sa := Person.showAge
	func_sa(p2)

	sum, err := add(-1, 2)
	if cm, ok := err.(*commonError); ok {
		fmt.Println(cm.code)
	} else {
		fmt.Println(sum)
	}

	e := errors.New("this is an error")
	we := fmt.Errorf("this is a weird error, %s", e)
	fmt.Println(we)
	fmt.Println(errors.Unwrap(we))
	fmt.Println(errors.Is(we, e))
}
