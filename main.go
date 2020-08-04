package main

import "fmt"

// MyInt : struct
type MyInt int

// Double func
func (i MyInt) Double() int {
	return int(i * 2)
}

func main() {
	myInt := MyInt(10)
	fmt.Println(myInt.Double())
}
