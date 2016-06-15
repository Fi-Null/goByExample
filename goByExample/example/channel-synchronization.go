package main

import "fmt"
import "time"

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	//程序将在接收到通道中 worker 发出的通知前一直阻塞。
	go worker(done)

	//如果你把 <- done 这行代码从程序中移除
	//，程序甚至会在 worker还没开始运行时就结束了
	<-done
}
