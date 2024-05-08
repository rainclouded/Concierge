import re
import app.Configs as cfg
from secrets import randbelow
from app.DatabaseController import DatabaseController
from app.UserObject import UserObject as User


class AuthenticationManager:
    """
    Class for handling the authentication of users as well as the creation
    of new users.
    """

    #Regex to find alphabetic characters
    GET_ALPHAPETIC_REGEX = '[a-zA-Z]'
    #Regex to find numeric characters
    GET_NUMERIC_REGEX = '[0-9]'

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


    """Computes the hash of the seeded password

        Args:
            id: seed of the password
            password: string of the password to process

        Returns:
            Hex hash of the password
    """
    def get_hash(self, id:int, password:str)->str:
        return cfg.PASSWORD_HASH_FUNCTION(
                f"{password}+{id}".encode()
            ).hexdigest()


    def validate_staff_login(self, username:str, password:str)->bool:
        """Validate the credentials of a user

            Args:
                username: String of the username to validate
                password: string of the password to validate

            Returns:
                If the user was successfully validated
        """
        user = \
            list(filter(lambda x : x.username == username, self.db.get_staff()))
        if len(user) != 1:
            return False
        return self.check_hash(user.pop(), password)


    def validate_staff_password(self, password:str)->bool:
        """Validate if a password can be used

            Args:
                password: string of the password to validate

            Returns:
                If the password meets all criteria
        """

        return not (
            len(password) < cfg.PASSWORD_MINIMAL_LENGTH
            or (
                cfg.PASSWORD_MUST_CONTAIN_LETTER 
                and not re.findall(self.GET_ALPHAPETIC_REGEX, password)
            )
            or (
                cfg.PASSWORD_MUST_CONTAIN_NUMBER 
                and not re.findall(self.GET_NUMERIC_REGEX, password)
            )
        )
    

    def validate_staff_username(self, password:str)->bool:
        """Validate if a password can be used

            Args:
                password: string of the password to validate

            Returns:
                If the password meets all criteria
        """
        return not (
            len(password) < cfg.USERNAME_MINIMAL_LENGTH
            or (
                cfg.USERNAME_MUST_CONTAIN_LETTER 
                and not re.findall(self.GET_ALPHAPETIC_REGEX, password)
            )
            or (
                cfg.USERNAME_MUST_CONTAIN_NUMBER 
                and not re.findall(self.GET_NUMERIC_REGEX, password)
            )
        )
        


    def create_new_guest(self, new_guest:User)->User:
        """Adds a new guest user

        Args:
            new_guest: The user to add

        Returns:
            The newly created user
        """
        if self.validate_new_guest(new_guest):
            new_guest.password = randbelow(cfg.MAX_GUEST_PASSWORD)
            new_guest.hash = self.get_hash(new_guest.username,new_guest.password)
            return self.db.create_guest(new_guest)
        return None
    

    def delete_user(self, user:User)->bool:
        """Remove a user from the database

            Args:
                user: The user to remove

            Returns:
                If the user was successfully deleted
        """
        return self.db.delete_user(user)


    def create_new_staff(self, new_user:User)->User:
        """Attempt to create a new staff in the database

            Args:
                new_user: dictionary containing at least a username and password

            Returns:
                If the staff was successfully created
        """
        if self.validate_new_staff(new_user):
            new_user.id = self.db.getLargestId()+1
            new_user.hash = self.get_hash(new_user.id, new_user.password)
            return self.db.create_staff(new_user)
        return None


    def validate_new_staff(self, new_user:User)->bool:
        """Validate if the credentials can be used

            Args:
                new_user: dictionary containing at least a username and password

            Returns:
                If the password meets all criteria
        """
        usernames = list(filter(lambda x: x.username, self.db.get_staff()))
        return (
            self.validate_staff_password(new_user.password)
            and self.validate_staff_username(new_user.username)
            and new_user.username not in usernames
            )
        
            
