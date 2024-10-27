"""
Module for the permission controller
"""
from app.permissions.PermissionInterface import PermissionInterface

class PermissionController():
    """
    The permission controller class manages
    whether the permissions provided allow for various
    actions
    """


    def __init__(self, permission_handler:PermissionInterface):
        self._permission_handler = permission_handler


    @property
    def permission_handler(self):
        return self._permission_handler


    @permission_handler.setter
    def permission_handler(self, new_permissions):
        self._permission_handler = new_permissions


    def can_delete_guest(self, token:str, public_key:str=None)->bool:
        """Verifies if the token permits guest deletion

            Args:            
                token is the jwt Token to be verified
                public_key is the public key to decode the token
            Returns:
                True if deletion is permitted
                False otherwise
        """
        public_key = self.get_public_key() if not public_key else public_key
        try:
            return self._permission_handler.can_delete_guest(token, public_key)
        except Exception as e:
            raise e


    def can_delete_staff(self, token:str, public_key:str=None)->bool:
        """Verifies if the token permits staff deletion

            Args:            
                token is the jwt Token to be verified
                public_key is the public key to decode the token
            Returns:
                True if deletion is permitted
                False otherwise
        """
        public_key = self.get_public_key() if not public_key else public_key
        try:
            return self._permission_handler.can_delete_staff(token, public_key)
        except Exception as e:
            raise e

    def can_update_guest(self, token:str, public_key:str=None)->bool:
        """Verifies if the token permits guest update

            Args:            
                token is the jwt Token to be verified
                public_key is the public key to decode the token
            Returns:
                True if update is permitted
                False otherwise
        """
        public_key = self.get_public_key() if not public_key else public_key
        try:
            return self._permission_handler.can_update_guest(token, public_key)
        except Exception as e:
            raise e

    def can_update_staff(self, token:str, public_key:str=None)->bool:
        """Verifies if the token permits staff update
        
            Args:            
                token is the jwt Token to be verified
                public_key is the public key to decode the token
            Returns:
                True if update is permitted
                False otherwise
        """        
        public_key = self.get_public_key() if not public_key else public_key
        try:
            return self._permission_handler.can_update_staff(token, public_key)
        except Exception as e:
            raise e

    def get_public_key(self)->str:
        """Get the current active public key

            Returns:
                string of public key
        """
        return self.permission_handler.get_public_key()
