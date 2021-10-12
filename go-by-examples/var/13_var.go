package main

import "fmt"

// var statement at package level.
var a, b = 1, 2

func main() {

	// var statement at function level.
	var i, j int = 1, 2
	var c, python, java = true, 3, "no way!"

	// short assignment statement
	// with implicit type
	k := 4

	// different type of variables
	// declared using a single var keyword
	var (
		name = "John"
		age  = 33
	)

	fmt.Println(i, j, c, k, python, java)
	fmt.Println("I'm", name, "I'm", age, "years old")
}
