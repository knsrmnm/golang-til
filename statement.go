package main

import "fmt"

func main() {
	num := 5
	if num%2 == 0 {
		fmt.Println("by 2")
	} else if num%3 == 0 {
		fmt.Println("by 3")
	} else {
		fmt.Println("else")
	}

	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}
		if i > 5 {
			break
		}
		fmt.Println(i)
	}

	l := []string{"python", "go", "java"}
	for i, v := range l {
		fmt.Println(i, v)
	}

	os := "mac"
	switch os {
	case "mac":
		fmt.Println("Mac!")
	case "windows":
		fmt.Println("windows!")
	}
}
