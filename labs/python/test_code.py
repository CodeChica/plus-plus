import unittest

from maths import *

class TestSum(unittest.TestCase):

    def test_sum(self):
        self.assertEqual(sum(1, 1), 2, "1 + 1 should be 2")
        self.assertEqual(sum(5, 12), 17, "5 + 12 should be 17")

    @unittest.skip("extension task")
    def test_multiply(self):
        self.assertEqual(multiply(1, 1), 1, "1 * 1 should be 1")
        self.assertEqual(multiply(5, 10), 50, "5 * 10 should be 50")

if __name__ == '__main__':
    unittest.main()
