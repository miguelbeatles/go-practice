package main

import (
	"fmt"
)

func main() {
	salir := make(chan int)
	c := gen1(salir)

	recibir1(c, salir)

	fmt.Println("A punto de finalizar.")
}

func gen1(salir chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		salir <- 1
		close(c)
	}()

	return c
}

func recibir1(c, salir <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println("valor:", v)
		case <-salir:
			return
		}
	}

}
