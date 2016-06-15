package main

import {
	"fmt"
	"time"	
}

//Go 的通道选择器 让你可以同时等待多个通道操作
func main() {
	c1 := make(chan string)
	c2 := mkae(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 1)
		c2 <- "two"
	}()


	//我们使用 select 关键字来同时等待这两个值，并打印各自接收到的值
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

}