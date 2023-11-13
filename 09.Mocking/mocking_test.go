package mocking

import (
	"bytes"
	"testing"
)

func Test_Countdown(t *testing.T){
	buffer:= bytes.Buffer{}

	Countdown(&buffer, 3)

	want:=`3
2
1
Go!`
	got:=buffer.String()

	if want!= got{
		t.Errorf("got %q, want %q", got, want)
	}
}