package main

import "fmt"

type rect struct {
	width, height int
}

// Methods on type rect struct
func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 5, height: 10}

	fmt.Println("r.area:", r.area())
	fmt.Println("r.preim:", r.perim())
}
