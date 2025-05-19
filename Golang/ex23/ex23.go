package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
You have officially achieved day 23 of Advent of Dimitri so by now you should know, that most of today's
programmers don't look at ones and zeros all day long. As we are using Computers that do binary and binary is really practical for a variety of things, you should still know
How to convert to and from binary?
Take the number
divide the number by two until you get 0 as a result
and remember the remainder
fill up the ones and zeros starting from the back.

Example:
5/2=2, remainder 1
2/2=1, remainder 0
1/2=0, remainder 1

=> 5 in binary is therefor 00000101


Today's Task:
The function convertToBinary will take an integer and convert it to 8Bit-Binary. Return it as a string for simplicity's sake
The function convertToDecimal does the exact opposite

*/

func convertToBinary(decimalNumber int) string {
	output := ""
	for decimalNumber > 0 {
		remainder := decimalNumber % 2
		decimalNumber = decimalNumber / 2
		output = strconv.Itoa(remainder) + output
	}

	return output
}

func convertToDecimal(binaryNumber string) int {

	sum := 0
	counter := 0
	for len(binaryNumber) > 0 {
		current := binaryNumber[len(binaryNumber)-1:]
		binaryNumber = binaryNumber[:len(binaryNumber)-1]

		currentAsNum, _ := strconv.Atoi(current)

		if currentAsNum > 0 {
			sum += int(math.Pow(2, float64(counter)))
		}

		counter++
	}

	return sum
}
func main() {
	/*
		fmt.Println(convertToBinary(16))
		fmt.Println(convertToDecimal("10001"))
		fmt.Println(convertToBinary(0))
		fmt.Println(convertToDecimal("00000000"))
		fmt.Println(convertToBinary(255))
	*/

	fmt.Println(convertToDecimal("11111111"))
	/*
		fmt.Println(convertToBinary(122))
		fmt.Println(convertToDecimal("01111010"))*/
}
