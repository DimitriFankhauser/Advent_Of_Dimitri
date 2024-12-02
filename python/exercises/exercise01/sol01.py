import unittest

from ex01 import main


class MyTestCase(unittest.TestCase):
    def test_main(self):
        result=main()
        self.assertEqual(result,42)


if __name__ == '__main__':
    unittest.main()
