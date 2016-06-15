package main

import "sort"
import "fmt"

type ByLenth []string

func (s ByLenth) Len() int {
	return len(s)
}

func (s ByLenth) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLenth) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruit := ByLenth{"peach", "banana", "kiwi"}
	sort.Sort(fruit)
	fmt.Println(fruit)
}
