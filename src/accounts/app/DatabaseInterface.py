from abc import ABC

class Databaseinterface(ABC):

    @classmethod
    @ABC.abstractmethod
    def get_all_users():
        raise NotImplementedError
    
    @classmethod
    @ABC.abstractmethod
    def get_all_staff():
        raise NotImplementedError
    
    @classmethod
    @ABC.abstractmethod
    def get_all_guests():
        raise NotImplementedError

    @classmethod
    @ABC.abstractmethod
    def add_staff():
        raise NotImplementedError
        
    @classmethod
    @ABC.abstractmethod
    def add_guest():
        raise NotImplementedError