"""
Module for UserService account service
"""
from secrets import randbelow
import app.Configs as cfg
from app.dto.UserObject import UserObject as User

class UserService():
    """
    Class for handling maintnance of user accounts
    """

    def __init__(self, database, authentication, validation):
        self.db = database
        self.auth = authentication
        self.validation = validation

    def get_user_type(self, username:str)->str:
        """Find the 'type' of a user
        """
        return next(
            (
                user.type for user in self.db.get_users()
                if user.username == username
            ),
            None
        )


    def create_new_guest(self, new_guest:User)->tuple[User, str]:
        """Adds a new guest user

        Args:
            new_guest: The user to add

        Returns:
            The newly created user
        """
        if self.validation.validate_new_guest(new_guest):
            new_guest.password = randbelow(cfg.MAX_GUEST_PASSWORD)
            new_guest.hash = \
                self.auth.get_hash(new_guest.username,new_guest.password)
            return (self.db.create_guest(new_guest), f'{new_guest.password}')
        return None


    def delete_user(self, username:str)->bool:
        """Remove a user from the database

            Args:
                user: The user to remove

            Returns:
                If the user was successfully deleted
        """
        return self.db.delete_user(username)


    def create_new_staff(self, new_user:User, password:str)->User:
        """Attempt to create a new staff in the database

            Args:
                new_user: dictionary containing at least a username and password

            Returns:
                If the staff was successfully created
        """
        if self.validation.validate_new_staff(new_user, password):
            new_user.id = self.db.get_largest_id()+1
            new_user.hash = self.auth.get_hash(new_user.id, password)
            return self.db.create_staff(new_user)
        return None

    def update_user(self, user:User)->tuple[User, str]:
        """
            Update the account associated with a username
            Args
        """
        return (
            self.create_new_guest(user)
            if self.delete_user(user.username)
            else (None, None)
        )
    
    def get_users(self, user_type=None):
        """
        Get all the users of a specific type or
        just all the users

        """
        users = None
        if user_type == cfg.GUEST_TYPE:
            users = self.db.get_guests()
        elif user_type == cfg.STAFF_TYPE:
            users = self.db.get_staff()
        else:
            users = self.db.get_users()
        return(users)
