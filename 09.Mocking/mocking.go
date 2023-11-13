package mocking

import (
	"fmt"
	"io"
	"os"
)
const goWording = "Go!"

func main(){
	Countdown(os.Stdout, 3)
}
func Countdown(writer io.Writer, count int){
	for cnt :=count; cnt > 0; cnt--{
	fmt.Fprintln(writer, cnt)
	}

	fmt.Fprint(writer, goWording)
}