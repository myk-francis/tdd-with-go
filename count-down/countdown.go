package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

func CountDown(out io.Writer) {
	for i := 3; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
	}

	time.Sleep(1 * time.Second)
	fmt.Fprint(out, "Go!")
}

func main()  {
	CountDown(os.Stdout)
}