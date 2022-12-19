package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main started")
	go f1() // scheduling f1
	f2()
	time.Sleep(1 * time.Second)
	// fmt.Scanln()
	fmt.Println("main completed")
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
