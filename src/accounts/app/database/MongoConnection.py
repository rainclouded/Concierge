"""
Module for the MongoDB connection
"""
from os import getenv
from pymongo import MongoClient
from app.database.DatabaseInterface import DatabaseInterface
import app.Configs as cfg

class MongoConnection(DatabaseInterface):
    """
    Implementation of DatabaseInterface which
    facilitates the connection to the MongoDB database
    """
    def __init__(self):
        self.uri = getenv('DB_URI')
        self.client = MongoClient(self.uri)
        self.database = self.client['accounts']
        if 'accounts' not in self.database.list_collection_names():
            self.database.create_collection('accounts')
            self.collection = self.database['accounts']
            self.collection.create_index([("username", 1)])
            self.add_guest(
                {
                    'id':0, 
                    'username':'11111', 
                    'hash':cfg.PASSWORD_HASH_FUNCTION(
                        "11111password".encode()
                        ).hexdigest(), 
                    'type':'guest'
                }
            )
            self.add_staff(
                {
                    'id':1, 
                    'username':'admin', 
                    'hash':cfg.PASSWORD_HASH_FUNCTION(
                        "1admin".encode()
                        ).hexdigest(), 
                    'type':'staff'
                }
            )
        else:
            self.collection = self.database['accounts']


    def get_all_users(self)->list[dict]:
        """Get the entire user base
        
            Returns:
                List of dictionaries corresponding to User objects
        """
        return self.collection.find({})


    def get_all_staff(self)->list[dict]:
        """Get the entire user base of staff users
        
            Returns:
                List of dictionaries corresponding to User objects
        """
        return self.collection.find({"type":"staff"})


    def get_all_guests(self)->list[dict]:
        """Get the entire user base of guest users
        
            Returns:
                List of dictionaries corresponding to User objects
        """
        return self.collection.find({"type":"guest"})


    def add_staff(self, new_staff:dict)->bool:
        """Add a staff user to the userbase
        
            Returns:
                True
        """
        try:
            self.collection.insert_one(new_staff)
        except Exception as e: raise e

        return True


    def add_guest(self, new_guest:dict)->bool:
        """Add a guest user to the userbase
        
            Returns:
                True
        """
        try:
            self.collection.insert_one(new_guest)
        except Exception as e: raise e

        return True


    def update_user(self, update_user)->bool:
        """Update user from in the userbase

            Returns:
                True if the user was successfully updated
                False otherwise
        """
        num_results = 0
        try:
            num_results = self.collection.replace_one(
                {'username' : update_user['username']},
                update_user,
                upsert = False #Do not error if the user does not exist
            )
        except Exception as e: raise e

        return bool(num_results)


    def delete_user(self, username:str)->dict:
        """Delete user from the userbase
        
            Returns:
                Dictionary corresponding to the user deleted
        """
        deleted_user = None
        try:
            deleted_user = self.collection.find_one_and_delete({'username':username})
        except Exception as e: raise e

        return deleted_user
    