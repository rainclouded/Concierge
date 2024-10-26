import unittest
import os
from unittest.mock import patch, MagicMock
from app.database.MongoConnection import MongoConnection

class TestMongoConnection(unittest.TestCase):

    @patch('app.database.MongoConnection.MongoClient')
    def setUp(self, ClientMock):
        os.environ['DB_URI'] = 'test'
        self.collection = MagicMock()
        self.database = MagicMock()
        ClientMock.return_value.__getitem__.return_value = self.database
        self.database.__getitem__.return_value = self.collection
        self.ClientMock = ClientMock
        self.database.list_collection_names.return_value = ['accounts']
        self.Connection = MongoConnection()

    def test_creation(self):
       self.ClientMock.assert_called_once_with('test')

    def test_get_all_users(self):
        self.Connection.get_all_users()
        self.collection.find.assert_called_once_with({})

    def test_get_all_staff(self):
        self.Connection.get_all_staff()
        self.collection.find.assert_called_once_with({"type":"staff"})

    def test_get_all_guests(self):
        self.Connection.get_all_guests()
        self.collection.find.assert_called_once_with({"type":"guest"})

    def test_add_staff(self):
        staff_dict = {'test':'value'}
        self.collection.insert_one.return_value = True
        self.Connection.add_staff(staff_dict)
        self.collection.insert_one.assert_called_once_with(staff_dict)

    def test_add_guest(self):
        guest_dict = {'test':'value'}
        self.collection.insert_one.return_value = True
        self.Connection.add_guest(guest_dict)
        self.collection.insert_one.assert_called_once_with(guest_dict)

    def test_update_user(self):
        staff_dict = {'username':'value'}
        self.collection.replace_one.return_value = True
        self.Connection.update_user(staff_dict)
        self.collection.replace_one.assert_called_once_with(
            staff_dict, 
            staff_dict, 
            upsert=False
            )

    def test_delete_user(self):
        user_dict = {'username':'value'}
        self.collection.find_one_and_delete.return_value = True
        self.Connection.delete_user('value')
        self.collection.find_one_and_delete.assert_called_once_with(user_dict)
