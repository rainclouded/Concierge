import unittest
import json
import re
from unittest.mock import MagicMock, patch
from app import app


class TestFlaskApp(unittest.TestCase):


    def setUp(self):
        self.client = app.test_client()
        self.client.testing = True
        self.app = app


    def test_index_get(self):
        response = self.client.get('/accounts')
        self.assertEqual(response.status_code, 200)
        self.assertEqual(
            response.json,
            {
                'message': 'You have contacted the accounts',
                'status': 'success'
            }
        )


    def test_index_post(self):
        response = self.client.post(
            '/accounts',
            data=json.dumps({'username': '4', 'type' : 'guest'}),
            content_type='application/json'
            )
        self.assertEqual(response.status_code, 200)
        self.assertRegex(
            response.json['message'],
            r'User created successfully. password: \d+'
        )
        self.assertEqual(response.status_code, 200)

        response = self.client.post(
            '/accounts',
            data=json.dumps({'username': 'newUser1', 'type' : 'staff', 'password' : 'newPassword1'}),
            content_type='application/json'
            )
        self.assertEqual(response.status_code, 200)
        self.assertEqual(
            response.json,
            {'message':'User created successfully. password: newPassword1', 'status' : 'success'}
        )
        self.assertEqual(response.status_code, 200)

        response = self.client.post(
            '/accounts',
            data=json.dumps({'username': '', 'type' : 'staff', 'password' : 'newPassword1'}),
            content_type='application/json'
            )
        self.assertEqual(response.status_code, 200)
        self.assertEqual(
            response.json,
            {'message':'Could not create user', 'status' : 'error'}
        )
        self.assertEqual(response.status_code, 200)


    def test_login_success(self):
        response = self.client.post(
            '/accounts/login_attempt',
            data=json.dumps({'username': 'test1', 'password': 'testPassword1'}),
            content_type='application/json',
        )
        self.assertEqual(response.status_code, 200)
        self.assertEqual(
            response.json,
            {
                'message': 'Welcome, test1!',
                'status': 'ok',
            },
        )
        response = self.client.post(
            '/accounts/login_attempt',
            data=json.dumps({'username': '5', 'password': '44444444'}),
            content_type='application/json',
        )
        self.assertEqual(response.status_code, 200)
        self.assertEqual(
            response.json,
            {
                'message': 'Welcome, 5!',
                'status': 'ok',
            },
        )

    def test_login_failure(self):
        response = self.client.post(
            '/accounts/login_attempt',
            data=json.dumps({'username': 'nouser', 'password': 'nopass'}),
            content_type='application/json',
        )
        self.assertEqual(
            response.json,
            {
                'message': 'Login Fail - Invalid Credentials',
                'status': 'error',
            },
        )
            
        self.assertEqual(response.status_code, 200)
        self.assertTrue(response.json['message'].startswith('Login Fail'))
        self.assertEqual(response.json['status'], 'error')



if __name__ == '__main__':
    unittest.main()
