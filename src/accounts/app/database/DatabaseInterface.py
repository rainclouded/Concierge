"""
Module for the database interface
"""
from abc import ABC, abstractmethod

class DatabaseInterface(ABC):
    """Interface to decouple the database from the server"""


    @abstractmethod
    def get_all_users(self):
        """Get all of the guests and staff from the database
        
            Returns:
                List of all users
            Raises:
                NotImplementedError if the method is not implemented
        """
        raise NotImplementedError


    @abstractmethod
    def get_all_staff(self):
        """Get all staff from the database
        
            Returns:
                List of all staff
            Raises:
                NotImplementedError if the method is not implemented
        """
        raise NotImplementedError


    @abstractmethod
    def get_all_guests(self):
        """Get all of the guests from the database
        
            Returns:
                List of all guests
            Raises:
                NotImplementedError if the method is not implemented
        """
        raise NotImplementedError


    @abstractmethod
    def add_staff(self, new_staff:dict):
        """Add a staff to the database
        
            Raises:
                NotImplementedError if the method is not implemented
        """
        raise NotImplementedError


    @abstractmethod
    def add_guest(self, new_guest:dict):
        """Add a guest to the database
        
            Raises:
                NotImplementedError if the method is not implemented
        """
        raise NotImplementedError


    @abstractmethod
    def update_user(self, update_user:dict):
        """Update the account info od the user whos username mathches
            update_user['username']. If the user does not exist, add them
        
            Raises:
                NotImplementedError if the method is not implemented
        """
        raise NotImplementedError


    @abstractmethod
    def delete_user(self, username:str):
        """
            Delete the user with the specified username
        """
        raise NotImplementedError
