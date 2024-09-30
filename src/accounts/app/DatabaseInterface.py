from abc import ABC, abstractmethod

class DatabaseInterface(ABC):
    """Interface to decouple the database from the server"""

    @classmethod
    @abstractmethod
    def get_all_users():
        """Get all of the guests and staff from the database
        
            Returns:
                List of all users
            Raises:
                NotImplementedError if the method is not implemented
        """
        raise NotImplementedError
    
    @classmethod
    @abstractmethod
    def get_all_staff():
        """Get all staff from the database
        
            Returns:
                List of all staff
            Raises:
                NotImplementedError if the method is not implemented
        """
        raise NotImplementedError
    
    @classmethod
    @abstractmethod
    def get_all_guests():
        """Get all of the guests from the database
        
            Returns:
                List of all guests
            Raises:
                NotImplementedError if the method is not implemented
        """
        raise NotImplementedError

    @classmethod
    @abstractmethod
    def add_staff():
        """Add a staff to the database
        
            Raises:
                NotImplementedError if the method is not implemented
        """
        raise NotImplementedError
        
    @classmethod
    @abstractmethod
    def add_guest():
        """Add a guest to the database
        
            Raises:
                NotImplementedError if the method is not implemented
        """
        raise NotImplementedError
    
    @classmethod
    @abstractmethod
    def get_largest_id():
        """Get the largest user id
        
            Raises:
                NotImplementedError if the method is not implemented
        """
        raise NotImplementedError