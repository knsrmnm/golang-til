package main

import (
	"fmt"
	"strconv"
	"time"
)

// constants
const (
	Username = "Ken Minami"
	ID       = "12345"
)

func foo() {
	fmt.Println(Username)
	fmt.Println(ID)
}

func bar() {
	var c []int
	//c = make([]int, 5)
	c = make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)
}

func strToInt() {
	var s string = "14"
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v\n", i, i)
}

func main() {
	var i int = 1
	fmt.Println(time.Now())
	fmt.Println(i)

	xi := 1
	fmt.Println(xi)

	foo()

	strToInt()

	bar()
}
