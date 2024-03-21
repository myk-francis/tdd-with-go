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

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep()  {
	s.Calls++
}

type DefaultSleeper struct {}

func (s *DefaultSleeper) Sleep()  {
	time.Sleep(1 * time.Second)
}

func CountDown(out io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
	}

	time.Sleep(1 * time.Second)
	fmt.Fprint(out, "Go!")
}

func main()  {
	sleeper := &DefaultSleeper{}
	CountDown(os.Stdout, sleeper)
}