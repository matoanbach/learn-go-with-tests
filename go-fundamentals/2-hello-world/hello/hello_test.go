package hello

import "testing"

func Hello() string {
	return "Hello, world\n"
}

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world\n"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
