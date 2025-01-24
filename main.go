package main

import (
	"fmt"
	"os/user"
	"time"
)

const Pi = 3.14

func init() {
	fmt.Println("Init")
}

func main() {
	fmt.Println("Hello world", time.Now())
	fmt.Println(user.Current())
	fmt.Println(Pi)
}
