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
        public_key = self.get_public_key() if not public_key else public_key
        try:
            return self._permission_handler.can_delete_guest(token, public_key)
        except Exception as e:
            raise e


    def can_delete_staff(self, token:str, public_key:str=None)->bool:
        public_key = self.get_public_key() if not public_key else public_key
        try:
            return self._permission_handler.can_delete_staff(token, public_key)
        except Exception as e:
            raise e

    def can_update_guest(self, token:str, public_key:str=None)->bool:
        public_key = self.get_public_key() if not public_key else public_key
        try:
            return self._permission_handler.can_update_guest(token, public_key)
        except Exception as e:
            raise e

    def can_update_staff(self, token:str, public_key:str=None)->bool:
        public_key = self.get_public_key() if not public_key else public_key
        try:
            return self._permission_handler.can_update_staff(token, public_key)
        except Exception as e:
            raise e

    def get_public_key(self):
        return self.permission_handler.get_public_key()