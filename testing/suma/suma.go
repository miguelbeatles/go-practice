package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7}
	a := mySUM(x...)
	fmt.Println("La suma es: ", a)
}
func mySUM(x ...int) int {
	var sum int
	for _, v := range x {
		sum += v
	}
	return sum
}
