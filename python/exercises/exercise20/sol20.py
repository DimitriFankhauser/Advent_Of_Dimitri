import unittest

from python.exercises.exercise20.ex20 import collectReindeer


class MyTestCase(unittest.TestCase):
    factory = ["Dasher", "Rudolph", "Vixen"]
    outdoors = ["Dancer", "Comet", "Prancer"]
    fireplace = ["Cupid", "Dunder", "Blixem"]

    def test_length(self):
        reindeer=collectReindeer(factory=self.factory, outdoors=self.outdoors, fireplace=self.fireplace)
        assert len(reindeer) ==9

    def test_order(self):
        reindeer = collectReindeer(factory=self.factory, outdoors=self.outdoors, fireplace=self.fireplace)
        assert reindeer == ['Blixem', 'Comet', 'Cupid', 'Dancer', 'Dasher', 'Dunder', 'Prancer', 'Rudolph', 'Vixen']

    def test_errorhandling_all(self):
        with self.assertRaises(ValueError):
            collectReindeer(factory=[], outdoors=[],fireplace=[])

    def test_errorhandling_fac(self):
        with self.assertRaises(ValueError):
            collectReindeer(factory=[], outdoors=self.outdoors, fireplace=self.fireplace)

    def test_errorhandling_out(self):
        with self.assertRaises(ValueError):
            collectReindeer(factory=self.factory, outdoors=[], fireplace=self.fireplace)


    def test_errorhandling_out(self):
        with self.assertRaises(ValueError):
            collectReindeer(factory=self.factory, outdoors=self.outdoors, fireplace=[])

if __name__ == '__main__':
    unittest.main()
