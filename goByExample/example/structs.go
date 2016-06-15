package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name : "Alice", age : 20})

	fmt.Println(person{name : "Fred"})//省略的字段将被初始化为零值

	fmt.Println(&person{name : "Ann", age : 40})

	s := person{name : "Sean", age : 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)//也可以对结构体指针使用. - 指针会被自动解引用

	sp.age = 51
	fmt.Println(sp.age)		

}
