package main

import (
	"bytes"
	"reflect"
	"testing"
)

func Test_Countdown(t *testing.T){
	t.Run("Test normal output", func(t *testing.T){

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
	})

	t.Run("Test sequence output", func(t *testing.T){
		want:=[]string{write,
					sleep,
					write,
					sleep,
					write,
					sleep,
					write,
}
spySleep:=SpyCountdownOperations{}
Countdown(&spySleep, 3, &spySleep)
got:= spySleep.Calls

if !reflect.DeepEqual(want, got){
	t.Errorf("got %v, want %v", got, want)
}
	})

}