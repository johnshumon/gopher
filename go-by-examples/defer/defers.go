package main

import "fmt"

func main() {
	fmt.Println("count")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
