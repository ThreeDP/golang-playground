package myintegers

import (
	"testing"
	"fmt"
)

func TestAdd(t *testing.T) {
	sum := add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d', result '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
