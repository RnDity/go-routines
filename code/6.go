package main

import (
	"sync"
	"io/ioutil"
	"time"
)

// dd if=/dev/urandom of=/tmp/bigfile bs=10M count=3
func work(ch chan string, wg *sync.WaitGroup) {
	for file := range ch {
		dat, _ := ioutil.ReadFile(file)
		if len(dat) != 30*1024*1024 {
			panic("file size not match")
		}
		wg.Done()
	}
}
// START OMIT
func run(numRoutines int) {
	wg := &sync.WaitGroup{}
	wg.Add(100)
	ch := make(chan string, 100)
	for i := 0; i < 100; i++ {
		ch <- "/tmp/bigfile"
	}
	for i := 0; i < numRoutines; i++ {
		go work(ch, wg)
	}
	wg.Wait()
	close(ch)

}

func main() {
	routines := []int{1, 4, 6, 7, 8, 9, 10, 15, 20, 50, 100}
	for _, r := range routines {
		start := time.Now()
		run(r)
		println(r, time.Since(start).Nanoseconds()/1000000, "ms")
	}
}
// END OMIT