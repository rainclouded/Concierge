from app.dto.UserObject import UserObject as User
from app.database.DatabaseInterface import DatabaseInterface

class DatabaseController:
    """
    Database controller, should be able to be used
    with any concrete implementation of the DatabaseInterface
    This class queries the database.
    """


    def __init__(self, database:DatabaseInterface):
        self._database = database


    @property
    def database(self):
        return self._database


    @database.setter
    def database(self, new_database):
        self._database = new_database


    def get_users(self)->list[User]:
        """
            List all of the users in the database

            Returns:
                List of user objects
        """
        return [User(**user) for user in self.database.get_all_users()]


    def get_staff(self)->list[User]:
        """
            List all of the staff in the database

            Returns:
                List of user objects
        """
        return [User(**staff) for staff in self.database.get_all_staff()]


    def get_guests(self)->list[User]:
        """
            List all of the guests in the database

            Returns:
                List of guest objects
        """
        return [User(**guest) for guest in self.database.get_all_guests()]


    def create_guest(self, new_guest:User)->User:
        """
            Add a guest to the database

            Args:
                new_guest: The guest User to add

            Returns:
                The user if successfully added else None
        """
        if self.database.add_guest(new_guest.__dict__):
            return new_guest
        return None


    def create_staff(self, new_staff:User)->User:
        """
            Add a staff User to the database

            Args:
                new_staff: The staff User to add

            Returns:
                The user added or none
        """
        if self.database.add_staff(new_staff.__dict__):
            return new_staff
        return None


    def get_largest_id(self)->int:
        """ Get the largest id in the database of all the staff

            Returns:
                The largest id
        """
        return max(int(user.id) for user in self.get_staff())


    def delete_user(self, user:User)->User:
        """Delete a user from the database
        
            Args:
                user: The User to delete
            Returns:
                bool if the deletion was successful
        """
        return (
                User(**user)
                if ( user := self.database.delete_user(user.username))
                else None
            )
