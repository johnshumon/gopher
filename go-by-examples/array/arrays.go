package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len= %d cap= %d %v\n", len(s), cap(s), s)
}

func main() {
	// array of integer
	var arr [10]int

	// assigning value in array index
	arr[0] = 10
	arr[9] = 9

	fmt.Println(arr)

	// array of string and values assigned to it
	strArr := [3]string{"hello", "world", "!"}
	fmt.Println(strArr)

	sliceArr := [10]int{11, 22, 33, 44, 55, 66, 77, 88, 99, 100}

	// slice of sliceArr
	sla := sliceArr[5:9]

	slice := []int{2, 3, 5, 7, 9, 11, 13, 17, 19}
	printSlice(slice)

	// slice the slice
	slice = slice[:0]
	printSlice(slice)

	// slice the slice
	slice = slice[:5]
	printSlice(slice)

	fmt.Println(sla)
}
