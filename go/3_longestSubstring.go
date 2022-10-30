package main

// use this cuz math.max uses float :(
func getBiggest(first int, second int) int {
	if first > second {
		return first
	} else {
		return second
	}
}

// window moving
func lengthOfLongestSubstring(s string) int {
	biggest := 0
	start := 0
	dict := make(map[byte]int)

	for end := 0; end < len(s); end++ {

		if _, ok := dict[s[end]]; ok {
			start = getBiggest(start, dict[s[end]]+1)
		}

		dict[s[end]] = end
		biggest = getBiggest(biggest, end-start+1)
	}
	return biggest
}
