package main

import "testing"

func TestMysum(t *testing.T) {
	type test struct {
		datos     []int
		respuesta int
	}

	tests := []test{
		test{[]int{1, 1, 1, 1}, 4},
		test{[]int{2, 4, 0, 1}, 7},
		test{[]int{3, 2, 0, 1}, 6},
		test{[]int{5, 2, 1, 1}, 9},
	}

	for _, x := range tests {
		v := mySUM(x.datos...)
		if v != x.respuesta {
			t.Error("Expected", x.respuesta, "Got", mySUM(x.datos...))
		}
	}

}
