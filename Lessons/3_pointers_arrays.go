package pointers_arrays

import "fmt"

func main() {
	fmt.Println(slice())
}

// Arrays
func arrays() [10]int {
	var a [10]int
	for i := 0; i < 10; i++ {
		a[i] = i
	}
	return a
}

// Slice
func slice() []int {
	var a []int
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	return a
}

// Pointers
func pointers() {
	var hello string = "Hello, World!"

	p := &hello
	fmt.Println("Address =", p)
	fmt.Println("Value =", *p)
	*p = "New value"
	fmt.Println("New value =", hello)
}

// Structures
type Point struct {
	X int
	Y int
}

func (p *Point) get_values() {
	fmt.Println(&p, "(", p.X, ",", p.Y, ")")
}

func Struct() {
	var struct1 = Point{
		X: 1,
		Y: 2,
	}

	fmt.Println(struct1)

	pointer1 := &struct1
	fmt.Println(pointer1)
}
