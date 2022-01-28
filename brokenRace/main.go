package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var result int

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go runProcess("P1", 20, wg)
	go runProcess("P2", 20, wg)

	wg.Wait()
	fmt.Println("Final result: ", result)
}

func runProcess(name string, total int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < total; i++ {
		z := result
		z++
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
		result = z
		fmt.Println(name, "->", i, "Partial result:", result)
	}
}
