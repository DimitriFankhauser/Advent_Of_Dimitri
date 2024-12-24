from random import randint


from python.exercises.exercise24.Character import Character


class Monster:
    def __init__(self, name, skill):
        self.name = name
        self.skill = skill

    def attack(self, character: Character):
        attackvalue=randint(1,15)
        character.decrease_health(attackvalue)