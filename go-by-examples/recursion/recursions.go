package main

import "fmt"

func factorial(n int) int {

	for n != 0 {
		return n * factorial(n-1)
	}
	return 1
}

func main() {
	fmt.Println("factorial-7: ", factorial(7))
}
