package main

import "fmt"

func main() {
	par := make(chan int)
	impar := make(chan int)
	salida := make(chan int)

	go enviar(par, impar, salida)

	recibir(par, impar, salida)

	fmt.Println("ya estuvo")
}

func enviar(p chan<- int, im chan<- int, s chan<- int) {

	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			p <- i
		} else {
			im <- i
		}
	}
	close(p)
	close(im)
	s <- 0

}

/*Con select podemos sacar la informacion de los canales funciona como case*/

func recibir(p <-chan int, im <-chan int, s <-chan int) {
	for {
		select {
		case v := <-p:
			fmt.Println("desde el canal par", v)
		case v := <-im:
			fmt.Println("desde el canal ipar", v)
		case v := <-s:
			fmt.Println("desde el canal salir", v)
			return
		}
	}
}
