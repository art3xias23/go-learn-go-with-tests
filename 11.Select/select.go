package selecting

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string){
timea := measureResponseTime(a)
timeb := measureResponseTime(b)

	if timea < timeb{
		return a
	}
	return b
}

func measureResponseTime(url string) time.Duration{
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}