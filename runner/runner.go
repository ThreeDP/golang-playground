package runner

import (
	"fmt"
	"net/http"
	"time"
)

const limitOfTenSec = 10 * time.Second

func Runner(a, b string) (winner string, err error) {
	return ConfigRunner(a, b, limitOfTenSec)
}

func ConfigRunner(a, b string, limitTime time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(limitTime):
		return "", fmt.Errorf("Exceeded wait timeout for %s and %s", a, b)
	}
}

func ping(URL string) (chan bool) {
	ch := make(chan bool)
	go func() {
		http.Get(URL)
		ch <- true
	}()
	return ch
}

/* execução da validação iterativa
func timeToResponse(URL string) time.Duration {
	init := time.Now()
	http.Get(URL)
	return time.Since(init)
}

func Runner(a, b string) (winner string) {
	endA := timeToResponse(a)
	endB := timeToResponse(b)
	if endA < endB {
		return a
	}
	return b
}
*/