package perimeter

import (
	"testing"
)

type Shape interface {
	Area() float64
}

func TestPerimeter(t *testing.T) {
	r := Rectangle{10.0, 10.0}
	result := r.Perimeter()
	expected := 40.0
	if result != expected {
		t.Errorf("result %.2f expected %.2f", result, expected)
	}
}

/* V1 da implmentação de tests com interface
func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, expected float64) {
		t.Helper()
		result := shape.Area()
		if result != expected {
			t.Errorf("result %.2f expected %.2f", result, expected)
		}
	}
	t.Run("Rectangle test area", func(t *testing.T) {
		r := Rectangle{12.0, 6.0}
		expected := 72.0
		checkArea(t, r, expected)
	})
	t.Run("Circle test area", func(t *testing.T) {
		c := Circle{10}
		expected := 314.1592653589793
		checkArea(t, c, expected)
	})
}
*/

func TestArea(t *testing.T) {
	testsArea := []struct {
		name		string
		shape 		Shape
		expected	float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, expected: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, expected: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, expected: 36.0},
	}
	for _, tt := range testsArea {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.shape.Area()
			if result != tt.expected {
				t.Errorf("%#v result %.2f, expected %.2f", tt.shape, result, tt.expected)
			}
		})
	}
}
