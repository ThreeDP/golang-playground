package iter

import (
	"testing"
	"fmt"
)

func TestIter(t *testing.T) {
	rep := Iter("a", 10)
	expected := "aaaaaaaaaa"

	if rep != expected {
		t.Errorf("expected '%s', result '%s'", expected, rep)
	}
}

func ExampleIter() {
	s := Iter("a", 10)
	fmt.Println(s)
	// Output: aaaaaaaaaa
}

func BenchmarkIter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iter("a", 5)
	}
}