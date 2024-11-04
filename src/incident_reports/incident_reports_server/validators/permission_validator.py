import os
import requests

class PermissionValidator:
    def validate_session_key_for_permission_name(self, sessionKey: str, permissionName: str) -> bool:
        if not sessionKey:
            return False
        permissions = self.get_session_permissions(sessionKey)
        return permissionName in permissions
    
    def get_session_permissions(self, sessionKey: str) -> list[int]:
        endpoint = os.getenv('SESSIONS_ENDPOINT')
        try:
            response: requests.Response = requests.get(f'{endpoint}/sessions/me', headers={"X-Api-Key": sessionKey})
            permissions = response.json().get('data', {}).get('sessionData', {}).get('SessionPermissionList', [])
            return permissions
        except requests.exceptions.RequestException as e:
            print(e)
        return []