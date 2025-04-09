package main

type Key [26]byte

func strKey(str string) (key Key) {
	for i := range str {
		key[str[i]-'a']++
	}

	return key
}

func groupAnagrams(strs []string) [][]string {
	groups := make(map[Key][]string)

	for _, v := range strs {
		key := strKey(v)
		groups[key] = append(groups[key], v)
	}

	res := make([][]string, 0, len(groups))

	for _, v := range groups {
		res = append(res, v)
	}

	return res
}
