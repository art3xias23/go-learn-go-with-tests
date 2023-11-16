package selecting

import (
	"fmt"
	"net/http"
	"time"
)

func RacerTime(a, b string) (winner string){
timea := measureResponseTime(a)
timeb := measureResponseTime(b)

	if timea < timeb{
		return a
	}
	return b
}

func RacerSelect(a, b string) (winner string){
	select{
	case <- ping(a):
		return a
	case <- ping(b):
		return b
	}
}

func Racer(a, b string) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
		case <- time.After(10 * time.Second):
			return "", fmt.Errorf("timed out waiting on %s and %s", a, b)

	}
}


func measureResponseTime(url string) time.Duration{
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func ping(url string) chan struct{}{
	ch:=make(chan struct{})
	go func(){
		http.Get(url)
		close(ch)
	}()
	return ch
}