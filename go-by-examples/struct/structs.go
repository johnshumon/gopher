package main

import "fmt"

// TwoD struct
type TwoD struct {
	x, y int
}

// Person struct
type Person struct {
	firstName string
	lastName  string
	email     string
	age       int
}

// Human struct
type Human struct {
	name string
	age  int
}

// NewPerson struct constructor
func NewPerson(fn, ln, em string) Person {

	p := Person{
		firstName: fn,
		lastName:  ln,
		email:     em,
	}
	p.age = 27

	return p
}

// NewHuman strcut constuctor
func NewHuman(name string) *Human {

	p := Human{name: name}
	p.age = 42
	return &p
}

var (
	t1 = TwoD{11, 13}
	t2 = TwoD{x: 2}
	t3 = TwoD{y: 3}
	t4 = TwoD{}
	p  = &TwoD{1, 2}
)

func main() {
	s1 := Person{
		firstName: "abu",
		lastName:  "shumon",
		email:     "abu@shumon.me",
		age:       30,
	}

	s2 := NewPerson("mihi", "khanam", "mimi@mimi.com")

	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)

	// ------
	fmt.Println(t1, t2, t3, t4, p)

	// ------
	fmt.Println(Human{"Bob", 20})
	fmt.Println(Human{name: "Alice", age: 30})
	fmt.Println(Human{name: "Fred"})
	fmt.Println(&Human{name: "Ann", age: 40})
	fmt.Println(NewHuman("John"))

	s := Human{name: "Sean", age: 50}
	fmt.Println(s.name, s.age)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
