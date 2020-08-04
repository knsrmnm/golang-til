package main

import "fmt"

type Vertex struct {
	X, Y int
	S    string
}

func main() {
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)
	v.X = 100
	fmt.Println(v.X, v.Y)

	var v2 Vertex
	fmt.Println(v2)
}
