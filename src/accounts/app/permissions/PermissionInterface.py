"""
Module for the permission interface
"""
from abc import ABC, abstractmethod

class PermissionInterface(ABC):
    """Interface to decouple the permission validation"""


    @abstractmethod
    def can_delete_guest(self, token:str, public_key:str):
        """Verify if the token permits guest deletion
        
            Returns:
                Boolean indicating permission
            Raises:
                NotImplementedError if the method is not implemented
        """
        raise NotImplementedError


    @abstractmethod
    def can_delete_staff(self, token:str, public_key:str):
        """Verify if the token permits staff deletion
        
            Returns:
                Boolean indicating permission
            Raises:
                NotImplementedError if the method is not implemented
        """
        raise NotImplementedError


    @abstractmethod
    def can_update_guest(self, token:str, public_key:str):
        """Verify if the token permits guest update
        
            Returns:
                Boolean indicating permission
            Raises:
                NotImplementedError if the method is not implemented
        """
        raise NotImplementedError


    @abstractmethod
    def can_update_staff(self, token:str, public_key:str):
        """Verify if the token permits staff update
        
            Returns:
                Boolean indicating permission
            Raises:
                NotImplementedError if the method is not implemented
        """
        raise NotImplementedError


    @abstractmethod
    def decode_token(self, token:str, public_key:str):
        """Decode a jwt token
        
            Returns:
                Decoded token value
            Raises:
                NotImplementedError if the method is not implemented
        """
        raise NotImplementedError


    @abstractmethod
    def get_public_key(self):
        """Get the valid public key
        
            Returns:
                Public key string
            Raises:
                NotImplementedError if the method is not implemented
        """
        raise NotImplementedError
    