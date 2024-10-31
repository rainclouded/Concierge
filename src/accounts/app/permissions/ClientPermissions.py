"""
Module for ClientPermissions
"""

import json
import os
import requests
import jwt
from app.permissions.PermissionInterface import PermissionInterface
from cryptography.hazmat.primitives import serialization

class ClientPermissionValidator(PermissionInterface):
    """
    ClientPermissions Integrates with the permissions service
    """
    CAN_DELETE_GUEST = "canDeleteGuestsAccounts"
    CAN_DELETE_STAFF = "canDeleteStaffAccounts"
    CAN_EDIT_STAFF =   "canEditStaffAccounts"
    CAN_EDIT_GUEST =   "canEditGuestAccounts"


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
        return self.validate_session_key_for_permission_name(token, self.CAN_DELETE_GUEST)


    def can_delete_staff(self, token:str, public_key:str)->bool:
        """Verifies if the token permits staff deletion

            Args:            
                token is the jwt Token to be verified
                public_key is the public key to decode the token
            Returns:
                True if deletion is permitted
                False otherwise
        """
        return self.validate_session_key_for_permission_name(token, self.CAN_DELETE_GUEST)



    def can_update_staff(self, token: str, public_key:str)->bool:
        """Verifies if the token permits staff update
        
            Args:            
                token is the jwt Token to be verified
                public_key is the public key to decode the token
            Returns:
                True if update is permitted
                False otherwise
        """
        return self.validate_session_key_for_permission_name(token, self.CAN_DELETE_GUEST)


    def can_update_guest(self, token: str, public_key:str)->bool:
        """Verifies if the token permits guest update

            Args:            
                token is the jwt Token to be verified
                public_key is the public key to decode the token
            Returns:
                True if update is permitted
                False otherwise
        """
        return self.validate_session_key_for_permission_name(token, self.CAN_DELETE_GUEST)
    
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
    
    @staticmethod
    def validate_session_key_for_permission_name(sessionKey: str, permissionName: str) -> bool:
        permissions = ClientPermissionValidator.get_session_permissions(sessionKey)
        return permissionName in permissions
    
    @staticmethod
    def get_session_permissions(sessionKey: str) -> list[int]:
        endpoint = os.getenv('SESSIONS_ENDPOINT')
        try:
            response = requests.get(f'{endpoint}/sessions/me')
            permissions = json.loads(response)['data']['sessionsData']['SessionPermissionList']
            return permissions
        except requests.exceptions.RequestException as e:
            print(e)
        return []