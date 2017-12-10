package main

import (
	"time"
	"net/http"
)

func say(s string, ch chan string) {
	time.Sleep(100 * time.Millisecond)
	ch <- s
}

func main() {
	ch := make(chan string)
	go say("hello", ch)
	go say("world", ch)
	println(<-ch)
	println(<-ch)
}
