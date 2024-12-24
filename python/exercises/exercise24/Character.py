class Character:
    def __init__(self, name:str, health:int, dexterity:int, intelligence:int, charisma:int, hasShield:bool):
        self.name=name
        self.health=health
        self.dexterity=dexterity
        self.intelligence=intelligence
        self.charisma=charisma
        self.hasShield=hasShield

    def decrease_health(self, attackvalue:int):
        if self.hasShield:
            attackvalue -= 2
        self.health=self.health-attackvalue
