/*
Loops & Comparisons:
It's halloween. There is a killer clown who has a list of all people in Clownsville with him.
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
func main() {
	victims := []string{"Julia", "Lenny", "Hans"}
	remainingVictims := removeHans(victims)
	fmt.Println(remainingVictims)
}
