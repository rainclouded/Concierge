from .server.Accounts import app
from .authentication.AuthenticationManager import AuthenticationManager
from .dto.UserObject import UserObject
from .database.DatabaseController import DatabaseController
from .Configs import DATABASE

__all__ = [app]

def create_database():
    return DatabaseController(DATABASE)