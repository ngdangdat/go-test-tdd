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

type countdownSleeper struct{}

func (s *countdownSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	start := 3
	for num := range 3 {
		sleeper.Sleep()
		cd := start - num
		fmt.Fprintf(writer, "%d\n", cd)
	}
	fmt.Fprintf(writer, "Go!")
}

func main() {
	Countdown(os.Stdout, &countdownSleeper{})
}
