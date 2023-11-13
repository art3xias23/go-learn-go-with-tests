package main

import (
	"bytes"
	"testing"
)

func Test_Countdown(t *testing.T){
	buffer:= bytes.Buffer{}
spySleep:= SpySleeper{}
cnt:=3 
	Countdown(&buffer, cnt, &spySleep)

	want:=`3
2
1
Go!`
	got:=buffer.String()

	if want!= got{
		t.Errorf("got %q, want %q", got, want)
	}

	if cnt!=spySleep.Calls{
		t.Errorf("SpySleep: got %d, wanted %d", spySleep.Calls, cnt)
	}
}