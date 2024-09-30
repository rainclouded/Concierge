
import hashlib
import re
import Configs as cfg
from DatabaseConroller import DatabaseController as db
class AuthenticationManager:
    """
    Class for handling the authentication of users as well as the creation
    of new users.
    """


    def check_hash(self, user:dict, password:str):
        """Validates the hash of password to that of user

            Args:
                user: Dictionary value that contains a "hash" value and an "id"
                    among other values
                password: string of the password to verify

            Returns:
                If the hashed password matches the user's hash
        """
        return (
            cfg.PASSWORD_HASH_FUNCTION(
                f"{password}+{user["id"]}".encode()
            ).hexdigest() == user["hash"]
        )


    def validate_staff_login(self, username:str, password:str):
        """Validate the credentials of a user

            Args:
                username: String of the username to validate
                password: string of the password to validate

            Returns:
                If the user was successfully validated
        """
        user = \
            list(filter(lambda x : x["username"] == username, db.get_staff()))
        if len(user) != 1:
            return False
        return self.check_hash(user.pop(), password)


    def validate_staff_password(self, password:str):
        """Validate if a password can be used

            Args:
                password: string of the password to validate

            Returns:
                If the password meets all criteria
        """
        if len(password) < cfg.PASSWORD_MINIMAL_LENGTH:
            return False
        
        if (
            cfg.PASSWORD_MUST_CONTAIN_LETTER 
            and not re.findall('[a-zA-Z]', password)
        ):
            return False
        
        if (
            cfg.PASSWORD_MUST_CONTAIN_NUMBER 
            and not re.findall('[0-9]', password)
        ):
            return False
        
        return True
    
    def validate_staff_username(self, password:str):
        """Validate if a password can be used

            Args:
                password: string of the password to validate

            Returns:
                If the password meets all criteria
        """
        if len(password) < cfg.USERNAME_MINIMAL_LENGTH:
            return False
        
        if (
            cfg.USERNAME_MUST_CONTAIN_LETTER 
            and not re.findall('[a-zA-Z]', password)
        ):
            return False
        
        if (
            cfg.USERNAME_MUST_CONTAIN_NUMBER 
            and not re.findall('[0-9]', password)
        ):
            return False
        
        return True


    def create_new_user(self):
        pass


    def validate_newUser(self,new_user:dict):
        pass