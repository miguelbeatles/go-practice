package main

/*en este ejemplo utilizamos atomic para solucionar el race*/

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("Los CPUs que tengo en mi compu es:", runtime.NumCPU())
	fmt.Println("Numero de gorutinas que corres actualmente es:", runtime.NumGoroutine())

	var count int64
	const gr = 50
	var goRutinas sync.WaitGroup
	/*con add de le decimos a main el numero de gorutinas que debe esperar para finalizar*/
	goRutinas.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			atomic.AddInt64(&count, 1)
			goRutinas.Done()
		}()
		fmt.Println("Numero de gorutinas que corres actualmente es:", runtime.NumGoroutine())
	}
	/*wait es para que la funcion espere a tods las subrutinas que semandaron en dd*/
	goRutinas.Wait()

	fmt.Println("el valor final es:", count)

}
