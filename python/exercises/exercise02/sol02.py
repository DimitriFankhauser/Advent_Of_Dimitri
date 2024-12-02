import unittest

from ex02 import main


class MyTestCase(unittest.TestCase):
    def test_neg1(self):
        result = main(-1, 10)
        self.assertEqual(result, -1)

    def test_neg2(self):
        result = main(100, -210)
        self.assertEqual(result, -1)

    def test_neg3(self):
        result = main(-100, -210)
        self.assertEqual(result, -1)

    def test_pos1(self):
        result = main(20, 10)
        self.assertEqual(result, 200)


if __name__ == '__main__':
    main()
