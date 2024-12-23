import unittest
from sol16 import calculate


# You may wonder: What is this file? This looks a lot different from everything else?
# This is a test. Don't worry, not an exam.
# The elves have produced a calculator for the good and mathematically gifted children.
# They programmed it (see sol16) but they want to be absolutely sure the calculator works as intended.
# The program works as follows: it takes 2 numbers and an operator (+,-,*,/, >,<,%)
# A unit-test is a test intended to check the smallest unit (typically a method or function).

# Some help:
# - tests must start with def test_<testname>(self):>
# - you can use assert <condition> to check if a condition is true

# For this exercise write tests for the operations named above. Yes there will be one or two small, logical errors in the code.
# Find and correct these. Then think about how and why you should test

class testCalculation(unittest.TestCase):

    def test_addition(self):
        result = calculate(1, 2, "+")
        self.assertEqual(result, 3)

    def test_type(self):
        result = calculate(1, 2, "+")
        self.assertFalse(type(result)==int)
        assert type(result) == float


if __name__ == '__main__':
    unittest.main()
