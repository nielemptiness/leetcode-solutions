package main

import (
	"fmt"
	ua "radioemptinessgoleetcode/IsOrderCorrect"
)

func main() {
	var input = "()[]{}"
	res := ua.IsValid(input)

	fmt.Println(res)

	var inp2 = "(]"
	res2 := ua.IsValid(inp2)

	fmt.Println(res2)
}
