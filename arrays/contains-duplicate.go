package main

func containsDuplicate(nums []int) bool {
	set := make(map[int]struct{})

	for _, num := range nums {
		if _, err := set[num]; err {
			return true
		}

		set[num] = struct{}{}
	}

	return false
}
