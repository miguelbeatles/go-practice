package main

import "fmt"

func main() {

	salir := make(chan int)

	c := llenar(salir)
	imprimir(c, salir)

}

func llenar(s chan<- int) <-chan int {
	c := make(chan int)
	go func() {

		for i := 0; i < 100; i++ {
			c <- i
		}
		s <- 0
		close(c)
	}()
	return c
}

func imprimir(c, s <-chan int) {

	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-s:
			return
		}
	}

}
