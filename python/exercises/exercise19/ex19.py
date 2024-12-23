# We talked about how classes are typically containers for data and/or behaviour.
# The first thing to implement in a class is often a form of toString().
# create an elf-object from the code below. Print it.
# then un-comment the method that is there to print out our elf's attributes
# the elves evolve like everyone does. Their special skills change, they get older and they change their weight.
# TASK: implement three methods: increment_age, change_skill, change_weight

class Elf:
    def __init__(self,name:str, height:int, weight:int,age: int,special_skill:str):
        self.name = name
        self.height = height
        self.weight = weight
        self.age = age
        self.special_skill = special_skill


    def __str__(self):
        return f"Name: {self.name} \n Height: {self.height} \n Weight: {self.weight} \n Age:{self.age} Special skill: {self.special_skill}"

    def increment_age(self):
        raise NotImplementedError


def main():
    elf=Elf("Sparkles",123,55,1234,"Dancing")


if __name__ == "__main__":
    main()