import unittest
from app.core.Services import Services

class TestServices(unittest.TestCase):


    def setUp(self):
        self.services = Services


    def test_set_up(self):
        self.assertIsNone(Services.get_database())
        self.assertIsNone(Services.get_authentication())
        self.assertIsNone(Services.get_user_service())
        self.assertIsNone(Services.get_permissions())
        self.assertIsNone(Services.get_validation())

        Services.set_up()

        self.assertIsNotNone(Services.get_database())
        self.assertIsNotNone(Services.get_authentication())
        self.assertIsNotNone(Services.get_user_service())
        self.assertIsNotNone(Services.get_permissions())
        self.assertIsNotNone(Services.get_validation())
