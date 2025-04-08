package main

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)

	for i, v := range nums {
		diff := target - v

		if index, ok := hash[diff]; ok {
			return []int{index, i}
		}

		hash[v] = i
	}

	return []int{}
}

// func twoSum(nums []int, target int) []int {
// 	var res []int
// 	n := len(nums)
//
// 	for i := range n - 1 {
// 		for j := i + 1; j < n; j++ {
// 			if nums[i]+nums[j] == target {
// 				res = append(res, i)
// 				res = append(res, j)
// 			}
// 		}
//
// 	}
//
// 	return res
// }
