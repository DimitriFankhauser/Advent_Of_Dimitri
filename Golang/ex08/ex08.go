package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
Like every modern company Santa Clause LLC also has their employees (elves) in a computer system.
However the elves got confused about their lists.
They have a list of elves stored in their system, which likely contains duplicates.
Read the list from the file elves.txt and make sure they haven't stored anyone twice
Remove the duplicates and write them into elves_unique.txt
As in ex07 there will be a unit-test, that you can run, to recreate the original list, in case you make a mistake or would like to redo the challenge

Hint: Python has a really cool data-structure that you can use for this
*/
func main() {
	elves, _ := os.Open("./ex08/elves.txt")
	elvesUnique, err := os.OpenFile("./ex08/elves_unique.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0660)
	if err != nil {
		log.Fatal(err)
	}
	defer elves.Close()
	defer elvesUnique.Close()

	scanner := bufio.NewScanner(elves)

	// Go doesn't have a set-implementation like Python's. For this reason We are creating a map that looks like this:
	// <Elfname>:<existing>
	set := make(map[string]bool)
	for k := range set {
		fmt.Println(k)
	}
	var i = 0
	for scanner.Scan() {
		current := scanner.Text()
		i++
		if !set[current] {
			set[current] = true
			elvesUnique.WriteString(current + "\n")
		}
	}
}
