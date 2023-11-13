package main

import (
	"fmt"
	"io"
	"os"
)
const goWording = "Go!"

func main(){
	Countdown(os.Stdout, 3, &SpySleeper{})
}
func Countdown(writer io.Writer, count int, slp Sleeper){
	for cnt :=count; cnt > 0; cnt--{
	fmt.Fprintln(writer, cnt)
	slp.Sleep()
	//time.Sleep(1*time.Second)
	}

	fmt.Fprint(writer, goWording)
}

type Sleeper interface {
	Sleep()
}

type SpySleeper struct{
	Calls int
}

func(s *SpySleeper) Sleep(){
	s.Calls++
}