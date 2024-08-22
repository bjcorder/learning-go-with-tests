package perimeter

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(Rectangle{10.0, 10.0})
	want := 40.0
	if got != want {
		t.Errorf("got '%.2f' want '%.2f'", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{width: 12, height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{base: 12, height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got '%g' want '%g'", tt.shape, got, tt.hasArea)
			}
		})
	}
	/*
		t.Run("rectangles", func(t *testing.T) {
			rectangle := Rectangle{12.0, 6.0}
			checkValues(t, rectangle, 72.0)
		})
		t.Run("circles", func(t *testing.T) {
			circle := Circle{10.0}
			checkValues(t, circle, 314.1592653589793)
		})
	*/
}

/*
func checkValues(t testing.TB, shape Shape, want float64) {
	t.Helper()
	got := shape.Area()
	if got != want {
		t.Errorf("got '%.2f' want '%.2f'", got, want)
	}
}
*/
