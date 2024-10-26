"""
Module for the permission interface
"""
from abc import ABC, abstractmethod

class PermissionInterface(ABC):
    """Interface to decouple the permission validation"""


    @abstractmethod
    def can_delete_user(self, token:str):
        """Verify if the token permits user deletion
        
            Returns:
                Boolean indicationg permission
            Raises:
                NotImplementedError if the method is not implemented
        """
        raise NotImplementedError


    @abstractmethod
    def can_update_user(self, token:str):
        """Verify if the token permits user update
        
            Returns:
                Boolean indicationg permission
            Raises:
                NotImplementedError if the method is not implemented
        """
        raise NotImplementedError

    def decode_token(self, token:str, public_key:str, algorithm:str):
        """Decode a jwt token
        
            Returns:
                Decoded token value
            Raises:
                NotImplementedError if the method is not implemented
        """
        raise NotImplementedError
