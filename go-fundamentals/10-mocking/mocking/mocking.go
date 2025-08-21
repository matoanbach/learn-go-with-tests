package mocking

import (
	"fmt"
	"io"
)

func Countdown(w io.Writer) {
	for _, s := range []string{"Print ", "3, ", "2, ", "1 ", "and ", "Go!"} {
		fmt.Fprintf(w, "%s", s)
	}
}
