package loops

import "fmt"

func main() {
	// simple for
	var sum = 0
	for index := 0; index < 5; index++ {
		sum += index
	}

	// short for <-> while
	for sum < 1000 {
		sum += 10
	}

	fmt.Println(sum)

	var a = 100

	// if
	if a == 100 {
		fmt.Println(a)
	}

	// switch
	switch a {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("default")
	}

	// defer, stack operations (last in -> first out)
	defer fmt.Println("defer op 1")
	defer fmt.Println("defer op 2")
	defer fmt.Println("defer op 3")
	fmt.Println("last op")

}

func first() int {
	var a = 1
	return a
}

type s struct {
	a string
	b string
	c string
}

func second() s {
	var a = "str1"
	var b = "str2"
	var c = "str3"

	return s{a, b, c}
}
