package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(3)

	go say(wg, "Hello1", time.Second * 3)
	go say(wg, "Hello2", time.Second * 2)
	go say(wg, "Hello3", time.Second * 1)

	wg.Wait()
}

func say(wg *sync.WaitGroup, s string, duration time.Duration) {
	time.Sleep(duration)
	fmt.Println(s)
	wg.Done()
}
