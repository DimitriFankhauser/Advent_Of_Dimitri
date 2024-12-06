/*
Loops & Comparisons:
Before we get to joyful christmas, it's halloween. There is a killer clown who has a list of all people in Clownsville with him.
The list only contains first names. The clown gives you his list as a list of strings like so ["Hans", "Malcom", "Charlie"]
Today the killer clown only kills Hans to still his bloodlust. Remove Hans from the list and return the rest
In this challenge Hans will only be present in the list ONCE */

package main

import "fmt"

func removeHans(victims []string) []string {
	remainingVictims := make([]string, 0)
	for i := 0; i < len(victims); i++ {
		if victims[i] != "Hans" {
			remainingVictims = append(remainingVictims, victims[i])
		}
	}
	return remainingVictims
}

// the param victims is a pointer, that points to a slice of strings (the var pointervictims)
func removeHansWithPointers(victims *[]string) {
	// no return-type because we are using reference params!
	// originalvictims has the value of the slice pointervictims
	originalVictims := *victims
	//create a slice with len0, since it's empty
	remainingVictims := make([]string, 0)

	for i := 0; i < len(originalVictims); i++ {
		if originalVictims[i] != "Hans" {
			//reassign remainingVictims with the current len, which is equal to i
			remainingVictims = append(remainingVictims, originalVictims[i])
		}
	}

	// Update the slice that the pointer is pointing to
	*victims = remainingVictims
}

func main() {
	victims := []string{"Julia", "Lenny", "Hans"}
	remainingVictims := removeHans(victims)

	pointerVictims := []string{"Julia", "Lenny", "Hans"}
	victimspointer := &pointerVictims
	removeHansWithPointers(victimspointer)
	fmt.Println(pointerVictims)
	fmt.Println(remainingVictims)
}
