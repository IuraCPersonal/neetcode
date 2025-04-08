package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make(map[rune]int)

	for _, ch := range s {
		count[ch]++
	}

	for _, ch := range t {
		if count[ch] == 0 {
			return false
		}

		count[ch]--
	}

	return true
}
