import unittest
from app import app
from app.validation.ValidationManager import ValidationManager
from app.database.Mockdata import Mockdata
from app.database.DatabaseController import DatabaseController
from app.dto.UserObject import UserObject as User

class TestAuthenticationManager(unittest.TestCase):
    TEST_DATA = [
        {
            'username' : 'test1',
            'id' : '1',
            'hash' : '',
            'type' : 'staff'
        },
        {
            'username' : 'test2',
            'id' : '2',
            'hash' : '',
            'type' : 'staff'
        },
        {
            'username' : 'test3',
            'id' : '3',
            'hash' : '',
            'type' : 'staff'
        },
        {
            'username' : '5',
            'id' : '',
            'hash' : '',
            'type' : 'guest'
        },
        {
            'username' : '6',
            'id' : '',
            'hash' : '',
            'type' : 'guest'
        },
        {
            'username' : '7',
            'id' : '',
            'hash' : '',
            'type' : 'guest'
        }
    ]

    def setUp(self):
        self.database = Mockdata()
        self.database.users = self.TEST_DATA
        self.database_controller = DatabaseController(self.database)
        self.validation = ValidationManager(self.database_controller)


    def test_validate_staff_password(self):
        invalid_passwords = [
            'password',
            '8888888',
            'shor',
            None,
            '',
            '@#$%^'
        ]

        valid_passwords = [
            'validPassword1',
            '9999999u',
            '!@#$%^&*()0P',
            '        9p'
        ]
        for password in valid_passwords:
            self.assertTrue(self.validation.validate_staff_password(password))
        for password in invalid_passwords:
            self.assertFalse(self.validation.validate_staff_password(password))


    def validate_staff_username(self):
        invalid_names = [
            'username',
            '8888888',
            'shor',
            None,
            '',
            '@#$%^'
        ]

        valid_names = [
            'validUsername1',
            '9999999u',
            '!@#$%^&*()0P',
            '        9p'
        ]

        for username in valid_names:
            self.assertTrue(self.validation.validate_staff_username(username))
        for username in invalid_names:
            self.assertFalse(self.validation.validate_staff_username(username))



    def test_validate_new_staff(self):
        valid_users = [
            (
                User(
                    **{
                        'username' : 'newUser1',
                        'id' : '',
                        'hash' : '',
                        'type' : 'staff'
                    }
                ),
                'newPassword1'
            ),
            (
                User(
                    **{
                        'username' : 'newUser2',
                        'id' : '',
                        'hash' : '',
                        'type' : 'staff'
                    }
                ),
                'newPassword2'
            )

        ]
        invalid_users = [
            (
                User(
                    **{
                        'username' : '',
                        'id' : '',
                        'hash' : '',
                        'type' : 'staff'
                    }
                ),
                'newPassword1'
            ),
            (
                User(
                    **{
                        'username' : 'newUser1',
                        'id' : '',
                        'hash' : '',
                        'type' : 'staff'
                    }
                ),
                ''
            ),
            (
                User(
                    **{
                        'username' : 'newuserfailure',
                        'id' : '',
                        'hash' : '',
                        'type' : 'staff'
                    }
                ),
                'newPassword1'
            ),
            (
                User(
                    **{
                        'username' : 'newUser1',
                        'id' : '',
                        'hash' : '',
                        'type' : 'staff'
                    }
                ),
                'newpasswordfailure'
            ),
            (
                User(
                    **{
                        'username' : 'test1',
                        'id' : '',
                        'hash' : '',
                        'type' : 'staff'
                    }
                ),
                'newPassword1'
            ),
            (
                User(
                    **{
                        'username' : '5',
                        'id' : '',
                        'hash' : '',
                        'type' : 'staff'
                    }
                ),
                'newPassword1'
            ),
            (
                User(
                    **{
                        'username' : None,
                        'id' : '',
                        'hash' : '',
                        'type' : 'staff'
                    }
                ),
                'newPassword1'
            ),
            (
                None,
                'newPassword1'
            ),
            (
                'username',
                None
            )
        ]

        for user, password in valid_users:
            self.assertTrue(self.validation.validate_new_staff(user, password))
        for user, password in invalid_users:
            self.assertFalse(self.validation.validate_new_staff(user, password))
