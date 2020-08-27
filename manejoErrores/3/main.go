package main

import (
	"fmt"
)

type errorPer struct {
	id  int
	err string
}

func (ep errorPer) Error() string {
	return fmt.Sprintf("aqui esta el error %v", ep.err)

}

func main() {
	e1 := errorPer{
		id:  21,
		err: "my own error",
	}
	foo(e1)
}

func foo(e error) {
	fmt.Println("foo corrio", e, "\n")
}
