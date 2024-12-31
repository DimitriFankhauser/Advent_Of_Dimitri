package main

import "fmt"

/*
now that we've established some basic class-knowledge, let's take it to another level!
Your task now is the following. Some elves are lactose-intolerant.
To accommodate them we'd like to know who can and who can't have dairy in their Christmas dinner.
Your task now is to copy the Elf-Class into this folder. Add a field "lactose_intolerant" to it and recreate the elf-object from ex13.
Add the information, that Sparkles is lactose intolerant to it
*/
type elf struct {
	id                uint32
	name              string
	balance           float32
	lactoseIntolerant bool
}

func main() {
	e := elf{}
	e.id = 24
	e.name = "Sparkles"
	e.balance = 24.12
	e.lactoseIntolerant = true

	f := elf{24, "Sparkles", 24.12, true}
	fmt.Println(e)
	fmt.Println(f)
}
