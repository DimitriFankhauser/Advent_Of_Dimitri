import unittest
import ex10

class MyTestCase(unittest.TestCase):
    nums = [9,4,1,4,2,8]

    def test_something(self):
        self.assertEqual(self.nums, ex10.numbers_found)


if __name__ == '__main__':
    unittest.main()
