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
	valCtx := context.WithValue(rootCtx, "k1", "v1")
	c1Ctx, cancel := context.WithCancel(valCtx)

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
	fmt.Println("[printNos] k1 = ", ctx.Value("k1"))
	defer wg.Done()

	newValCtx := context.WithValue(ctx, "k1", "new-v1")

	wg.Add(1)
	timeoutCtx1, cancel1 := context.WithTimeout(newValCtx, 5*time.Second)
	defer cancel1()
	go printEvenNos(timeoutCtx1, wg)

	timeoutCtx2, cancel2 := context.WithTimeout(newValCtx, 10*time.Second)
	defer cancel2()
	wg.Add(1)
	go printOddNos(timeoutCtx2, wg)

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
	fmt.Println("[printEvenNos] k1 = ", ctx.Value("k1"))
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
	fmt.Println("[printOddNos] k1 = ", ctx.Value("k1"))
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
