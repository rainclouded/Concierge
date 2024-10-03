import unittest
from app import app
from app.UserObject import UserObject as User

class TestAuthenticationManager(unittest.TestCase):

    def testCreation(self):
        test_data = [
            {
            'username' : 'testuser',
            'password' : 'testpass',
            'hash' : 'testhash',
            'id' : 'testid'
            }
        ]
        for test_case in test_data:
            self.assertDictEqual(test_case, User(test_data).__dict__)
        

