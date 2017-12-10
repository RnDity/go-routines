package main

import (
	"time"
)

func say(s string) {
	time.Sleep(100 * time.Millisecond)
	println(s)
}

func main() {
	say("hello")
	say("world")
}
