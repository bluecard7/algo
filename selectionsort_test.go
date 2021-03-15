package main

import (
	"testing"
)

func _minPos(nums []int, from int) int {
	pos := from
	for from < len(nums) {
		if nums[from] < nums[pos] {
			pos = from
		}
		from++
	}
	return pos
}

func selectionSort(nums ...int) []int {
	for i := range nums {
		pos := _minPos(nums, i)
		nums[i], nums[pos] = nums[pos], nums[i]
	}
	return nums
}

func TestSelectionSort(t *testing.T) {
	t.Log(selectionSort(5, 2, 4, 2, 3, 1))
}
