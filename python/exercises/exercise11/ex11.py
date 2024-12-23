# Day 11, the advent was beautiful and quite until now
# the evil killer clown is back in town. But the evil killer clown can only wreck havoc on the town, if santa doesn't stop him.
# Santa gets into a fist-fight with the clown. Can he save Christmas?
#
#
#
# Task:
# Santa and the clown each have 3 health points.
# You will receive two Lists with symbols representing the characters' actions in a chronological order
# "-" means the character attacks
# "+" means a character regains one health points
# "0" means a character got distracted and is eating Christmas cookies

# if both characters attack at the same time, the attacks cancel each other out
# if both characters recover at the same time, both characters gain one health point each
# if one character attacks and the other is distracted or tries to recover the attack goes through

# Have a look at the Result-Enum in this Directory and return the corresponding enum entry
# If one of the characters has health less than 1, they lose
# if they have the same health at the end, nobody wins
from Result import Result

# access these variables through battle.health_santa or battle.health_clown
class battle:
    health_santa = 3
    health_clown = 3

def main(attacks_santa, attacks_clown):
    return Result.NOBODY_WINS
