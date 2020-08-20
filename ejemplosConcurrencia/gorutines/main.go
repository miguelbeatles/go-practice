package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	/*WaitGroup sirve para lanzar procesos concurrentes*/
	var wg sync.WaitGroup
	fmt.Println("numero de cpus", runtime.NumCPU())
	fmt.Println("numero de rutinas", runtime.NumGoroutine())
	wg.Add(2)
	/*Primera gorutina*/
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		wg.Done()
	}()
	/*segunda gorutina*/
	go func() {

		for i := 9; i > 0; i-- {
			fmt.Println(i)
		}
	}()
	wg.Done()
	fmt.Println("numero de rutinas", runtime.NumGoroutine())
	wg.Wait()

	fmt.Println("numero de rutinas", runtime.NumGoroutine())
	fmt.Println("Finalizo con exito")

}
