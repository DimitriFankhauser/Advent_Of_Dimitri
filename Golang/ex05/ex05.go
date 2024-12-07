package main

import (
	"fmt"
	"strings"
)

/*
Santa has a naughty and nice list of all the kids
However Santa does not want the kids to know until Christmas, if they are on the nice or naughty list.
Santa encrypts and so should you
Santa uses a cesar cypher. cesar cyphers work by substituting every letter in a word by another letter that is a fixed number of places in the alphabet away
let's say this number is 2.
[A,B,C,D,E,F]
ABC becomes CDE.
This challenge does not include a unit-test. Implement a basic cesar cypher and get a jolly message.
You only shift in the normal alphabet (abcdefghijklmnopqrstuvwxyz), you don't have to account for special characters. You also have to find out shift-number in order to get the jolly message.
HINT: this can also work with negative shifts, meaning you have to turn CDE -> ABC*/

const alphabet string = "abcdefghijklmnopqrstuvwxyz"
const specialCharacters string = ", !"

func cesarCypher(message string) string {
	result := ""
	for i := 0; i < len(message); i++ {
		letter := message[i : i+1] //that way we get a string of len 1 instead of a rune (message[i])
		if strings.Index(specialCharacters, letter) < 0 {
			indexInAlphabet := strings.Index(alphabet, letter)
			indexInAlphabet -= 4 //shift by four letters
			if indexInAlphabet < 0 {
				indexInAlphabet = 26 + indexInAlphabet //letters like c will lead to a negative index due to the shift.
			}
			result += alphabet[indexInAlphabet : indexInAlphabet+1]

		} else {
			result += letter
		}
	}
	return result
}

func main() {
	message := cesarCypher("kviex nsf, csyzi fiir zivc rmgi xlmw ciev!")
	fmt.Println(message)
}
