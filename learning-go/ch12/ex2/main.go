package main

import (
	"fmt"
)

func printNumbers() {
	c1 := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			c1 <- i
		}
		close(c1)
	}()

	c2 := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			c2 <- i
		}
		close(c2)
	}()

	for {
		select {
		case val, ok := <-c1:
			if !ok {
				c1 = nil
				break
			}
			fmt.Println(val)
		case val, ok := <-c2:
			if !ok {
				c2 = nil
				break
			}
			fmt.Println(val)
		}

		if c1 == nil && c2 == nil {
			break
		}
	}
}

func main() {
	printNumbers()
}
