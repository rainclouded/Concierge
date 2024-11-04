import unittest

if __name__=='__main__':
    loader = unittest.TestLoader()
    suite = loader.discover('.', pattern='*unit_test.py')
    testrunner = unittest.TextTestRunner()
    testResult = testrunner.run(suite)
    exit(0 if testResult.wasSuccessful() else 1)