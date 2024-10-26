import unittest

if __name__=='__main__':
    loader = unittest.TestLoader()
    suite = loader.discover('.', pattern='*integration_test.py')
    testrunner = unittest.TextTestRunner()
    testrunner.run(suite)
