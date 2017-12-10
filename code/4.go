package main

import (
	"time"
	"sync"
)

func say(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
	println(<-ch)
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan string)
	go say(ch, wg)
	go say(ch, wg)
	ch <- "hello"
	ch <- "world"
	wg.Wait()
}
