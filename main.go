package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello" + "World")
	var s  string = "Hello World"
	fmt.Println(strings.Replace(s, "H", "X", 1))
}
