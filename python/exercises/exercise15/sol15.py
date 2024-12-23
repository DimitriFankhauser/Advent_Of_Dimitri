import unittest
from ex15 import findPrimeNumber

class MyTestCase(unittest.TestCase):
    def test_firstNumPrime(self):
        nums=[2,24,55,26,75,100]
        assert findPrimeNumber(nums) ==2

    def test_lastNumPrime(self):
        nums=[24,55,10,23]
        assert findPrimeNumber(nums) ==23

    def test_primeInTheMiddle(self):
        nums = [24, 55, 7,10, 23]
        assert findPrimeNumber(nums) == 7

    def test_negatives(self):
        nums=[24,55,10,-134,4,8]
        assert findPrimeNumber(nums) ==-1

    def test_OutOfRange(self):
        nums=[24,55,10,134,4,8]
        assert findPrimeNumber(nums) ==-1

if __name__ == '__main__':
    unittest.main()
