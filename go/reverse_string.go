package main

import (
	"fmt"
)

func reverseString(s []string) {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i += 1
		j -= 1
	}
}

func main() {
	input := []string{"h", "e", "l", "l", "o"}
	reverseString(input)
	fmt.Println(input)
}
