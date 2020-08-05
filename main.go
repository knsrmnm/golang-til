package main

import "fmt"

// Human : struct
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("invalid type %T\n", v)
	}
}

func main() {
	do(10)
	do("Mike")
	do(true)
}
