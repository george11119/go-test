package main

import (
	"fmt"
	"sync"
)

func printNumbers() {
	c := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	for range 2 {
		go func() {
			defer wg.Done()

			for i := 1; i <= 10; i++ {
				c <- i
			}
		}()
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	var wg2 sync.WaitGroup
	wg2.Add(1)

	go func() {
		defer wg2.Done()
		for v := range c {
			fmt.Println(v)
		}
	}()

	wg2.Wait()
}

func main() {
	printNumbers()
}
