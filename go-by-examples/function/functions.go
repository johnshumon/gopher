package main

import "fmt"

// Syntax for declaring a function:
// func functionname(parametername type) returntype {
//function body
func add(x int, y int) int {
	return x + y
}

// Add_same_type: x and y share same type.
// that's x type is omitted.
func add_same_type(x, y int) int {
	return x + y
}

// Swap: multiple return types
func swap(x, y string) (string, string) {
	return y, x
}

// Spilt: returned values can be named.
// as return statement is without arguments
// therefore, spilt function returns
// named return values.
func spilt(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	return
}

func main() {
	fmt.Println(add(17, 13))
	fmt.Println(add_same_type(13, 13))
	fmt.Println(swap("world", "hello"))
	fmt.Println(spilt(99))
}
