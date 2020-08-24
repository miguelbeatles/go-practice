package main

import (
	"fmt"
	"sync"
)

/*Utilizando paron de diseño Fan In con canales esto consiste en meter informacion de varios canales en uno*/
func main() {
	par := make(chan int)
	impar := make(chan int)
	fanin := make(chan int)

	go enviar(par, impar)
	go recibir(par, impar, fanin)

	for v := range fanin {
		fmt.Println(v)
	}
}

func enviar(p, i chan<- int) {
	for j := 0; j <= 100; j++ {
		if j%2 == 0 {
			p <- j
		} else {
			i <- j
		}
	}
	close(p)
	close(i)
}

func recibir(p, i <-chan int, f chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for v := range p {
			f <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range i {
			f <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(f)
}
