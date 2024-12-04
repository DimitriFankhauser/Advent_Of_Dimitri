# Santa has a naughty and nice list of all the kids
# However Santa does not want the kids to know until Christmas, if they are on the nice or naughty list.
# Santa encrypts and so should you
# Santa uses a cesar cypher. cesar cyphers work by substituting every letter in a word by another letter that is a fixed number of places in the alphabet away
# let's say this number is 2.
# [A,B,C,D,E,F]
# ABC becomes CDE.
# This challenge does not include a unit-test. Implement a basic cesar cypher and get a jolly message.
# You only shift in the normal alphabet (abcdefghijklmnopqrstuvwxyz), you don't have to account for special characters. You also have to find out shift-number in order to get the jolly message.
# HINT: this can also work with negative shifts, meaning you have to turn CDE -> ABC

def main(word:str):
    print(word)



if __name__=='__main__':
    main("kviex nsf, csyzi fiir zivc rmgi xlmw ciev!")