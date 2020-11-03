package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	producer := producer{index: 0, listenerList: []listener{}}

	println("Running...")

	go producer.produce(&wg, 30)

	lis1 := createListener("lis-1")
	lis2 := createListener("lis-2")

	time.Sleep(time.Millisecond * 3500)
	producer.register(lis1)

	time.Sleep(time.Millisecond * 5500)
	producer.register(lis2)

	time.Sleep(time.Millisecond * 4000)
	producer.unregister(lis2)

	time.Sleep(time.Millisecond * 5500)
	producer.register(lis2)


	wg.Wait()
}
