package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutines") //使用 go f(s) 在一个 Go 协程中调用这个函数。
	//这个新的 Go 协程将会并行的执行这个函数调用。

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var intput string
	fmt.Scanln(&input)
	fmt.Println("done")
}
