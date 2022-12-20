package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	rootCtx := context.Background()
	c1Ctx, cancel := context.WithCancel(rootCtx)
	defer cancel()
	fmt.Println("Hit ENTER to stop....")
	go func() {
		fmt.Scanln()
		cancel()
	}()
	wg.Add(1)
	go printNos(c1Ctx, wg)
	wg.Wait()
}

func printNos(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
LOOP:
	for i := 1; ; i++ {
		select {
		case <-ctx.Done():
			break LOOP
		default:
			fmt.Println(i * 10)
			time.Sleep(1 * time.Second)
		}

	}
}
