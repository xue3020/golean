package main

import (
	"fmt"
	_ "runtime"
	_ "sync"
	_ "time"
)

/*
无缓冲通道上的发送操作会阻塞，直到另一个goroutine在该通道上执行接收操作，
这时值才能发送成功，两个goroutine将继续执行。相反，如果接收操作先执行，
接收方的goroutine将阻塞，直到另一个goroutine在该通道上发送一个值。
*/
func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}
func main() {
	ch := make(chan int)
	go recv(ch)
	ch <- 10
	fmt.Println("发送成功")
}
