/*
	Create a goroutine "GeneratePrimes" that will generate the prime numbers between the given start and end
	The main function should receive the generated prime numbers and print them one after another
*/

package main

import (
	"fmt"
	"time"
)

//consumer
func main() {
	ch := GeneratePrimes(3, 100)
	for no := range ch {
		fmt.Println(no)
	}
}

//producer
func GeneratePrimes(start, end int) <-chan int {
	ch := make(chan int)
	go func() {
		for no := start; no <= end; no++ {
			if isPrime(no) {
				time.Sleep(500 * time.Millisecond)
				ch <- no
			}
		}
		close(ch)
	}()
	return ch
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
