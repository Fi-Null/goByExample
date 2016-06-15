package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Println("%s -> %s", k, v)
	}

	for i, c := range "go" {
		fmt.Println(i, c) //range 在字符串中迭代 unicode 编码。
		//第一个返回值是rune 的起始字节位置，然后第二个是 rune 自己
	}

}
