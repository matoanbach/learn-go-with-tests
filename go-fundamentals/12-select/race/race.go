package race

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	if measureResponseTime(a) < measureResponseTime(b) {
		return a
	}
	return b
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
