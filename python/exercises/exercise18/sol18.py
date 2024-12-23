import unittest

from ex18 import assignSeats

class MyTestCase(unittest.TestCase):
    elves = ['Jingle', 'Tinsel', 'Peppermint', 'Snowball', 'Holly', 'Sugarplum', 'Frosty', 'Merry', 'Twinkle',
             'Candycane',
             'Gingersnap', 'Sparkle', 'Sprinkle', 'Jolly', 'Chestnut', 'Cinnamon', 'Pinecone', 'Kringle',
             'Noel', 'Bells']
    def test_allAssigned(self):
        assignments=assignSeats(self.elves)
        for elf in self.elves:
            assert elf in assignments

    def test_NoAdditions(self):
        assignments=assignSeats(self.elves)
        assert len(assignments)==len(self.elves)

    def test_emptyList(self):
        with self.assertRaises(ValueError):
            assignSeats([])


if __name__ == '__main__':
    unittest.main()
