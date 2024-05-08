import unittest
from app import app
from app.AuthenticationManager import AuthenticationManager
from app.Mockdata import Mockdata

class TestAuthenticationManager(unittest.TestCase):
    def setUp(self):
        self.database = Mockdata()
        self.am = AuthenticationManager(self.database)

    def test_check_hash(self):
        pass