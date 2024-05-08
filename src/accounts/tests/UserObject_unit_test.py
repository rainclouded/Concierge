import unittest
from app import app
from app.UserObject import UserObject as User

class TestUserObject(unittest.TestCase):

    def test_creation(self):
        test_data = [
            {
            'username' : 'testuser',
            'hash' : 'testhash',
            'id' : 'testid',
            'type' : 'guest'
            },
            {
            'username' : None,
            'hash' : None,
            'id' : None,
            'type' : None
            }
        ]
        for test_case in test_data:
            self.assertDictEqual(test_case, User(**test_case).__dict__)

    def test_comparison(self):
        test_user_1 = {
            'username' : 'testuser',
            'hash' : 'testhash',
            'id' : 'testid',
            'type' : 'guest'
        }
        test_user_2 = {
            'username' : 'testuser',
            'hash' : 'testhash',
            'type' : 'guest',
            'id' : 'testid',

        }
        test_user_3 = {
            'username' : 'testuser',
            'hash' : 'testhash',
            'extra_field_1' : None,
            'type' : 'guest',
            'id' : 'testid',
            'extra_field_2' : 'value'

            }

        test_user_4 = {
            'username' : 'wrong_name',
            'hash' : 'testhash',
            'type' : 'guest',
            'id' : 'testid',

        }
        self.assertTrue(User(**test_user_1) == User(**test_user_1))
        self.assertTrue(User(**test_user_1) == User(**test_user_2))
        self.assertTrue(User(**test_user_1) == User(**test_user_3))
        self.assertFalse(User(**test_user_1) == User(**test_user_4))



        

