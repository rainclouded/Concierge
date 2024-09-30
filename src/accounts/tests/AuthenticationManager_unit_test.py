import unittest
import json
from app import app
from app.AuthenticationManager import AuthenticationManager

class TestAuthenticationManager(unittest.TestCase):
    def setUp(self):
        self.am = AuthenticationManager()

    def test_check_hash(self):
        pass