import unittest

from python.exercises.exercise14.Elf import Elf
from python.exercises.exercise14.ex14 import main


class MyTestCase(unittest.TestCase):
    def test_attributes(self):
        x=main()
        self.assertEqual(type(x), Elf)
        self.assertEqual(x.__class__, Elf)
        self.assertEqual(x.elf_name,"Sparkles")
        self.assertEqual(x.elf_id,24)
        self.assertEqual(x.balance,24.12)
        self.assertEqual(x.lactose_intolerant,True)


if __name__ == '__main__':
    unittest.main()

