package main

import (
	"fmt"
)

func topKFrequent(nums []int, k int) []int {
	if len(nums) < 0 || k < 1 || k > len(nums) {
		return nil
	}

	result := []int{}

	m := make(map[int]int)

	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v] = m[v] + 1
		} else {
			m[v] = 1
		}
	}

	for i := 0; i < k; i++ {
		maxCount := 1
		val := 0
		for k, v := range m {
			if v >= maxCount {
				maxCount = v
				val = k
			}
		}
		delete(m, val)
		result = append(result, val)
	}

	return result

}

func main() {
	nums := []int{1, 2, 2, 4, 4, 5}
	fmt.Println(topKFrequent(nums, 2))
	go topKFrequent(nums, 3)
}
