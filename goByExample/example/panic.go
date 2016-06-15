package main

import (
	"os"
)

//panic 意味着有些出乎意料的错误发生。通常我们
//用它来表示程序正常运行中不应该出现的，后者我么没有处理好的错误。
func main() {
	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
