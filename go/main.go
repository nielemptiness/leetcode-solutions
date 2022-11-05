package main

import (
	"fmt"
)

func main() {
	input := "PAYPALISHIRING"

	res := convert(input, 3)
	fmt.Println("Input is: " + input)
	fmt.Println("Res == " + res)
	fmt.Println("Is valid: ", res == "PAHNAPLSIIGYIR")
}
