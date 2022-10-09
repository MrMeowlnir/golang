package main

import (
	"fmt"
	"log"
)

func main() {
	//	a := []int{1, 2, 3}
	//	c := a[3]
	//	panic("AAAAAAAAAAAAAA!")
	divide(1, 0)
	fmt.Println("Code after panic...")
}

func divide(a, b int) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Panic!", err)
		}
	}()
	fmt.Println(a / b)
}
