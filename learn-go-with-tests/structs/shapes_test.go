package structs

import "testing"

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t *testing.T,shape Shape, expected float64) {
		got := shape.Perimeter()
		if expected != got {
			t.Errorf("expected: %.2f, but got: %.2f", expected, got)
		}
	}
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		checkPerimeter(t, rectangle, 40.0)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		checkPerimeter(t, circle, 62.83185307179586476925286766559)
	})
}

func TestArea(t *testing.T) {
	
	areaTests := []struct {
		shape Shape
		expected float64
	}{
		{Rectangle{12.0, 6.0}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, st := range areaTests {
		got := st.shape.Area()
		if st.expected != got {
			t.Errorf("expected: %.2f, but got: %.2f", st.expected, got)
		}
	}
}