package main

import (
	"fmt"
	"io"
	"os"
)

type Sleeper interface {
	Sleep()
}

func CountDown(out io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
	}

	fmt.Fprint(out, "Go!")
}

func main()  {
	CountDown(os.Stdout)
}