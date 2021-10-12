package main

import "fmt"

func main() {
	i := 3

	p := &i // point to i
	fmt.Println("p:", *p)

	*p = 13 // set i through pointer p
	fmt.Println("p:", *p)

}
