"""
Module for MockPermissions
"""
import jwt
import datetime
from app.permissions.PermissionInterface import PermissionInterface
from cryptography.hazmat.primitives import serialization

class MockPermissions(PermissionInterface):
    """
    MockPermissions mocks the permission service
    """

    def __init__(self):
        self._key_pair = (None, None)
        self._test_message = {}


    def can_delete_guest(self, token:str, public_key:str)->bool:
        """Verifies if the token permits guest deletion

            Args:            
                token is the jwt Token to be verified
                public_key is the public key to decode the token
            Returns:
                True if deletion is permitted
                False otherwise
        """
        return (
            self.decode_token(token, public_key)['expiry']
            >= datetime.datetime.now().timestamp()
        )


    def can_delete_staff(self, token:str, public_key:str)->bool:
        """Verifies if the token permits staff deletion

            Args:            
                token is the jwt Token to be verified
                public_key is the public key to decode the token
            Returns:
                True if deletion is permitted
                False otherwise
        """
        return (
            self.decode_token(token, public_key)['expiry']
            >= datetime.datetime.now().timestamp()
        )


    def can_update_staff(self, token: str, public_key:str)->bool:
        """Verifies if the token permits staff update
        
            Args:            
                token is the jwt Token to be verified
                public_key is the public key to decode the token
            Returns:
                True if update is permitted
                False otherwise
        """
        try:
            return (
                self.decode_token(token, public_key)['expiry']
                >= datetime.datetime.now().timestamp()
            )
        except jwt.PyJWTError as e:
            raise e


    def can_view_users(self, token: str, public_key:str)->bool:
        """Verifies if the token permits account retrieval
        
            Args:            
                token is the jwt Token to be verified
                public_key is the public key to decode the token
            Returns:
                True if action is permitted
                False otherwise
        """
        try:
            return (
                self.decode_token(token, public_key)['expiry']
                >= datetime.datetime.now().timestamp()
            )
        except jwt.PyJWTError as e:
            raise e


    def can_update_guest(self, token: str, public_key:str)->bool:
        """Verifies if the token permits guest update

            Args:            
                token is the jwt Token to be verified
                public_key is the public key to decode the token
            Returns:
                True if update is permitted
                False otherwise
        """
        try:
            return (
                self.decode_token(token, public_key)['expiry']
                >= datetime.datetime.now().timestamp()
            )
        except jwt.PyJWTError as e:
            raise e

    def decode_token(self, token:str, public_key:str,\
                     algorithm:str='ES256'):
        """Decodes a jwt token

            Args:            
                token is the jwt Token to be decoded
                public_key is the public key to decode the token
            Returns:
                deconded token data
        """
        try:
            return jwt.decode(
                token.encode(),
                serialization.load_pem_public_key(public_key.encode()),
                [algorithm]
                )
        except jwt.PyJWTError as e:
            raise e


    def get_public_key(self):
        """
        Returns the testing public key
        """
        return (
            """-----BEGIN PUBLIC KEY-----MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE+tognnc+cFv4SK9KTuw7BIAVkZKr ET7NVlEYW+n+4XMSlK8ZOlUTuYw35b6aJsT7GWrGGsOBE7I+g3x6nikmxg==-----END PUBLIC KEY-----"""
        )
    