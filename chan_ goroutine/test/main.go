package main

import (
	"fmt"
	"sync"
	_ "time"
)

var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 1000; i++ {
		lock.Lock()
		x = x + 1
		// fmt.Println(x)
		lock.Unlock()
	}
	wg.Done()
}
func main() {
	wg.Add(3)
	go add()
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
