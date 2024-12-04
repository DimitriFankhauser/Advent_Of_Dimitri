# It's becoming christmas now.
# You now receive a christmas tree in the form of a string
# branches can point to the left "/", be straight "|" or point to the right "\"
# if a branch is straight, add two christmas ornaments to it, so it looks like this:  *|* and return the entire string
# example input: //\\||
# turn it into this: //*|*\*|**|*
# nothing special happens if two straight branches are next to each other
# if there is no straight branch return "No straight branch in this tree"
# if you don't know the concept of return-types: now is the time to google it and figure out what -> str means
def main(tree: str)->str:
    return "No straight branch in this tree"
