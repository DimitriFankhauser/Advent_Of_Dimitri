import unittest

from python.exercises.exercise22.ex22 import checkPassword,BadPasswordError
from python.exercises.exercise23.ex23 import convertToBinary, convertToDecimal


class MyTestCase(unittest.TestCase):
    def test_minimum(self):
        bin=convertToBinary(0)
        dec=convertToDecimal("00000000")
        self.assertEqual(bin,"00000000")
        self.assertEqual(dec,0)

    def test_maximum(self):
        bin=convertToBinary(255)
        dec=convertToDecimal("11111111")
        self.assertEqual(bin,"11111111")
        self.assertEqual(dec,255)

    def test_normalCase(self):
        bin=convertToBinary(122)
        dec=convertToDecimal("01111010")
        self.assertEqual(bin,"01111010")
        self.assertEqual(dec,122)

if __name__ == '__main__':
    unittest.main()
