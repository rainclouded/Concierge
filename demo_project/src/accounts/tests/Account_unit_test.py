import unittest
import json
from app import app


class TestFlaskApp(unittest.TestCase):
    def setUp(self):
        self.client = app.test_client()
        self.client.testing = True
        self.app = app

    def test_index(self):
        response = self.client.get("/accounts/")
        self.assertEqual(response.status_code, 200)
        self.assertEqual(
            response.json,
            {"message": "Hello, World From Accounts", "status": "success"},
        )

    def test_login_success(self):
        response = self.client.post(
            "/accounts/login_attempts",
            data=json.dumps({"username": "admin", "password": "admin"}),
            content_type="application/json",
        )
        self.assertEqual(response.status_code, 200)
        self.assertEqual(
            response.json,
            {
                "key": None,
                "account_id": None,
                "message": "Welcome, admin!",
                "status": "ok",
            },
        )

    def test_login_failure(self):
        response = self.client.post(
            "/accounts/login_attempts",
            data=json.dumps({"username": "admin", "password": "wrong"}),
            content_type="application/json",
        )
        self.assertEqual(response.status_code, 200)
        self.assertTrue(response.json["message"].startswith("Login Fail"))
        self.assertEqual(response.json["status"], "error")

    def test_login_attempt_count(self):
        self.client.post(
            "/accounts/login_attempts",
            data=json.dumps({"username": "admin", "password": "wrong"}),
            content_type="application/json",
        )
        self.client.post(
            "/accounts/login_attempts",
            data=json.dumps({"username": "admin", "password": "wrong"}),
            content_type="application/json",
        )
        response = self.client.post(
            "/accounts/login_attempts",
            data=json.dumps({"username": "admin", "password": "wrong"}),
            content_type="application/json",
        )
        self.assertEqual(response.json["message"], "Login Fail #2: Invalid Credentials")
        self.assertEqual(response.json["status"], "error")


if __name__ == "__main__":
    unittest.main()
