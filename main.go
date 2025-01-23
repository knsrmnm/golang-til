package main

import (
	"fmt"
	"os/user"
	"time"
)

func init() {
	fmt.Println("Init")
}

func main() {
	fmt.Println("Hello world", time.Now())
	fmt.Println(user.Current())
}
