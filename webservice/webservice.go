package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
	_ "net/http/pprof"
)

func countBests(url string, ch chan int) {
	res, err := http.Get(url)
	if err != nil || res.StatusCode != http.StatusOK {
		ch <- 0
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	count := 0
	for _, s := range strings.Split(string(body), " ") {
		if strings.HasPrefix(strings.ToLower(s), "naj") {
			count++
		}
	}
	ch <- count
}

func handler(w http.ResponseWriter, _ *http.Request) {
	ch := make(chan int)
	go countBests("http://gazeta.pl/", ch)
	go countBests("http://pudelek.pl/", ch)
	fmt.Fprintf(w, "%v", <-ch + <-ch)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

// curl -s http://localhost:8080/debug/pprof/trace?seconds=2 > trace.out
// PPROF_BINARY_PATH=. go tool pprof --gv -alloc_objects http://localhost:8080/debug/pprof/heap

