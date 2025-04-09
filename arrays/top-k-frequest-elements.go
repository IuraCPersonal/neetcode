package main

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	freq := make([][]int, len(nums)+1)

	for _, n := range nums {
		count[n]++
	}

	for k, v := range count {
		freq[v] = append(freq[v], k)
	}

	res := make([]int, 0, k)

	for i := len(freq) - 1; i > 0; i-- {
		for _, n := range freq[i] {
			res = append(res, n)

			if len(res) == k {
				return res
			}
		}
	}

	return res
}
