package main

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println("counter :", counter)
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	{
		counter++
	}
	mutex.Unlock()
}
