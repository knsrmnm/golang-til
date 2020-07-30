package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("logging")

	log.Fatalln("error!")

	fmt.Println("ok")
}
