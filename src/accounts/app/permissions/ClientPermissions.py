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
    CAN_VIEW_USERS =   "canViewAccounts"


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


    def can_view_users(self, token:str, public_key:str)->bool:
        """Verifies if the token permits user retrieval

            Args:            
                token is the jwt Token to be verified
                public_key is the public key to decode the token
            Returns:
                True if action is permitted
                False otherwise
        """
        return self.validate_session_key_for_permission_name(token, self.CAN_VIEW_USERS)


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
        """
        Validates the session key for the specified permission

        Args:
            sessionKey the session key to validate
            permissionName the permission name to look for
        Returns:
            True if permission granted, otherwise false
        """
        permissions = ClientPermissionValidator.get_session_permissions(sessionKey)
        print(f'Permission {permissionName} in permissions? {permissionName in permissions}')
        return permissionName in permissions


    @staticmethod
    def get_session_permissions(sessionKey: str) -> list[int]:
        """
        Retrieve the list of permissions from the session api

        Args:
            sessionKey is the current jwt key for the session
        Returns:
            List of available values
        """
        endpoint = os.getenv('SESSIONS_ENDPOINT')
        
        try:
            response = requests.get(f'{endpoint}/sessions/me', headers={'x-api-key':sessionKey})
            print(json.loads(response.text)['data'])
            print(json.loads(response.text)['data']['sessionData'])
            print(json.loads(response.text)['data']['sessionData']['SessionPermissionList'])
            permissions = json.loads(response.text)['data']['sessionData']['SessionPermissionList']
            return permissions
        except requests.exceptions.RequestException as e:

            print(f'Error in get_session_permissions: {e}')
        return []

