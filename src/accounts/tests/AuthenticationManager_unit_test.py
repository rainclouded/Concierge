import unittest
from app import app
from app.database.DatabaseController import DatabaseController
from app.authentication.AuthenticationManager import AuthenticationManager
from app.database.Mockdata import Mockdata

class TestAuthenticationManager(unittest.TestCase):
    
    def setUp(self):
        self.database = Mockdata()
        self.database_controller = DatabaseController(self.database)
        self.am = AuthenticationManager(self.database_controller)

    def test_get_hash(self):
        pass
    def test_check_hash(self):
        pass
    def test_authenticate_staff_login(self):
        pass