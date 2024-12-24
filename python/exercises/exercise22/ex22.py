# Santa has a special lock for his sleigh, so the elves don't take it out for a ride whenever they feel like it
# Santa always forgets the passcode for his lock so he just wants it to be "Merry X-MAS"
# This is a bad idea, since the elves can easily guess or brute-force this problem
# Task: Santa will suggest some Passcodes and you will need to validate them
# The password must be at least 8 characters long.
# It must consist of uppercase and lowercase letters
# The password must consist of at least one Number and one special character (!,?,$,£,#,@)
# There cannot be more than three of each type in a row
# If the password does not fit these criteria, raise a BadPasswordError
class BadPasswordError(Exception):
    pass


def checkPassword(suggestion:str) -> bool:
    alphabet_lowercase=list("abcdefghijklmnopqrstuvwxyz")
    alphabet_uppercase=list("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
    special_characters=list("!?$£#@")
    numbers=list("0123456789")
    raise NotImplementedError


