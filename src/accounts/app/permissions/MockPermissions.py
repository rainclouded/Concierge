import jwt
from app.permissions.PermissionInterface import PermissionInterface

class MockPermissions(PermissionInterface):

    def can_delete_user(self, token:str, type:str)->bool:
        return True
    

    def can_update_user(self, token: str, type:str)->bool:
        return True


    def decode_token(self, token:str, public_key:str,\
                     algorithm:str='ES256'):
        return jwt.decode(token, public_key, [algorithm])
    