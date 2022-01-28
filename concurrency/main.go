package main

import (
	"fmt"
	"sync"
)

func worker(channel chan int, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for value := range channel {
		fmt.Println("WorkerId ", id, " is processing message ", value)
	}

}

func main() {
	fmt.Println("Proccess has started")
	messageChannel := make(chan int)
	wg := new(sync.WaitGroup)
	workersAmount := 10
	messagesAmount := 100

	for workerId := 0; workerId < workersAmount; workerId++ {
		wg.Add(1)
		go worker(messageChannel, workerId, wg)
	}

	for message := 1; message <= messagesAmount; message++ {
		messageChannel <- message
	}

	close(messageChannel)

	wg.Wait()

	fmt.Println("Proccess has ended")
}
