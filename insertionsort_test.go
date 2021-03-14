package main

import (
	"testing"
)

func insertionSort(pred func(x, y int) bool, nums ...int) []int {
	for i, v := range nums {
		j := i
		for 0 < j && pred(nums[j-1], v) {
			nums[j] = nums[j-1]
			j--
		}
		nums[j] = v
	}
	return nums
}

func TestInsertionSort(t *testing.T) {
	t.Log("Ascending", insertionSort(func(x, y int) bool { return x > y }, 5, 2, 4, 6, 1, 3))
	t.Log("Descending", insertionSort(func(x, y int) bool { return x < y }, 5, 2, 4, 6, 1, 3))
}
