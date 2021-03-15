package main

import (
	"testing"
)

func merge(left, right []int) []int {
	buf := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			buf = append(buf, left[i])
			i++
		} else {
			buf = append(buf, right[j])
			j++
		}
	}
	if i < len(left) {
		buf = append(buf, left[i:]...)
	}
	if j < len(right) {
		buf = append(buf, right[j:]...)
	}
	return buf
}

func mergeSort(nums ...int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := int(len(nums) / 2)
	left := mergeSort(nums[0:mid]...)
	right := mergeSort(nums[mid:]...)
	return merge(left, right)
}

func TestMergeSort(t *testing.T) {
	t.Log(mergeSort(5, 2, 3, 4, 6))
}
