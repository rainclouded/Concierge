import unittest

import sys
import os

sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), '../../')))
from app import app

class TestFlaskApp(unittest.TestCase):
    def setUp(self):
        self.client = app.test_client()
        self.client.testing = True
        self.app = app

    def test_index(self):
        response = self.client.get("/permissions/")
        self.assertEqual(response.status_code, 200)
        self.assertEqual(
            response.json,
            {"message": "Hello, World From Permissions", "status": "success"},
        )


if __name__ == "__main__":
    unittest.main()
