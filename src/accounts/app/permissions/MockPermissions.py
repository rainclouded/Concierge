import jwt
import datetime
from app.permissions.PermissionInterface import PermissionInterface
from cryptography.hazmat.primitives import serialization

class MockPermissions(PermissionInterface):

    def __init__(self):
        self._key_pair = (None, None)
        self._test_message = {}


    def can_delete_guest(self, token:str, public_key:str)->bool:
        return (
            self.decode_token(token, public_key)['expiry']
            >= datetime.datetime.now().timestamp()
        )
    

    def can_delete_staff(self, token:str, public_key:str)->bool:
        return (
            self.decode_token(token, public_key)['expiry']
            >= datetime.datetime.now().timestamp()
        )
    
    def can_update_staff(self, token: str, public_key:str)->bool:
        try:
            return (
                self.decode_token(token, public_key)['expiry']
                >= datetime.datetime.now().timestamp()
            )
        except jwt.PyJWTError as e:
            raise e

    def can_update_guest(self, token: str, public_key:str)->bool:
        try:
            return (
                self.decode_token(token, public_key)['expiry']
                >= datetime.datetime.now().timestamp()
            )
        except jwt.PyJWTError as e:
            raise e

    def decode_token(self, token:str, public_key:str,\
                     algorithm:str='ES256'):
        try:
            return jwt.decode(
                token,
                serialization.load_pem_public_key(public_key.encode()),
                [algorithm]
                )
        except jwt.PyJWTError as e:
            raise e
    
    def get_public_key(self):
        return (
            """-----BEGIN PUBLIC KEY-----MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE+tognnc+cFv4SK9KTuw7BIAVkZKr ET7NVlEYW+n+4XMSlK8ZOlUTuYw35b6aJsT7GWrGGsOBE7I+g3x6nikmxg==-----END PUBLIC KEY-----"""
        )