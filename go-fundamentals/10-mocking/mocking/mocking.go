package mocking

import "fmt"

func Countdown() {
	for _, s := range []string{"Print ", "3, ", "2, ", "1, ", "and ", "Go!", "\n"} {
		fmt.Print(s)
	}
}
