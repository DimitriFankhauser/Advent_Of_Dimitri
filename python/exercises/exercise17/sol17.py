import unittest
from ex17 import findYears


class MyTestCase(unittest.TestCase):
    years = [2025, 2031, 2036, 2042, 2053, 2059, 2064, 2070, 2081, 2087, 2092, 2098, 2104, 2110, 2121]

    def test_years(self):
        self.assertEqual(self.years, findYears())  # add assertion here

    def test_length(self):
        self.assertEqual(len(self.years), len(findYears()))


if __name__ == '__main__':
    unittest.main()
