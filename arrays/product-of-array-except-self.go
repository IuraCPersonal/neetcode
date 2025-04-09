package main

import "fmt"

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	fmt.Print(res)

	for i := range len(nums) {
		res[i] = 1
		for j := range len(nums) {
			if i != j {
				res[i] *= nums[j]
			}
		}

	}

	return res
}
