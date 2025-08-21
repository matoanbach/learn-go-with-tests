package main

import (
	"main/mocking"
	"os"
)

func main() {
	mocking.Countdown(os.Stdout)
}
