package main

import "fmt"

import "os"

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	fmt.Printf("%v\n", p)

	//如果值是一个结构体，%+v 的格式化输出内容将包括结构体的字段名。
	fmt.Printf("%+v\n", p)
	//%#v 形式则输出这个值的 Go 语法表示。例如，值的运行源代码片段。
	fmt.Printf("%#v\n", p)
	//打印值的类型
	fmt.Printf("%T\n", p)

	fmt.Printf("%t\n", true)
	fmt.Printf("%d\n", 123)
	//输出二进制表示形式
	fmt.Printf("%b\n", 14)
	//输出给定整数的对应字符
	fmt.Printf("%c\n", 33)
	//%x 提供十六进制编码
	fmt.Printf("%x\n", 456)
	//%f 进行最基本的十进制格式化
	fmt.Printf("%f\n", 78.9)

	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	fmt.Printf("%s\n", "\"string\"")
	//%x 输出使用 base-16 编码的字符串，每个字节使用 2 个字符表示
	fmt.Printf("%x\n", "hex this")
	//输出一个指针的值
	fmt.Printf("%p\n", &p)

	fmt.Printf("|%6d|%6d|\n", 12, 345)
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%6s|%6s|\n", "foo", "b")
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")
	//Sprintf 则格式化并返回一个字符串而不带任何输出
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)
	// Fprintf 来格式化并输出到 io.Writers而不是 os.Stdout
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
