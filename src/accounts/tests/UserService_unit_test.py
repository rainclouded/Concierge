import unittest
from unittest.mock import patch
from app import app
from app.database.DatabaseController import DatabaseController
from app.user_service.UserService import UserService
from app.database.Mockdata import Mockdata
from app.dto.UserObject import UserObject as User
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
        self.database.users = self.TEST_DATA
        self.database_controller = DatabaseController(self.database)
        self.us = UserService(self.database_controller)

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
            'username' : '8',
            'id' : '',
            'hash' : '',
            'type' : 'staff'
            }
        )
    
        user = self.us.create_new_staff(new_user)

        self.assertEqual(new_user, user)

        
    def test_delete_user(self):
        pass
