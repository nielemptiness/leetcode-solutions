package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	occurred := map[int]bool{}
	for e := range nums {
		if !occurred[nums[e]] {
			occurred[nums[e]] = true
			continue
		}
		return true
	}
	return false
}

func duplicates() {
	inp := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	res := containsDuplicate(inp)

	fmt.Println(res)
}
