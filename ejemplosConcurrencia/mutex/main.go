package main

/*Con este ejemplo nos aseguramos de que no tengamos race en nestras rutinas*/

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Los CPUs que tengo en mi compu es:", runtime.NumCPU())
	fmt.Println("Numero de gorutinas que corres actualmente es:", runtime.NumGoroutine())

	count := 0
	const gr = 50
	var goRutinas sync.WaitGroup
	var mutex sync.Mutex
	goRutinas.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			/*con el metodo de mutex Lock bloqueamos la variable final para que no escriban al mismo tiempo las rutinas en dicha variable*/
			mutex.Lock()
			final := count
			/*ya no se requiere ya que nosostros controlaremos la concurrencia con mutex*/
			//runtime.Gosched()
			final++
			count = final
			goRutinas.Done()
			/*Con mutex Unlock liberamos la variable para de mas rutinas*/
			mutex.Unlock()
		}()
		fmt.Println("Numero de gorutinas que corres actualmente es:", runtime.NumGoroutine())
	}

	goRutinas.Wait()

	fmt.Println("el valor final es:", count)

}
