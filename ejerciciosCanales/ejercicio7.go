package main

import (
	"fmt"
	"sync"
)

func main() {

	c := make(chan int)
	var wg1 sync.WaitGroup
	var numGorutines = 10
	wg1.Add(numGorutines)

	for i := 0; i <= numGorutines; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				c <- j
			}
		}()
	}
	for i := 0; i < 100; i++ {
		fmt.Println(i, <-c)
	}

}
