"""
Module for the permission interface
"""
from abc import ABC, abstractmethod

class PermissionInterface(ABC):
    """Interface to decouple the permission validation"""


    @abstractmethod
    def can_delete_guest(self):
        """Get all of the guests and staff from the database
        
            Returns:
                List of all users
            Raises:
                NotImplementedError if the method is not implemented
        """
        raise NotImplementedError


    @abstractmethod
    def can_delete_staff(self):
        """Get all staff from the database
        
            Returns:
                List of all staff
            Raises:
                NotImplementedError if the method is not implemented
        """
        raise NotImplementedError


    @abstractmethod
    def can_update_guest(self):
        """Get all of the guests from the database
        
            Returns:
                List of all guests
            Raises:
                NotImplementedError if the method is not implemented
        """
        raise NotImplementedError



