package main

import (
	"fmt"
	"sort"
)

/*
Santa is ready to go deliver some presents, but his reindeer are scattered around the north-pole.
Some are in the factory, looking for snacks
Some are outside, glancing at the stars
And a few are even by the fireplace enjoying the cozy warmth.
Collect them all and sort them alphabetically, so Santa can bridle them for their work
Throw an error, if you can't find them all (there are 9)
*/

func main() {
	factory := []string{"Dasher", "Rudolph", "Vixen"}
	outdoors := []string{"Dancer", "Comet", "Prancer"}
	fireplace := []string{"Cupid", "Dunder", "Blixem"}

	factory = append(factory, outdoors...)

	factory = append(factory, fireplace...)
	if len(factory) < 9 || len(factory) > 9 {
		panic("unmatching number of reindeer")
	} else {
		sort.Strings(factory)
		fmt.Println("factory:", factory)
	}

}
