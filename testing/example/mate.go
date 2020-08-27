package main

func Sum(x ...int) int {
	var sum int
	for _, v := range x {
		sum += v
	}
	return sum
}
