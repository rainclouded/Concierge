"""
Module for core services
"""
import app.Configs as cfg
from app.authentication.AuthenticationManager import AuthenticationManager
from app.database.DatabaseController import DatabaseController
from app.user_service.UserService import UserService
from app.validation.ValidationManager import ValidationManager
from app.permissions.PermissionController import PermissionController

class Services():
    """
    This class manages singleton objects for dependency injection
    """
    _database = None
    _authentication = None
    _user_service = None
    _permissions = None
    _validation = None

    @classmethod
    def set_up(cls):
        """
        Set up the class' services
        """
        cls._database = DatabaseController(cfg.create_database())
        cls._authentication = AuthenticationManager(cls._database)
        cls._validation = ValidationManager(cls._database)
        cls._user_service = UserService(
            cls._database,
            cls._authentication,
            cls._validation
            )
        cls._permissions = PermissionController(cfg.create_permissions())
    
    @classmethod
    def inject_set_up(cls, database=None, authentication=None, validation=None, permissions=None):
        """
        Set up the class' services
        """
        cls._database = database
        cls._authentication = authentication
        cls._validation = validation
        cls._user_service = UserService(
            cls._database,
            cls._authentication,
            cls._validation
            )
        cls._permissions = permissions

    @classmethod
    def get_database(cls)->DatabaseController:
        """Get the database singleton
        
            Returns:
                DatabaseController instance
        """
        return cls._database

    @classmethod
    def get_authentication(cls)->AuthenticationManager:
        """Get the authentication singleton
        
            Returns:
                AuthenticationManager instance
        """
        return cls._authentication

    @classmethod
    def get_user_service(cls)->UserService:
        """Get the User Service singleton
        
            Returns:
                UserService instance
        """
        return cls._user_service

    @classmethod
    def get_permissions(cls)->PermissionController:
        """Get the permissions singleton
        
            Returns:
                PermissionController instance
        """
        return cls._permissions

    @classmethod
    def get_validation(cls)->ValidationManager:
        """Get the validation singleton
        
            Returns:
                ValidationManager instance
        """
        return cls._validation
