package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	alarm string
	count int
)

func init() {
	flag.StringVar(&alarm, "a", "Alarm", "Gives back a number of alarms")
	flag.IntVar(&count, "c", 0, "Give a number to print alarm")
	flag.Parse()
}

func main() {
	print()
}

func print() {
	for i := 1; i <= count; i++ {
		fmt.Println(alarm, i, "!")
		time.Sleep(1 * time.Second)
	}
}
