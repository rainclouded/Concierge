"""
Module for account VerificationManager
"""
import re
import app.Configs as cfg
from app.database.DatabaseController import DatabaseController
from app.dto.UserObject import UserObject as User


class ValidationManager():
    """
    Class for handling the validation of user credentials
    """

    #Regex to find alphabetic characters
    GET_ALPHAPETIC_REGEX = '[a-zA-Z]'
    #Regex to find numeric characters
    GET_NUMERIC_REGEX = '[0-9]'

    def __init__(self, database:DatabaseController):
        self.db = database


    def validate_staff_password(self, password:str)->bool:
        """Validate if a password can be used

            Args:
                password: string of the password to validate

            Returns:
                If the password meets all criteria
        """

        return password is not  None and not (
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


    def validate_staff_username(self, username:str)->bool:
        """Validate if a password can be used

            Args:
                password: string of the password to validate

            Returns:
                If the password meets all criteria
        """
        return username is not None and not (
            len(username) < cfg.USERNAME_MINIMAL_LENGTH
            or (
                cfg.USERNAME_MUST_CONTAIN_LETTER
                and not re.findall(self.GET_ALPHAPETIC_REGEX, username)
            )
            or (
                cfg.USERNAME_MUST_CONTAIN_NUMBER
                and not re.findall(self.GET_NUMERIC_REGEX, username)
            )
        )

    def validate_guest_username(self, username:str)->bool:
        """Validate if a password can be used

            Args:
                password: string of the password to validate

            Returns:
                If the password meets all criteria
        """
        return username is not None and username.isdigit()

    def validate_new_staff(self, new_user:User, password:str)->bool:
        """Validate if the credentials can be used

            Args:
                new_user: dictionary containing at least a username and password

            Returns:
                If the password meets all criteria
        """
        usernames = [user.username for user in self.db.get_staff()]
        return (
            new_user is not None
            and self.validate_staff_password(password)
            and self.validate_staff_username(new_user.username)
            and new_user.username not in usernames
            )


    def validate_new_guest(self, new_user:User)->bool:
        """Validate if the credentials can be used

            Args:
                new_user: dictionary containing at least a username and password

            Returns:
                If the password meets all criteria
        """
        usernames = [user.username for user in self.db.get_guests()]
        print(usernames)
        print(f"""testing user thing {(
            new_user is not None
            and self.validate_guest_username(new_user.username)
            and new_user.username not in usernames
            )}""")
        return (
            new_user is not None
            and self.validate_guest_username(new_user.username)
            and new_user.username not in usernames
            )