package main

import "fmt"

/*
A Class is a "blueprint" for a data-object. Meaning:
In a class you may define how your data is supposed to look like.
Classes have different reasons for their existence: mostly to achieve a certain behaviour or to store data in a way that makes sense
have a look at the elf-class in Elf.py
Your task is to create a single elf-object from the class with the following properties:
id=24
name= Sparkles
balance= 24.12*/

type elf struct {
	id      uint32
	name    string
	balance float32
}

func main() {
	e := elf{}
	e.id = 24
	e.name = "Sparkles"
	e.balance = 24.12

	f := elf{24, "Sparkles", 24.12}
	fmt.Println(e)
	fmt.Println(f)

}
