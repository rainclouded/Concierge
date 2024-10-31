import json
import os
import requests

class PermissionValidator:
    @staticmethod
    def validate_session_key_for_permission_name(sessionKey: str, permissionName: str) -> bool:
        permissions = PermissionValidator.get_session_permissions(sessionKey)
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