package main

import "fmt"

func main() {
	// make 了一个通道，最多允许缓存 2 个值
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
