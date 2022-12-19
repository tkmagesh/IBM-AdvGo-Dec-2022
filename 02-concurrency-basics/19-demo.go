package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := stopSignal()
	ch := generateFibonocci(stopCh)
	for no := range ch {
		fmt.Println(no)
	}
}

func stopSignal() <-chan struct{} {
	stopCh := make(chan struct{})
	fmt.Println("Hit ENTER to stop....")
	go func() {
		fmt.Scanln()
		stopCh <- struct{}{}
	}()
	return stopCh
}

func generateFibonocci(stopCh <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
		x, y := 0, 1
	LOOP:
		for {
			select {
			case <-stopCh:
				fmt.Println("Stopping the fibonocci series generation...")
				break LOOP
			case ch <- x:
				time.Sleep(500 * time.Millisecond)
				x, y = y, x+y
			}
		}
		close(ch)
	}()
	return ch
}
