package main

import (
	_"runtime"
	"sync"
	"fmt"
	"time"
)

var wg sync.WaitGroup

func fn1(ch chan int){
	for i := 1; i<=10; i ++{
		ch <- i
		fmt.Printf("写入数据%v成功\n",i)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Done()
}

func fn2(ch chan int){
	for i :=1; i<=10; i++{
		ai :=<-ch
		fmt.Printf("读取数据%v成功\n",ai)
		time.Sleep(time.Millisecond)
	}
	wg.Done()
}


func main(){

	var ch = make(chan int,10)

	wg.Add(1)
	go fn1(ch)
	wg.Add(1)
	go fn2(ch)
	wg.Wait()

}