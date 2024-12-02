import unittest

from ex03 import main


class MyTestCase(unittest.TestCase):
    def test_hansremoved(self):
        victims=["Leopold", "Hans", "Hanspeter", "Persephone"]
        result = main(victims)
        victims_hans_removed=["Leopold","Hanspeter", "Persephone"]
        self.assertEqual( victims_hans_removed,result)

