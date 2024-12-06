import os
import requests

class PermissionValidator:
    
    # Validates if the provided sessionKey has the specified permissionName.
    # Args:
    #     sessionKey (str): The session key to validate.
    #     permissionName (str): The permission name to check within the session's permissions.
    # Returns:
    #     bool: True if the sessionKey has the specified permission, False otherwise.
    def validate_session_key_for_permission_name(self, sessionKey: str, permissionName: str) -> bool:
        # If no sessionKey is provided, return False (invalid session)
        if not sessionKey:
            return False
        
        # Retrieve the list of permissions for the given sessionKey
        permissions = self.get_session_permissions(sessionKey)
        
        # Return True if the permissionName is in the permissions list, otherwise return False
        return permissionName in permissions
    
    # Retrieves the list of permissions associated with a sessionKey.
    # Args:
    #     sessionKey (str): The session key whose permissions are to be fetched.
    # Returns:
    #     list[int]: A list of permission IDs for the given sessionKey.
    def get_session_permissions(self, sessionKey: str) -> list[int]:
        # Retrieve the endpoint URL for the sessions service from environment variables
        endpoint = os.getenv('SESSIONS_ENDPOINT')
        
        try:
            # Make a GET request to the sessions service using the provided sessionKey
            response: requests.Response = requests.get(f'{endpoint}/sessions/me', headers={"X-API-Key": sessionKey})
            
            # Extract the list of permissions from the response JSON
            permissions = response.json().get('data', {}).get('sessionData', {}).get('SessionPermissionList', [])
            
            # Return the list of permissions
            return permissions
        except requests.exceptions.RequestException as e:
            # Catch any exceptions that occur during the request and print the error
            print(e)
        
        # Return an empty list in case of an error or failed request
        return []
