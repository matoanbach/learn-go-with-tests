package perimeter

import "testing"

func TestPerimeter(t *testing.T) {
	rec := Rectangle{10.0, 10.0}
	got := rec.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rec := Rectangle{12.0, 6.0}
	got := rec.Area()
	want := 72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
