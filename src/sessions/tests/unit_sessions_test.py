import unittest
from unittest.mock import patch

import sys
import os
sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), '../../')))
from app import app


class FlaskAppTests(unittest.TestCase):
    def setUp(self):
        self.app = app.test_client()
        self.app.testing = True

    def test_get_sessions(self):
        response = self.app.get("/sessions/")
        self.assertEqual(response.status_code, 200)
        self.assertEqual(
            response.json,
            {"message": "Hello, World From Sessions", "status": "success"},
        )

    def test_post_sessions_success(self):
        with patch("requests.post") as mock_post:
            mock_post.return_value.json.return_value = {"status": "ok"}
            mock_post.status_code = 200
            
            response = self.app.post(
                "/sessions/", json={"username": "user", "password": "pass"}
            )
            self.assertEqual(response.status_code, 200)
            data = response.json
            self.assertEqual(data["status"], "ok")
            self.assertIn("session_key", data)

    def test_post_sessions_failure(self):
        with patch("requests.post") as mock_post:
            mock_post.return_value.json.return_value = {
                "status": "error",
                "message": "Invalid credentials",
            }
            mock_post.status_code = 200
            
            response = self.app.post(
                "/sessions/", json={"username": "user", "password": "wrong_pass"}
            )
            self.assertEqual(response.status_code, 400)
            data = response.json
            self.assertEqual(data["status"], "error")
            self.assertEqual(data["message"], "Invalid credentials")

    def test_post_sessions_no_username(self):
        with patch("requests.post") as mock_post:
            mock_post.return_value.json.return_value = {
                "status": "error",
                "message": "Missing username and or password",
            }
            mock_post.status_code = 400
            response = self.app.post("/sessions/", json={})
            self.assertEqual(response.status_code, 400)
            self.assertEqual(
                response.json, {"status": "error", "message": "Missing username and or password"}
            )

    def test_invalid_method(self):
        response = self.app.put("/sessions/")
        self.assertEqual(response.status_code, 405)


if __name__ == "__main__":
    unittest.main()
