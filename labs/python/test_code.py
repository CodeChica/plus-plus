import unittest

from sum import sum

class TestSum(unittest.TestCase):

    def test_sum(self):
        self.assertEqual(sum(1, 1), 2, "1 + 1 should be 2")
        self.assertEqual(sum(5, 12), 17, "5 + 12 should be 17")

if __name__ == '__main__':
    unittest.main()
