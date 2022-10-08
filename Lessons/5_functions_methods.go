package main

import "fmt"

// callback fucn
func dosmth(callback func(int, int) int, s string) int {
	result := callback(5, 1)
	fmt.Println(s)
	return result
}

// Closures
func totalPrice(initPrice int) func(int) int {
	sum := initPrice
	return func(x int) int {
		sum += x
		return sum
	}
}

// methods
type Point struct {
	X, Y int
}

func (p *Point) movePoint(dx, dy int) {
	p.X += dx
	p.Y += dy
}

func main() {
	// callbacs works
	fmt.Println("Callbacks")
	sumCallback := func(n1, n2 int) int {
		return n1 + n2
	}
	res := dosmth(sumCallback, "plus")
	fmt.Println(res)

	multCallback := func(n1, n2 int) int {
		return n1 * n2
	}
	res = dosmth(multCallback, "multiple")
	fmt.Println(res)

	// closures works
	fmt.Println("Closures")
	orderPrice := totalPrice(1)
	fmt.Println(orderPrice(1))
	fmt.Println(orderPrice(1))

	// methods
	fmt.Println("Methods")
	p := Point{1, 2}
	fmt.Println(p)
	p.movePoint(1, 1)
	fmt.Println(p)
}
