package main

import "fmt"

// Vertex struct
type Vertex struct {
	X, Y int
}

// Scale func : pointer receiver
func (v *Vertex) Scale(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
}

// Area func : value receiver
func (v Vertex) Area() int {
	return v.X * v.Y
}

// New : constructor
func New(x, y int) *Vertex {
	return &Vertex{x, y}
}

func main() {
	v := New(3, 4)
	v.Scale(10)
	fmt.Println(v.Area())
}
