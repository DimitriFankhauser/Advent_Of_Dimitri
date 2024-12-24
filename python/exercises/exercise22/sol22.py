import unittest

from python.exercises.exercise22.ex22 import checkPassword,BadPasswordError


class MyTestCase(unittest.TestCase):
    def test_happy_case(self):
        response=checkPassword("asdFj8!")
        assert response is True

    def test_multiplicity(self):
        with self.assertRaises(BadPasswordError):
            checkPassword("$$$$$$$$")
        with self.assertRaises(BadPasswordError):
            checkPassword("AAAAAAAA")
        with self.assertRaises(BadPasswordError):
            checkPassword("aaaaaaaa")
        with self.assertRaises(BadPasswordError):
            checkPassword("00000000")
    def test_multiplicity_in_normal(self):
        with self.assertRaises(BadPasswordError):
            checkPassword("asddddFj8!")
        with self.assertRaises(BadPasswordError):
            checkPassword("asdFFFFj8!")
        with self.assertRaises(BadPasswordError):
            checkPassword("asdFj8888!")
        with self.assertRaises(BadPasswordError):
            checkPassword("asdFj8!!!!")

    def test_length(self):
        def test_multiplicity_in_normal(self):
            with self.assertRaises(BadPasswordError):
                checkPassword("asd!8")
            with self.assertRaises(BadPasswordError):
                checkPassword(" ")

if __name__ == '__main__':
    unittest.main()
