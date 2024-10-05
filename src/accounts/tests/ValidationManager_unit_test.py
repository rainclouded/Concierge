import unittest
from app import app
from app.validation.ValidationManager import ValidationManager
from app.database.Mockdata import Mockdata
from app.database.DatabaseController import DatabaseController

class TestAuthenticationManager(unittest.TestCase):
    def setUp(self):
        self.database = Mockdata()
        self.database_controller = DatabaseController(self.database)
        self.validation = ValidationManager(self.database_controller)

    def test_validate_staff_password(self):
        pass
    def validate_staff_username(self):
        pass
    def test_create_new_guest(self):
        pass
    def test_delete_user(self):
        pass
    def test_create_nw_staff(self):
        pass
    def test_validate_new_staff(self):
        pass