package main

import "fmt"

/*We talked about how classes are typically containers for data and/or behaviour.
The first thing to implement in a class is often a form of toString().
create an elf-object from the code below. Print it.
then un-comment the method that is there to print out our elf's attributes
the elves evolve like everyone does. Their special skills change, they get older and they change their weight.
TASK: implement three methods: increment_age, change_skill, change_weight
*/

type Elf struct {
	name          string
	height        int
	weight        int
	age           int
	special_skill string
}

func (re *Elf) incrementAge() {
	re.age++
}

func (re *Elf) changeSkill(skill string) {
	re.special_skill = skill
}

func (re *Elf) changeWeight(weight int) {
	re.weight = weight
}
func main() {
	var a = Elf{"Bernard", 115, 55, 1002, "Leadership"}
	fmt.Println(a)
	a.incrementAge()
	fmt.Println(a)
	a.changeSkill("Singing")
	fmt.Println(a)
	a.changeWeight(57)
	fmt.Println(a)
}
