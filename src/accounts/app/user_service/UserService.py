import app.Configs as cfg
from secrets import randbelow
from app.authentication.AuthenticationManager import AuthenticationManager
from app.validation.ValidationManager import ValidationManager
from app.database.DatabaseController import DatabaseController
from app.dto.UserObject import UserObject as User

class UserService():
    """
    Class for handling maintnance of user accounts
    """
    
    def __init__(self, database:DatabaseController):
        self.db = database
        self.auth = AuthenticationManager(database)
        self.validation = ValidationManager(database)

    def create_new_guest(self, new_guest:User)->User:
        """Adds a new guest user

        Args:
            new_guest: The user to add

        Returns:
            The newly created user
        """
        
        new_guest.password = randbelow(cfg.MAX_GUEST_PASSWORD)
        new_guest.hash = self.auth.get_hash(new_guest.username,new_guest.password)
        return self.db.create_guest(new_guest)
    

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
        if self.validation.validate_new_staff(new_user):
            new_user.id = self.db.getLargestId()+1
            new_user.hash = self.auth.get_hash(new_user.id, new_user.password)
            return self.db.create_staff(new_user)
        return None