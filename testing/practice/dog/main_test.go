package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	n := Years(20)
	if n != 140 {
		t.Error("Expected", 140, "Got", n)
	}
}

func TestYearsTwo(t *testing.T) {
	n := YearsTwo(10)
	if n != 70 {
		t.Error("Expected", 70, "Got", n)
	}
}

func BenchmarkYear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(111)
	}
}

func BenchmarkYearTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(111)
	}
}
func ExampleYears() {
	fmt.Println(Years(10))
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(10))
}
