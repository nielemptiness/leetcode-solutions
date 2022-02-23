package main

import (
	"fmt"
)

func rotate(nums []int, k int) []int {

	l := len(nums)
	k = k % l
	for i := 0; i < k; i++ {
		last := nums[l-1]
		for j := l - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = last
	}
	return nums
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	count := 3
	res := rotate(arr, count)
	fmt.Println(res)

	arr2 := []int{-1}
	count2 := 2
	res2 := rotate(arr2, count2)
	fmt.Println(res2)

	arr3 := []int{1, 2}
	count3 := 3
	res3 := rotate(arr3, count3)
	fmt.Println(res3)

	arr4 := []int{1, 2, 3}
	count4 := 4
	res4 := rotate(arr4, count4)
	fmt.Println(res4)
}
