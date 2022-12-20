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
	c1Ctx, cancel := context.WithTimeout(rootCtx, time.Duration(10)*time.Second)
	defer cancel()
	valCtx := context.WithValue(c1Ctx, "k1", "v1")
	fmt.Println("Hit ENTER top stop or wait for 10s....")
	go func() {
		fmt.Scanln()
		cancel()
	}()
	wg.Add(1)
	go printNos(valCtx, wg)
	wg.Wait()
}

func printNos(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[printNos] k1 = ", ctx.Value("k1"))
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
	fmt.Println("Nos done!")
}
