package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	sync.Mutex
}

func (this *Counter) Increment() {
	this.Lock()
	{
		this.count++
	}
	this.Unlock()
}

var counter Counter

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println("counter :", counter.count)
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	counter.Increment()
}
