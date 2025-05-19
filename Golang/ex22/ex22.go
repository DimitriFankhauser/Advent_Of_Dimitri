package main

import (
	"fmt"
	"strings"
)

/*
Santa has a special lock for his sleigh, so the elves don't take it out for a ride whenever they feel like it
Santa always forgets the passcode for his lock so he just wants it to be "Merry X-MAS"
This is a bad idea, since the elves can easily guess or brute-force this problem
Task: Santa will suggest some Passcodes and you will need to validate them
The password must be at least 8 characters long.
It must consist of uppercase and lowercase letters
The password must consist of at least one Number and one special character (!,?,$,£,#,@)
There cannot be more than three of each type in a row
If the password does not fit these criteria, raise a BadPasswordError
*/

func checkPassword(proposedPassword string) bool {
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercase := "abcdefghijklmnopqrstuvwxyz"
	specialCharacters := "!?$£#@"

	lowercaseFlag := false
	uppercaseFlag := false
	specialCharFlag := false
	lengthFlag := false

	if len(proposedPassword) > 7 {
		lengthFlag = true
	}

	for _, char := range proposedPassword {
		if strings.Contains(uppercase, string(char)) {
			uppercaseFlag = true
		}
		if strings.Contains(lowercase, string(char)) {
			lowercaseFlag = true
		}
		if strings.Contains(specialCharacters, string(char)) {
			specialCharFlag = true
		}
	}

	return lowercaseFlag && uppercaseFlag && specialCharFlag && lengthFlag

}
func main() {
	fmt.Println(checkPassword("asdFj8!"))
	fmt.Println(checkPassword("$$$$$$$$"))
	fmt.Println(checkPassword("AAAAAAAA"))
	fmt.Println(checkPassword("aaaaaaaa"))
	fmt.Println(checkPassword("00000000"))
	fmt.Println(checkPassword("asddddFj8!"))
	fmt.Println(checkPassword("asdFFFFj8!"))
	fmt.Println(checkPassword("asdFj8888!"))
	fmt.Println(checkPassword("asdFj8!!!!"))
	fmt.Println(checkPassword("asd!8"))
	fmt.Println(checkPassword(" "))
}
