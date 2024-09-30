from Mockdata import Mockdata

class DatabaseController:
    def __init__(self):
        self.db = Mockdata()

    def get_users(self):
        return self.db.get_all_users()

    def get_staff(self):
        return self.db.get_all_staff()

    def get_guests(self):
        return self.db.get_all_guests()
        
    def create_guest(self, new_guest:dict):
        pass

    def create_staff(self, new_staff:dict):
        pass
    #ToDo: Add support for connecting/committing and other db administration