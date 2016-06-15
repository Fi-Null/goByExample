package main

import (
	"fmt"
	"os"
)

func main() {

	//当使用 os.Exit 时 defer 将不会 执行，
	//所以这里的 fmt.Println将永远不会被调用。
	defer fmt.Println("!")

	os.Exit(3)
}
