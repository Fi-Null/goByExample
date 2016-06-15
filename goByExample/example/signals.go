package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

//Go 通过向一个通道发送 os.Signal
//值来进行信号通知。我们将创建一个通道来接收这些通知
func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
