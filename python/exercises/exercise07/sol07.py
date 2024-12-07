import unittest


class MyTestCase(unittest.TestCase):
    presents = ["Chocolate", "Flowers", "LinuxONE", "Whiskey",
                "Cookies", "Friendship", "Hugs", "Puppet",
                "Candy cane", "Plant", "DVDs", "Clean Code-  by Robert C Martin", "Socks", "dog"]

    # run this test to reset your files
    """
    def test_reset(self):
        file1=open("presents.txt","w")
        for present in self.presents:
            file1.write(present+"\n")
        file1.close()

        file2=open("big_crates.txt","w")
        file2.truncate(0)
        file2.close()
        
        file3=open("small_crates.txt","w")
        file3.truncate(0)
        file3.close()
        """

    def test_small_crates(self):
        file1 = open("small_crates.txt", "r")
        lines = file1.readlines()
        correctlyPackaged = True
        for line in lines:
            line = line.strip()
            if len(line) > 5:
                correctlyPackaged = False
        assert correctlyPackaged

    def test_big_crates(self):
        file1 = open("big_crates.txt", "r")
        lines = file1.readlines()
        correctlyPackaged = True
        for line in lines:
            line = line.strip()
            if len(line) < 5:
                correctlyPackaged = False
        assert correctlyPackaged
