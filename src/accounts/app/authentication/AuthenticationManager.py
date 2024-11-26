"""
    Module for handling the authentication of users and their passwords
"""
import app.Configs as cfg
from app.dto.UserObject import UserObject as User

class AuthenticationManager:
    """
    Class for handling the authentication of users and their passwords
    """

    def __init__(self, database):
        self.db = database


    def check_hash(self, user:User, password:str)->bool:
        """Validates the hash of password to that of user

            Args:
                user: Dictionary value that contains a hash value and an id
                    among other values
                password: string of the password to verify

            Returns:
                If the hashed password matches the user's hash
        """
        if user.type == cfg.STAFF_TYPE:
            return self.get_hash(user.id, password) == user.hash
        else:
            return self.get_hash(user.username, password) == user.hash


    def get_hash(self, user_id:int, password:str)->str:
        """Computes the hash of the seeded password

            Args:
                id: seed of the password
                password: string of the password to process

            Returns:
                Hex hash of the password
        """
        return cfg.PASSWORD_HASH_FUNCTION(
                f"{user_id}{password}".encode()
            ).hexdigest()


    def authenticate_user_login(self, username:str, password:str)->User:
        """Validate the credentials of a user

            Args:
                username: String of the username to validate
                password: string of the password to validate

            Returns:
                The User obj on success, None on failure
        """
        user = \
            list(filter(lambda x : x.username == username, self.db.get_users()))
        if len(user) == 1 and self.check_hash(user[0], password):
            return user[0]
        return None
