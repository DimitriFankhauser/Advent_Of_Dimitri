import unittest

from ex04 import main


class MyTestCase(unittest.TestCase):
    def test_normalcase(self):
        result = main("//|\||")
        self.assertEqual("//*|*\*|**|*", result)

    def test_no_straight_branches(self):
        result = main("//\/\/")
        self.assertEqual("No straight branch in this tree", result)

    def test_only_straight_branches(self):
        result = main("|||")
        self.assertEqual("*|**|**|*", result)
