class MockPermissionValidator:
    def validate_session_key_for_permission_name(self, sessionKey: str, permissionName: str) -> bool:
        return True
    
    def get_session_permissions(self, sessionKey: str) -> list[int]:
        return []