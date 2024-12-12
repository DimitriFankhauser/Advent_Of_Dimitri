# In the previous exercise you were probably wondering:
# What do I do, if the attacks of Santa or the clown make no sense?
# You throw an error
# in this tiny little program your task is the following:
# - take the input of f
# - divide it by three
# - print the output
# - if the input is not a number, raise an exception
# - OR divide by three and catch the exception, that will be thrown
# - print the result to the console

def main():
    running = True
    while running:
        f = input("Enter a number: ")
        # Your Code here
        f /= 3
        print(str(int(f)))


if __name__ == "__main__":
    main()
