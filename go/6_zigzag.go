package main

import "fmt"

func getColumnStep(index int, step int, row int) int {
	return index + step - 2*row
}
func convert(s string, numRows int) string {
	if numRows <= 0 {
		return ""
	}
	if numRows == 1 {
		return s
	}
	var result string
	step := 2*numRows - 2

	for i := 0; i < numRows; i++ {
		for j := i; j < len(s); j += step {
			fmt.Println(string(s[j]))
			result += string(s[j])
			index := getColumnStep(j, step, i)
			if i != 0 && i != numRows-1 && index < len(s) {
				fmt.Println(string(s[index]))
				result += string(s[index])
			}
		}
	}
	return result
}

func testZigzag() {
	input := "PAYPALISHIRING"

	res := convert(input, 3)
	fmt.Println("Input is: " + input)
	fmt.Println("Res == " + res)
	fmt.Println("Is valid: ", res == "PAHNAPLSIIGYIR")
}
