# Okay Santa thinks that coals are a tiny bit outdated for the bad kids, so instead he gifts them a recursion-quiz for punishment.
# You may wonder: what is recursion?
# Recursion in the most simple sense happens, when a function or method calls itself.
# This means: a function must give itself parameters to continue to solve the puzzle. Under a certain condition, the puzzle is solved and the function can stopp calling itself
# If this condition doesn't occur, this will go to infinity and yes that is bad

# To make sure the bad kids can solve the recursion quiz you need to test it by adding to the following code:
# Your mission is quite simple: you receive a list of numbers, find the first prime number and return it.
# The numbers will not go above 100, however it is your job to make sure santa doesn't accidentally enter negatives into his unit-tests (wink wink)
# If santa makes such a bad mistake simply return -1
# You can find a list of primes on the internet.


nums=[4,122,6,8,2,13,2]

def findPrimeNumber(nums) ->int:
    #TODO: Define an exit-condition
    if True:
        raise NotImplementedError
    else:
        return findPrimeNumber(nums[1:])

# this is just for your convenience.
def main():
    print(findPrimeNumber(nums))

if __name__ == "__main__":
    main()