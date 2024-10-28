import unittest
from unittest.mock import patch
from app import app
from app.database.DatabaseController import DatabaseController
from app.user_service.UserService import UserService
from app.database.Mockdata import Mockdata
from app.dto.UserObject import UserObject as User
from app.authentication.AuthenticationManager import AuthenticationManager
from app.validation.ValidationManager import ValidationManager
import app.Configs as cfg

class TestUserService(unittest.TestCase):
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
        self.database.users = [*self.TEST_DATA]#Deepcopy hack
        self.database_controller = DatabaseController(self.database)
        self.us = UserService(
            self.database_controller,
            AuthenticationManager(self.database_controller),
            ValidationManager(self.database_controller)
            )


    @patch('app.user_service.UserService.randbelow')
    def test_create_new_guest(self, mock_randbelow):
        mock_randbelow.return_value = 500
        new_guest = User(
            **{
            'username' : '8',
            'id' : '',
            'hash' : '',
            'type' : 'guest'
            }
        )

        user, _ = self.us.create_new_guest(new_guest)

        mock_randbelow.assert_called_with(cfg.MAX_GUEST_PASSWORD)
        self.assertEqual(user, new_guest)


    def test_create_new_staff(self):
        new_user = User(
            **{
            'username' : 'newUser99',
            'id' : '',
            'hash' : '',
            'type' : 'staff'
            }
        )
        valid_staff = User(
            **{
            'username' : 'newUser99',
            'id' : 4,
            'hash' : '7668b598580e46d2e6347841824c847c84994feaffff228aadfacf059e5ef5bb',
            'type' : 'staff'
            }
        )
        user = self.us.create_new_staff(User(**{}), 'nothing')

        self.assertIsNone(user)

        user = self.us.create_new_staff(new_user, 'testPassword99')

        self.assertTrue(valid_staff, user)


    def test_delete_user(self):
        remaining_valid_users = [
            User(**{
                'username' : 'test1',
                'id' : '1',
                'hash' : '',
                'type' : 'staff'
            }),
            User(**{
                'username' : 'test2',
                'id' : '2',
                'hash' : '',
                'type' : 'staff'
            }),
            User(**{
                'username' : '5',
                'id' : '',
                'hash' : '',
                'type' : 'guest'
            }),
            User(**{
                'username' : '6',
                'id' : '',
                'hash' : '',
                'type' : 'guest'
            })
        ]
        staff_to_delete = User(
            **{
                'username' : 'test3',
                'id' : '3',
                'hash' : '',
                'type' : 'staff'
            }
        )

        guest_to_delete = User(
            **{
                'username' : '7',
                'id' : '',
                'hash' : '',
                'type' : 'guest'
            }
        )
        self.assertIsNone(self.us.delete_user(User(**{})))
        self.assertEqual(
            self.us.delete_user(guest_to_delete.username),
            guest_to_delete
            )
        self.assertCountEqual(
            self.database_controller.get_users(),
            [*remaining_valid_users, staff_to_delete]
            )

        self.assertEqual(
            self.us.delete_user(staff_to_delete.username),
            staff_to_delete
            )
        self.assertCountEqual(
            self.database_controller.get_users(),
            remaining_valid_users
            )
