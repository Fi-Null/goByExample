package main

import "fmt"

func main() {
	//使用 make(chan val-type) 创建一个新的通道。通道类型就是他们需要传递值的类型
	message := make(chan string)

	//使用 channel <- 语法 发送 一个新的值到通道中
	go func() {message <- "ping"}()

	//使用 <-channel 语法从通道中 接收 一个值
	msg := <- messages
	fmt.Println(msg)

	//默认发送和接收操作是阻塞的，直到发送方和接收方都准备完毕。这个特性允许我们，
	//不使用任何其它的同步操作，来在程序结尾等待消息 "ping"。
}