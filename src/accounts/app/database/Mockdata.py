"""
Module for the mock database
"""
from app.database.DatabaseInterface import DatabaseInterface

class Mockdata(DatabaseInterface):
    """
    Class used to mock the database and
    provide mock data; just for testing
    """
    def __init__(self):
        #Set test data
        self._users = [
        {#password: testPassword1
            'username' : 'test1',
            'id' : '1',
            'hash' : '2ab34e3ec1de9d16996e303582da30758f88712c5196212d7e07914a859cc21b',
            'type' : 'staff'
        },
        {
            'username' : 'test2',
            'id' : '2',
            'hash' : '',
            'type' : 'staff'
        },
        {
            'username' : 'test3',
            'id' : '3',
            'hash' : '',
            'type' : 'staff'
        },
        {#password: 44444444
            'username' : '5',
            'id' : '',
            'hash' : 'fb86d0f8ce24539b550e58c8398343cea6c09836614deefa06db7a460822d78c',
            'type' : 'guest'
        },
        {
            'username' : '6',
            'id' : '',
            'hash' : '',
            'type' : 'guest'
        },
        {
            'username' : '7',
            'id' : '',
            'hash' : '',
            'type' : 'guest'
        }
    ]


    @property
    def users(self):
        return self._users


    @users.setter
    def users(self, new_list):
        self._users = new_list


    def get_all_users(self)->list[dict]:
        """Get the entire user base
        
            Returns:
                List of dictionaries corresponding to User objects
        """
        return self.users


    def get_all_staff(self)->list[dict]:
        """Get the entire user base of staff users
        
            Returns:
                List of dictionaries corresponding to User objects
        """
        return list(filter(lambda x : x['type'] == 'staff', self.users))


    def get_all_guests(self)->list[dict]:
        """Get the entire user base of guest users
        
            Returns:
                List of dictionaries corresponding to User objects
        """
        return list(filter(lambda x : x['type'] == 'guest', self.users))


    def add_staff(self, new_staff:dict)->True:
        """Add a staff user to the userbase
        
            Returns:
                True
        """
        self.users.append(new_staff)
        return True


    def add_guest(self, new_guest:dict)->True:
        """Add a guest user to the userbase
        
            Returns:
                True
        """
        self.users.append(new_guest)
        return True


    def delete_user(self, username)->dict:
        """Delete user from the userbase
        
            Returns:
                Dictionary corresponding to the user deleted
        """
        for idx, user in enumerate(self.users):
            if user['username'] == username:
                return self.users.pop(idx)
        return None

    def update_user(self, update_user:dict)->bool:
        """Update user from in the userbase

            Returns:
                True if the user was successfully updated
                False otherwise
        """
        for idx, user in enumerate(self.users):
            if user['username'] == update_user['username']:
                self.users[idx] = update_user
                return True
        return False
    