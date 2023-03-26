package hello

import (
	"testing"
	"bytes"
)

func TestGreet(t *testing.T) {
	buf := bytes.Buffer{}
	Greet(&buf, "Davy")
	result := buf.String()
	expected := "Hello, Davy"
	if expected != result {
		t.Errorf("result '%s', expected '%s'", result, expected)
	}
}