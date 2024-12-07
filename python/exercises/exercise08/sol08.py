import unittest


class sol08(unittest.TestCase):
    elves_duplicates = ["Jingle", "Tinsel", "Peppermint", "Snowball", "Holly", "Sugarplum", "Frosty", "Merry",
                        "Twinkle", "Candycane",
                        "Gingersnap", "Sparkle", "Sprinkle", "Jolly", "Tinsel", "Chestnut", "Chestnut", "Cinnamon",
                        "Pinecone", "Kringle", "Noel", "Bells", "Sprinkle", "Sprinkle"]

    elves_singleton = ["Jingle", "Tinsel", "Peppermint", "Snowball", "Holly", "Sugarplum", "Frosty", "Merry", "Twinkle",
                       "Candycane", "Gingersnap", "Sparkle", "Jolly", "Chestnut", "Cinnamon", "Pinecone", "Kringle",
                       "Noel", "Bells", "Sprinkle"]

    """
    def test_reset(self):
        file1 = open("elves.txt", "w")
        file1.truncate(0)
        for elf in self.elves_duplicates:
            file1.write(elf + "\n")
        file1.close()
        file2 = open("elves_unique.txt", "w")
        file2.truncate(0)
        file2.close()
    """

    def test_length(self):
        file1 = open("elves_unique.txt", "r")
        elves = file1.readlines()
        file1.close()
        flag= len(self.elves_singleton)==len(elves)
        if not flag:
            print("Lists differ in length")
        assert flag

    def test_equality(self):
        file1 = open("elves_unique.txt", "r")
        elves = file1.readlines()

        file1.close()
        flag= self.elves_singleton.sort() == elves.sort()
        if not flag:
            print("Lists are unequal")
        assert flag

    if __name__ == '__main__':
        unittest.main()
