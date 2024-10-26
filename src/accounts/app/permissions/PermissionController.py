from app.permissions.PermissionInterface import PermissionInterface

class PermissionController():

    def __init__(self, permission_handler:PermissionInterface):
        self._permission_handler = permission_handler

    @property
    def permission_handler(self):
        return self._permission_handler
    
    @permission_handler.setter
    def permission_handler(self, new_permissions):
        self._permission_handler = new_permissions

    def can_delete_guest(self, token:str)->bool:
        return self._permission_handler.can_delete_guest(token)

    def can_delete_staff(self, token:str)->bool:
        return self._permission_handler.can_delete_staff(token)

    def can_update_guest(self, token:str)->bool:
        return self._permission_handler.can_update_guest(token)