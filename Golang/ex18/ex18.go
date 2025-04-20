package main

import (
	"fmt"
	"math/rand/v2"
)

/*
Dictionaries can be really helpful to map values to their respective definition.
As Christmas is coming up Bernard, the Head Elf is already planning the Christmas dinner. The seats in the big hall are numbered.
Certain elves want to sit next to each other, while others don"t want to sit next to each other.
Bernard can"t cater to everyone"s wishes, so he assigns the seats randomly.
Task: you receive a list of elves and randomly assign a seat to each elf. Return a dictionary in the form of <elf-name>:<seat>
Throw an exception if the list you receive is empty*/

func main() {
	var elves = [...]string{"Jingle", "Tinsel", "Peppermint", "Snowball", "Holly", "Sugarplum", "Frosty", "Merry", "Twinkle", "Candycane",
		"Gingersnap", "Sparkle", "Sprinkle", "Jolly", "Chestnut", "Cinnamon", "Pinecone", "Kringle",
		"Noel", "Bells"}

	numbers := []int{}

	// Iterate from 1 to 20 (inclusive) and append each number to the slice.
	for i := 1; i <= 20; i++ {
		numbers = append(numbers, i)
	}

	var m map[int]string
	m = make(map[int]string)

	for i := 0; i < len(elves); i++ {
		//generate a random index within the range of "numbers"
		index := rand.IntN(len(numbers))
		num := numbers[index]
		if m[num] == "" {
			fmt.Println("success")
			m[num] = elves[i]
			//recreate numbers without the one we picked out
			numbers = append(numbers[:index], numbers[index+1:]...)
		} else {
			fmt.Println("fail")
		}

	}

	for key, value := range m { // Order not specified
		fmt.Println(key, value)
	}
	fmt.Println(len(m))
}
