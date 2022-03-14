package reverseString

import (
	"fmt"
)

func ReverseString(s []string) {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i += 1
		j -= 1
	}
}

func Reverse() {
	input := []string{"h", "e", "l", "l", "o"}
	ReverseString(input)
	fmt.Println(input)
}
