import unittest

from main import compare

class TestComparisonMethod(unittest.TestCase):
    def test_ver1_greater_example_1(self):
        self.assertEqual(compare('1.2', '1.1'), 1)
    
    def test_ver1_greater_example_2(self):
        self.assertEqual(compare('1.2.9.9.9.9', '1.2.9'), 1)
    
    def test_ver1_greater_example_3(self):
        self.assertEqual(compare('1.10', '1.1'), 1)
    
    def test_ver1_greater_example_4(self):
        self.assertEqual(compare('0.1', '0.0.1'), 1)
    
    def test_ver1_equal_ver2_example_1(self):
        self.assertEqual(compare('1.1', '1.1'), 0)
    
    def test_ver1_equal_ver2_example_2(self):
        self.assertEqual(compare('01.1', '1.01'), 0)
    
    def test_ver1_equal_ver2_example_3(self):
        self.assertEqual(compare('01.01', '1.1'), 0)
    
    def test_ver1_equal_ver2_example_4(self):
        self.assertEqual(compare('1.2.9.9.9.9', '1.2.9.9.9.9'), 0)

    def test_ver2_greater_example_1(self):
        self.assertEqual(compare('0.1', '1.1'), -1)

    def test_ver2_greater_example_2(self):
        self.assertEqual(compare(' 1.3', '1.3.4'), -1)

    def test_ver2_greater_example_3(self):
        self.assertEqual(compare('0.0.0.0.1', '1.0.0.01'), -1)

    def test_ver2_greater_example_3(self):
        self.assertEqual(compare('1.2', '1.2.9.9.9.9'), -1)

if __name__ == '__main__':
    unittest.main()