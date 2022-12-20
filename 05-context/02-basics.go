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
	go printEvenNos(c1Ctx, wg)
	wg.Add(1)
	go printOddNos(c1Ctx, wg)
	wg.Wait()
}

func printEvenNos(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
LOOP:
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			break LOOP
		default:
			fmt.Println("Even : ", i*2)
			time.Sleep(1 * time.Second)
		}
	}
	fmt.Println("Even Done")
}

func printOddNos(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
LOOP:
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			break LOOP
		default:
			fmt.Println("Odd : ", (i*2)+1)
			time.Sleep(1 * time.Second)
		}

	}
	fmt.Println("Odd Done")
}
