import unittest
from datetime import date

from ex09 import calculate_days_til_christmas

class sol09(unittest.TestCase):

    def test_days(self):
        today = date.today()
        christmas = date.fromisoformat(str(today.year)+"-12-25")
        td = (christmas - today).days
        calulatedDiff=calculate_days_til_christmas()
        self.assertEqual(td, calulatedDiff)




