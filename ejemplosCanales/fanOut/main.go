package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*Utilizando paron de diseño Fan Out con canales esto consiste  en dicidir todo el trabajo en varias rutinas*/
func main() {
	a := make(chan int)
	b := make(chan int)

	go agregar(a)
	go fanOutIn(a, b)

	for v := range b {
		fmt.Println(v)
	}

}

func agregar(a chan int) {
	for i := 0; i <= 100; i++ {
		a <- i
	}
	close(a)
}

func fanOutIn(a, b chan int) {
	var wg sync.WaitGroup
	const rutinas = 10
	wg.Add(rutinas)
	for i := 0; i <= rutinas; i++ {
		go func() {
			for v := range a {
				func(v1 int) {
					b <- tiempo(v1)
				}(v)
			}
			wg.Done()
		}()
		wg.Wait()
		close(b)
	}
}

func tiempo(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
