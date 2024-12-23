import unittest

from python.exercises.exercise19.ex19 import Elf


class MyTestCase(unittest.TestCase):

    elf=Elf("Sparkles",123,55,1234,"Dancing")

    def test_increment_age(self):
        self.elf.increment_age()
        self.assertEqual(self.elf.age,1235)

    def test_changeWeight(self):
        self.elf.change_weight(59)
        self.assertEqual(self.elf.weight,59)

    def test_change_skill(self):
        self.elf.change_skill("Python")
        self.assertEqual(self.elf.special_skill,"Python")


if __name__ == '__main__':
    unittest.main()
