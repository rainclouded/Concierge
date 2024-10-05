import app.Configs as cfg
from app.database.DatabaseController import DatabaseController
from app.dto.UserObject import UserObject as User

class AuthenticationManager:
    """
    Class for handling the authentication of users and their passwords
    """

    def __init__(self, database):
        self.db = DatabaseController(database)


    def check_hash(self, user:User, password:str)->bool:
        """Validates the hash of password to that of user

            Args:
                user: Dictionary value that contains a hash value and an id
                    among other values
                password: string of the password to verify

            Returns:
                If the hashed password matches the user's hash
        """
        if user.type == cfg.GUEST_TYPE:
            return self.get_hash(user.id, password) == user.hash
        else:
            return self.get_hash(user.username, password) == user.hash


    def get_hash(self, id:int, password:str)->str:
        """Computes the hash of the seeded password

            Args:
                id: seed of the password
                password: string of the password to process

            Returns:
                Hex hash of the password
        """
        return cfg.PASSWORD_HASH_FUNCTION(
                f"{password}+{id}".encode()
            ).hexdigest()


    def authenticate_user_login(self, username:str, password:str)->bool:
        """Validate the credentials of a user

            Args:
                username: String of the username to validate
                password: string of the password to validate

            Returns:
                If the user was successfully validated
        """
        user = \
            list(filter(lambda x : x.username == username, self.db.get_users()))

        return len(user) == 1 and self.check_hash(user.pop(), password)