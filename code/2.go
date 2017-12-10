package main

import (
	"time"
)

func say(s string) {
	time.Sleep(100 * time.Millisecond)
	println(s)
}

func main() {
	go say("hello")
	go say("world")
	time.Sleep(150 * time.Millisecond)
}
