package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Counter struct {
	mu sync.Mutex
	c  map[string]int
}

func (c *Counter) Inc(key string) {
	c.mu.Lock()
	c.c[key]++
	c.mu.Unlock()
}

func (c *Counter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.c[key]
}

func main() {
	key := "test"
	c := Counter{c: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc(key)
	}
	time.Sleep(1 * time.Second)
	fmt.Println(c.Value(key))

	// waitgroup
	var wg sync.WaitGroup
	var counter uint64

	for i := 0; i < 10; i++ {
		wg.Add(1)
		n := i
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddUint64(&counter, 1)
			}
			fmt.Printf("%d goroutine working...\n", n)
			time.Sleep(300 * time.Millisecond)
		}()
	}

	wg.Wait()
	fmt.Printf("All done. Counter = %d\n", counter)
}
