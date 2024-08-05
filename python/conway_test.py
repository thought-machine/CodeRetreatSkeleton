import unittest

from python import conway

class TestConway(unittest.TestCase):

    def test_do_something(self):
        self.assertEqual(conway.DoSomething(), 123)

if __name__ == '__main__':
    unittest.main()
