import unittest
from ex11 import *

class MyTestCase(unittest.TestCase):
    def test_santa_wins(self):
        attacks_santa = ["-", "+", "+", "-", "-", "-"]
        attacks_clown = ["-", "+", "-", "+", "0", "0"]
        self.assertEqual(Result.SANTA_WINS,main(attacks_santa,attacks_clown))
    def test_clown_wins(self):
        attacks_santa = ["-", "+", "+", "+", "0", "0"]
        attacks_clown = ["-", "+", "-", "-", "-", "-"]
        self.assertEqual(Result.CLOWN_WINS, main(attacks_santa,attacks_clown))

    def test_nobody_wins(self):
        attacks_santa = ["0", "0", "0", "0"]
        attacks_clown = ["0", "0", "0", "0"]
        self.assertEqual(Result.CLOWN_WINS, main(attacks_santa, attacks_clown))
if __name__ == '__main__':
    unittest.main()
