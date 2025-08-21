package print

import (
	"bytes"
	"fmt"
)

func Greet(writer *bytes.Buffer, name string) (n int, err error) {
	// fmt.Printf("Hello, %s", name)
	n, err = fmt.Fprintf(writer, "Hello, %s", name)
	if err != nil {
		return -1, err
	}
	return n, err
}
