
"""
Testing module for FlaskApp
"""
import unittest
import json
import datetime
import jwt
from cryptography.hazmat.primitives import serialization
from app import app
from app.server.Accounts import get_port, set_services
from app.authentication.AuthenticationManager import AuthenticationManager
from app.database.DatabaseController import DatabaseController
from app.user_service.UserService import UserService
from app.validation.ValidationManager import ValidationManager
from app.permissions.PermissionController import PermissionController
from app.database.Mockdata import Mockdata
from app.permissions.MockPermissions import MockPermissions

class TestFlaskApp(unittest.TestCase):
    """
    Tests for FlaskApp
    """

    def setUp(self):
        database = DatabaseController(Mockdata())
        authenticaion = AuthenticationManager(database)
        validation = ValidationManager(database)
        permissions = PermissionController(MockPermissions())
        user_service = UserService(database,  authenticaion, validation)
        set_services(database, authenticaion, user_service, permissions)

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
            data=json.dumps(
                {
                    'username': 'newUser1', 
                    'type' : 'staff', 
                    'password' : 'newPassword1'
                    }
            ),
            content_type='application/json'
            )
        self.assertEqual(response.status_code, 200)
        self.assertEqual(
            response.json,
            {
                'message':'User created successfully. password: newPassword1',
                'status' : 'success'
            }
        )
        self.assertEqual(response.status_code, 200)

        response = self.client.post(
            '/accounts',
            data=json.dumps(
                {
                    'username': '',
                    'type' : 'staff',
                    'password' : 'newPassword1'
                }
            ),
            content_type='application/json'
            )
        self.assertEqual(response.status_code, 401)
        self.assertEqual(
            response.json,
            {'message':'Could not create user', 'status' : 'error'}
        )


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
                'data':{'id':'1','type':'staff','username':'test1'},
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
                'data':{'id':'1','type':'guest','username':'5'},
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

        self.assertEqual(response.status_code, 401)
        self.assertTrue(response.json['message'].startswith('Login Fail'))
        self.assertEqual(response.json['status'], 'error')


    def test_get_port(self):
        self.assertEqual(8080,get_port())


    def test_delete(self):
        test_private_key= serialization.load_pem_private_key(
            """-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEIPzzVpfUDwUXOc+rF3o/FUdACe7hhv7dl5hgmFIJoR3ooAoGCCqGSM49 AwEHoUQDQgAE+tognnc+cFv4SK9KTuw7BIAVkZKrET7NVlEYW+n+4XMSlK8ZOlUT uYw35b6aJsT7GWrGGsOBE7I+g3x6nikmxg==\n-----END EC PRIVATE KEY-----""".encode('utf-8'),
            password = None
        )
        token = jwt.encode(
                {
                    'expiry':(
                        datetime.datetime.now()
                        +datetime.timedelta(days=999)
                        ).timestamp()
                },
                test_private_key,
                algorithm='ES256'
            )

        response = self.client.post(
            '/accounts/delete',
            headers={'X-Api-Key':token},
            data=json.dumps({'username': 'test1'}),
            content_type='application/json',
        )
        self.assertEqual(response.status_code, 200)
        self.assertEqual(
            response.json,
            {
                'message': 'test1 Successfully deleted!',
                'status': 'ok',
            },
        )
        response = self.client.post(
            '/accounts/delete',
            headers={'X-Api-Key':token},
            data=json.dumps({'username': 'not_real'}),
            content_type='application/json',
        )
        self.assertEqual(response.status_code, 403)
        self.assertEqual(
            response.json,
            {
                'message': 'Action not permitted',
                'status': 'forbidden',
            },
        )

        token = jwt.encode(
                {
                    'expiry':(
                        datetime.datetime.now()
                        -datetime.timedelta(days=999)
                        ).timestamp()
                },
                test_private_key,
                algorithm='ES256'
            )


    def test_update(self):
        test_private_key= serialization.load_pem_private_key(
            """-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEIPzzVpfUDwUXOc+rF3o/FUdACe7hhv7dl5hgmFIJoR3ooAoGCCqGSM49 AwEHoUQDQgAE+tognnc+cFv4SK9KTuw7BIAVkZKrET7NVlEYW+n+4XMSlK8ZOlUT uYw35b6aJsT7GWrGGsOBE7I+g3x6nikmxg==\n-----END EC PRIVATE KEY-----""".encode('utf-8'),
            password = None
        )
        token = jwt.encode(
                {
                    'expiry':(
                        datetime.datetime.now()
                        +datetime.timedelta(days=999)
                        ).timestamp()
                },
                test_private_key,
                algorithm='ES256'
            )

        response = self.client.put(
            '/accounts/update',
            headers={'X-Api-Key':token},
            data=json.dumps({'username': '5'}),
            content_type='application/json',
        )
        self.assertEqual(response.status_code, 200)
        self.assertRegex(
            response.json['message'],
            r'5 Successfully updated! password: \d+'
        )

        response = self.client.put(
            '/accounts/update',
            headers={'X-Api-Key':token},
            data=json.dumps({'username': '99'}),
            content_type='application/json',
        )
        self.assertEqual(response.status_code, 401)
        self.assertEqual(
            response.json,
            {
                'message': 'Update could not be completed.',
                'status': 'error',
            },
        )
        