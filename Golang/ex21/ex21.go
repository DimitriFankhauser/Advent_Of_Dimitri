package main

import "fmt"

/*
Do you like reindeer? I love them, so here"s another reindeer-centric exercise:
Now that we"ve gathered all reindeer to go on this global, package-delivery mission, there is one thing left to do:
Make sure all reindeer have enough bells mounted on them. They need a minimum of 3 and a maximum of 7 bells
Return a tuple with the reindeer, that need to be corrected
*/

type Reindeer struct {
	name  string
	bells int
}

func main() {
	reindeerWithoutBells := []string{"Blixem", "Comet", "Cupid", "Dancer", "Dasher", "Dunder", "Prancer", "Rudolph", "Vixen"}
	ratings := []int{1, 17, 3, 4, 238, 7, 13, -7, 2}

	reindeerTotal := []Reindeer{}
	for x, reindeer := range reindeerWithoutBells {
		reindeerTotal = append(reindeerTotal, Reindeer{reindeer, ratings[x]})
	}
	fmt.Println("all reindeer:", reindeerTotal)

	disobedientReindeer := []Reindeer{}
	for _, reindeer := range reindeerTotal {
		if reindeer.bells < 3 || reindeer.bells > 7 {
			disobedientReindeer = append(disobedientReindeer, Reindeer{reindeer.name, reindeer.bells})
		}
	}

	fmt.Println("disobedientReindeer:", disobedientReindeer)

}
