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
	wg.Add(1)
	go printEvenNos(ctx, wg)
	wg.Add(1)
	go printOddNos(ctx, wg)
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
	fmt.Println("Nos Done")
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
			time.Sleep(500 * time.Millisecond)
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
			time.Sleep(300 * time.Millisecond)
		}

	}
	fmt.Println("Odd Done")
}
