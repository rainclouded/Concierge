import unittest
from app import app
from app.AuthenticationManager import AuthenticationManager
from app.Mockdata import Mockdata

class TestAuthenticationManager(unittest.TestCase):
    def setUp(self):
        self.database = Mockdata()
        self.am = AuthenticationManager(self.database)

    def test_get_hash(self):
        pass
    def test_check_hash(self):
        pass
    def test_validate_staff_login(self):
        pass
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