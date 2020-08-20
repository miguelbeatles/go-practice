package main

import "fmt"

type persona struct {
	nombre string
	edad   int
}

type humano interface {
	hablar()
}

func (p *persona) hablar() {
	fmt.Println("Hola mi nombre es:", p.nombre)
}

func diAlgo(h humano) {
	h.hablar()

}

func main() {

	p1 := persona{
		nombre: "Miguel",
		edad:   27,
	}
	diAlgo(&p1)

}
