package main

import (
	"fmt"
	"sort"
)

func divMod(num, divider int) (whole, remainder int) {
	fmt.Println("divider:", divider)
	whole, remainder = num/divider, num%divider
	fmt.Println("whole: ", whole)
	fmt.Println("remainder:", remainder)
	return
}

func getSortedMapKeys(mapping map[int]string) []int {
	keys := make([]int, len(mapping))

	i := 0
	for key := range mapping {
		keys[i] = key
		i++
	}
	sort.SliceStable(keys, func(a, b int) bool {
		return keys[a] > keys[b]
	})
	return keys
}
func intToRoman(num int) string {
	mapping := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	var result string

	keys := getSortedMapKeys(mapping)

	for i := 0; i < len(keys); i++ {
		// lots of corner cases
		// var wholePart int
		// wholePart, num = divMod(num, keys[i])
		// fmt.Println(mapping[keys[i]])
		// if wholePart == 1 {
		// 	result += mapping[keys[i]]
		// }

		// correct solution, but slow & not memory efficient cuz of map + sorting
		if num == 0 {
			continue
		}
		for num >= keys[i] {
			num -= keys[i]
			result += mapping[keys[i]]
		}

	}
	return result
}

var values = []struct {
	value  int
	symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman2(num int) string {
	var out string
	for _, c := range values {
		fmt.Println(c)
		for num >= c.value {
			num -= c.value
			out += c.symbol
		}
		if num <= 0 {
			break
		}
	}
	return out
}

func testIntToRoman() {
	// res := intToRoman(1994)
	// fmt.Println(res)
	// fmt.Println(res == "MCMXCIV")
	res := intToRoman(20)
	fmt.Println(intToRoman2(20))
	fmt.Println(res)
	fmt.Println(res == "XX")
}
