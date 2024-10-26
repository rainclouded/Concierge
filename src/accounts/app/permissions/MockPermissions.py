from app.permissions.PermissionInterface import PermissionInterface

class MockPermissions(PermissionInterface):

    def can_delete_guest(self):
        return True

    def can_delete_staff(self):
        return True

    def can_update_guest(self):
        return True
