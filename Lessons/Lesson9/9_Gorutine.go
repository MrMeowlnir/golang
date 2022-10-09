package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Main stream")
	ch := make(chan int)
	// create new goroutine
	go sayHello(ch)
	for i := range ch {
		fmt.Println(i)
	}
	// selectOne
	fmt.Println("SelectOne")
	data := make(chan int)
	exit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-data)
		}
		exit <- 0
	}()
	selectOne(data, exit)

	fmt.Println("End Main stream")
}

func say(word string) {
	fmt.Println(word)
}

func sayHello(ch chan int) {
	// sleep
	fmt.Println("func sayHello start")
	for t := 0; t < 5; t++ {
		time.Sleep(100 * time.Millisecond)
		ch <- t
		say("Hello!")

	}
	fmt.Println("func sayHello ends")
	close(ch)
}

func selectOne(data, exit chan int) {
	x := 0
	for {
		select {
		case data <- x:
			x += 1
		case <-exit:
			fmt.Println("selectOne exit!")
			return
		}
	}
}
