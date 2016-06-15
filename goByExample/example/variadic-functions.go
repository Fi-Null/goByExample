package main

import "fmt"

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int {1, 2, 3, 4}
	sum(nums...)//如果你的 slice 已经有了多个值，
	//想把它们作为变参使用，你要这样调用 func(slice...)
}
