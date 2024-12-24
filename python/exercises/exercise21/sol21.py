import unittest

from python.exercises.exercise21.ex21 import checkBells


class MyTestCase(unittest.TestCase):
    reindeer = ['Blixem', 'Comet', 'Cupid', 'Dancer', 'Dasher', 'Dunder', 'Prancer', 'Rudolph', 'Vixen']
    ratings = [1, 17, 3, 4, 238,7,13,-7,2]
    def test_equality(self):
        res = dict(map(lambda i, j: (i, j), self.reindeer, self.ratings))
        reindeer_to_correct=checkBells(res)
        self.assertEqual(reindeer_to_correct,('Blixem', 'Comet', 'Dasher', 'Prancer', 'Rudolph', 'Vixen'))

    def test_length(self):
        res = dict(map(lambda i, j: (i, j), self.reindeer, self.ratings))
        reindeer_to_correct = checkBells(res)
        self.assertEqual(len(reindeer_to_correct), len(('Blixem', 'Comet', 'Dasher', 'Prancer', 'Rudolph', 'Vixen')))

if __name__ == '__main__':
    unittest.main()
