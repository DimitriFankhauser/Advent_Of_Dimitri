package main

import (
	"fmt"
	"slices"
)

/*
Okay Santa thinks that coals are a tiny bit outdated for the bad kids, so instead he gifts them a recursion-quiz for punishment.
You may wonder: what is recursion?
Recursion in the most simple sense happens, when a function or method calls itself.
This means: a function must give itself parameters to continue to solve the puzzle. Under a certain condition, the puzzle is solved and the function can stopp calling itself
If this condition doesn't occur, this will go to infinity and yes that is bad

To make sure the bad kids can solve the recursion quiz you need to test it by adding to the following code:
Your mission is quite simple: you receive a list of numbers, find the first prime number and return it.
The numbers will not go above 100, however it is your job to make sure santa doesn't accidentally enter negatives into his unit-tests (wink wink)
If santa makes such a bad mistake simply return -1
You can find a list of primes on the internet.
*/

func findPrimeNumber(nums []int16) int16 {
	primes := []int16{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	for i := 0; i < len(nums); i++ {
		current := nums[i]
		if current < 0 || current > 100 {
			return -1
		}
		if slices.Contains(primes, current) {
			return current
		}
	}
	return -1
}

func main() {
	nums := []int16{2, 24, 55, 26, 75, 100}
	fmt.Println("The prime is", findPrimeNumber(nums))

	nums = []int16{24, 55, 10, 23}
	fmt.Println("The prime is", findPrimeNumber(nums))

	nums = []int16{24, 55, 7, 10, 23}
	fmt.Println("The prime is", findPrimeNumber(nums))

	nums = []int16{24, 55, 7, 10, 23}
	fmt.Println("The prime is", findPrimeNumber(nums))

	nums = []int16{24, 55, 10, -134, 4, 8}
	fmt.Println("The prime is", findPrimeNumber(nums))

	nums = []int16{24, 55, 10, 134, 4, 8}
	fmt.Println("The prime is", findPrimeNumber(nums))

}
