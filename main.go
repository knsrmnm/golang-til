package main

import "fmt"

// Vertex struct
type Vertex struct {
	X, Y int
}

// Area func : value receiver
func (v Vertex) Area() int {
	return v.X * v.Y
}

// Scale func : pointer receiver
func (v *Vertex) Scale(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
}

// Vertex3D struct
type Vertex3D struct {
	Vertex
	Z int
}

// Area3D func : value receiver
func (v Vertex3D) Area3D() int {
	return v.X * v.Y * v.Z
}

// Scale3D func : pointer receiver
func (v *Vertex3D) Scale3D(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
	v.Z = v.Z * i
}

// New : constructor
func New(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}

func main() {
	v := New(3, 4, 5)
	v.Scale(10)
	fmt.Println(v.Area())
	fmt.Println(v.Area3D())
}
