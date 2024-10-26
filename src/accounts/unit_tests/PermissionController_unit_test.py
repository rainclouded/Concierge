import unittest
import jwt
import datetime
from cryptography.hazmat.primitives import serialization
from app.permissions.PermissionController import PermissionController
from app.permissions.MockPermissions import MockPermissions

class TestPermissionController(unittest.TestCase):


    def setUp(self):
        self.test_private_key= serialization.load_pem_private_key(
            """-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEIPzzVpfUDwUXOc+rF3o/FUdACe7hhv7dl5hgmFIJoR3ooAoGCCqGSM49 AwEHoUQDQgAE+tognnc+cFv4SK9KTuw7BIAVkZKrET7NVlEYW+n+4XMSlK8ZOlUT uYw35b6aJsT7GWrGGsOBE7I+g3x6nikmxg==\n-----END EC PRIVATE KEY-----""".encode('utf-8'),
            password = None
        )

        self.test_public_key= (
            """-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE+tognnc+cFv4SK9KTuw7BIAVkZKr ET7NVlEYW+n+4XMSlK8ZOlUTuYw35b6aJsT7GWrGGsOBE7I+g3x6nikmxg==\n-----END PUBLIC KEY-----"""
        )

        self.valid_tokens = [
            jwt.encode(
                {
                    'expiry':(
                        datetime.datetime.now()
                        +datetime.timedelta(days=999)
                        ).timestamp()
                },
                self.test_private_key,
                algorithm='ES256'
            ),
            jwt.encode(
                {
                    'other':'value',
                    'expiry':(
                        datetime.datetime.now()
                        +datetime.timedelta(days=999)
                        ).timestamp()
                },
                self.test_private_key,
                algorithm='ES256'
            )
        ]

        self.invalid_tokens = [
            'invalid',
            'not.a.token',
            jwt.encode(
                {
                    'expiry':(
                        datetime.datetime.now()
                        -datetime.timedelta(days=999)
                        ).timestamp()
                },
                self.test_private_key,
                algorithm='ES256'
            )[:-2]+'00'
        ]
        self.permissions = PermissionController(MockPermissions())


    def test_can_delete_staff(self):
        errors = [
            jwt.exceptions.DecodeError,
            jwt.exceptions.DecodeError,
            jwt.exceptions.InvalidSignatureError,

        ]
        for token in self.valid_tokens:
            self.assertTrue(
                self.permissions.can_delete_staff(token, self.test_public_key)
                )
            token = jwt.encode(
                {
                    'expiry':(
                        datetime.datetime.now()
                        -datetime.timedelta(days=999)
                        ).timestamp()
                },
                self.test_private_key,
                algorithm='ES256'
            )
            self.assertFalse(self.permissions.can_delete_staff(token, self.test_public_key))
        for token, error in zip(self.invalid_tokens, errors):
            with self.assertRaises(error):
                self.permissions.can_delete_staff(token, self.test_public_key)

    def test_can_delete_guest(self):
        errors = [
            jwt.exceptions.DecodeError,
            jwt.exceptions.DecodeError,
            jwt.exceptions.InvalidSignatureError,

        ]
        for token in self.valid_tokens:
            self.assertTrue(
                self.permissions.can_delete_guest(token, self.test_public_key)
                )
            token = jwt.encode(
                {
                    'expiry':(
                        datetime.datetime.now()
                        -datetime.timedelta(days=999)
                        ).timestamp()
                },
                self.test_private_key,
                algorithm='ES256'
            )
            self.assertFalse(self.permissions.can_delete_guest(token, self.test_public_key))
        for token, error in zip(self.invalid_tokens, errors):
            with self.assertRaises(error):
                self.permissions.can_delete_guest(token, self.test_public_key)

    def test_can_update_guest(self):
        errors = [
            jwt.exceptions.DecodeError,
            jwt.exceptions.DecodeError,
            jwt.exceptions.InvalidSignatureError,

        ]
        for token in self.valid_tokens:
            self.assertTrue(
                self.permissions.can_update_guest(token, self.test_public_key)
                )
            token = jwt.encode(
                {
                    'expiry':(
                        datetime.datetime.fromtimestamp(0)
                        ).timestamp()
                },
                self.test_private_key,
                algorithm='ES256'
            )
            self.assertFalse(self.permissions.can_update_guest(token, self.test_public_key))
        for token, error in zip(self.invalid_tokens, errors):
            with self.assertRaises(error):
                self.permissions.can_update_guest(token, self.test_public_key)


    def test_can_update_staff(self):
        errors = [
            jwt.exceptions.DecodeError,
            jwt.exceptions.DecodeError,
            jwt.exceptions.InvalidSignatureError,

        ]
        for token in self.valid_tokens:
            self.assertTrue(
                self.permissions.can_update_staff(token, self.test_public_key)
                )
            token = jwt.encode(
                {
                    'expiry':(
                        datetime.datetime.now()
                        -datetime.timedelta(days=999)
                        ).timestamp()
                },
                self.test_private_key,
                algorithm='ES256'
            )
            self.assertFalse(self.permissions.can_update_staff(token, self.test_public_key))
        for token, error in zip(self.invalid_tokens, errors):
            with self.assertRaises(error):
                self.permissions.can_update_staff(token, self.test_public_key)

