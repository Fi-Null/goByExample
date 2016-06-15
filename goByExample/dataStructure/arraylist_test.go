package dataStructure

import (
	"fmt"
	"testing"
)

func TestArrayList(t *testing.T) {
	fmt.Println("fef")

	list := New()
	list.Append("hello")
	list.Append("world")

	fmt.Println(list)

}
