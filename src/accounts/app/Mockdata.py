from . import DatabaseInterface

class Mockdata(DatabaseInterface):
    """
    Class used to mock the database and
    provide mock data
    """

    users = [
        {
            'username' : 'test1',
            'identifier' : '1',
            'hash' : '',
            'type' : 'staff'
        }
    ]

    def get_all_users(self):
        return self.users
    

    def get_all_staff(self):
        return list(filter(lambda x : x['type'] == 'staff', self.users))


    def get_all_guests(self):
        return list(filter(lambda x : x['type'] == 'guest', self.users))


    def add_staff(self):
        pass

    def add_guest(self):
        pass