from Mockdata import Mockdata
from UserObject import UserObject as User

class DatabaseController:
    def __init__(self):
        self.db = Mockdata()

    def get_users(self):
        return [User(user) for user in self.db.get_all_users()]

    def get_staff(self):
        return [User(staff) for staff in self.db.get_all_staff()]

    def get_guests(self):
        return [User(guest) for guest in self.db.get_all_guests()]
        
    def create_guest(self, new_guest:User):
        if self.db.add_guest(self._user_to_dict(new_guest)):
            return new_guest
        return None

    def create_staff(self, new_staff:User):
        if self.db.add_staff(self._user_to_dict(new_staff)):
            return new_staff
        return None

    def get_largest_id(self):
        return (list(filter(lambda x: x['id'], self.get_users())))
    
    def delete_user(self, user:User):
        self.db.delete_user(user.username)
        return True
    
    def _user_to_dict(self, user:User):
        if user.type == 'guest':
            return {
                'username' : user.username,
                'hash' : user.hash,
                'id' : '',
                'type' : user.type
            }
            
        if user.type == 'staff':
            return {
                'username' : user.username,
                'hash' : user.hash,
                'id' : user.id,
                'type' : user.type
            }
      
    #ToDo: Add support for connecting/committing and other db administration