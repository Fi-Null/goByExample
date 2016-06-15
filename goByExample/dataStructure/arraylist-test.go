package main

import (
	"fmt"
	"goByExample/dataStructure/arraylist"
	"goByExample/dataStructure/linkedlist"
)

func main() {
	fmt.Println("fef")

	list := arraylist.New()
	list.Append("hello")
	list.Append("world")

	fmt.Println(list)

}
