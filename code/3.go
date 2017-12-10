package main

import (
	"time"
	"sync"
)

func say(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
	println(s)
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go say("hello", wg)
	go say("world", wg)
	wg.Wait()
}
