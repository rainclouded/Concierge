import unittest
from app import app
from app.database.DatabaseController import DatabaseController
from app.authentication.AuthenticationManager import AuthenticationManager
from app.database.Mockdata import Mockdata

class TestUserService(unittest.TestCase):
    
    def setUp(self):
        self.database = Mockdata()
        self.database_controller = DatabaseController(self.database)
        self.am = AuthenticationManager(self.database_controller)

    def test_create_new_guest(self):
        pass
    def test_create_new_staff(self):
        pass
    def test_delete_user(self):
        pass
