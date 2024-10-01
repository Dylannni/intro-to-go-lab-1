package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Println("Hello from goroutine", i)
}

func main() {
	i := 1
	for i < 6 {
		go fmt.Println("Hello World", i)
		i += 1
	}

	//for i := 0; i < 5; i++ {
	//	go hello(i)
	//}
	time.Sleep(1 * time.Second)
}
