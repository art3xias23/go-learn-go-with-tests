package main

import (
	"fmt"
	"io"
	"os"
	"time"
)
const goWording = "Go!"

func main(){
	Countdown(os.Stdout, 3, &DefaultSleeper{})
}
func Countdown(writer io.Writer, count int, slp Sleeper){
	for cnt :=count; cnt > 0; cnt--{
	fmt.Fprintln(writer, cnt)
	slp.Sleep()
	}

	fmt.Fprint(writer, goWording)
}

type Sleeper interface {
	Sleep()
}

type SpySleeper struct{
	Calls int
}

type DefaultSleeper struct{}

func(d *DefaultSleeper) Sleep(){
	time.Sleep(1*time.Second)
}

func(s *SpySleeper) Sleep(){
	s.Calls++
}

type SpyCountdownOperations struct{
	Calls []string
}

func(s *SpyCountdownOperations) Write([]byte)(n int, err error){
	s.Calls=append(s.Calls, write)
	return
}

func(s *SpyCountdownOperations) Sleep(){
	s.Calls=append(s.Calls, sleep)
}

const sleep = "sleep"
const write = "write"