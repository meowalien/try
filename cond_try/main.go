package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go funcA(wg)
	time.AfterFunc(time.Second, func() {
		cd.L.Lock()
		defer cd.L.Unlock()
		defer wg.Done()

		doIt = true
		cd.Broadcast()
	})
	wg.Wait()
}

var doIt = false
var cd = sync.Cond{L: &sync.Mutex{}}

func funcA(wg *sync.WaitGroup) {
	defer wg.Done()
	cd.L.Lock()
	defer cd.L.Unlock()
	if !doIt {
		cd.Wait()
	}
	fmt.Println("Done")
}
