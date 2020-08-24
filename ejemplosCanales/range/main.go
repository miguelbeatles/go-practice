package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		/*ocupamos la funcion de close para cerrar el canal debido a que si no los
		hacemos seguira  extraeyendo valores y se va a bloquear */
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("finalizando")

}

func enviar(c chan<- int) {
	c <- 21
}

func recibir(c <-chan int) {
	fmt.Println(<-c)
}
