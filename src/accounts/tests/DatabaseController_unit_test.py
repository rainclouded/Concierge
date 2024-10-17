import unittest
from app import app
from app.database.DatabaseController import DatabaseController
from app.database.Mockdata import Mockdata
from app.dto.UserObject import UserObject as User

class TestDatabaseController(unittest.TestCase):
    TEST_DATA = [
        {
            'username' : 'tesrt1',
            'id' : '1',
            'hash' : '',
            'type' : 'staff'
        },
        {
            'username' : 'trest2',
            'id' : '2',
            'hash' : '',
            'type' : 'staff'
        },
        {
            'username' : 'tesrt3',
            'id' : '3',
            'hash' : '',
            'type' : 'staff'
        },
        {
            'username' : '5',
            'id' : '',
            'hash' : '',
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

    def setUp(self):
        self.data = Mockdata()
        self.data.users = [*self.TEST_DATA]#Deepcopy hack
        self.db = DatabaseController(self.data)
        self.maxDiff = None
        
    def test_get_users(self):
        test_users = [User(**user) for user in self.TEST_DATA]
        retrieved_users = self.db.get_users()

        self.assertTrue(all([type(user) == User for user in retrieved_users]))
        self.assertCountEqual(retrieved_users, test_users)

    def test_get_staff(self):
        test_staff = [User(**staff) for staff in self.TEST_DATA[0:3]]
        retrieved_staff = self.db.get_staff()

        self.assertTrue(all([type(staff) == User for staff in retrieved_staff]))
        self.assertCountEqual(retrieved_staff, test_staff)
        

    def test_get_guests(self):
        test_guests = [User(**guest) for guest in self.TEST_DATA[3:6]]
        retrieved_guests = self.db.get_guests()

        self.assertTrue(all([type(guest) == User for guest in retrieved_guests]))
        self.assertCountEqual(retrieved_guests, test_guests)


    def test_create_guest(self):
        new_user = User(**{
            'username' : 'new',
            'hash' : '',
            'id' : '',
            'type' : 'guest'
        })

        self.db.create_guest(new_user)
        guests = self.db.get_guests()
        newly_added_guest = list(filter(lambda x : x.username == 'new', guests))
        self.assertTrue(len(newly_added_guest) == 1)
        self.assertEqual(newly_added_guest.pop(), new_user)
        self.assertTrue(len(guests) == 4)
        
    def test_create_staff(self):
        new_user = User(**{
            'username' : 'new',
            'hash' : '',
            'id' : '',
            'type' : 'staff'
        })

        self.db.create_staff(new_user)
        staff = self.db.get_staff()
        newly_added_guest = list(filter(lambda x : x.username == 'new', staff))
        self.assertTrue(len(newly_added_guest) == 1)
        self.assertEqual(newly_added_guest.pop(), new_user)
        self.assertTrue(len(staff) == 4)
        pass

    def test_gest_largest_id(self):
        self.assertEqual(3, self.db.get_largest_id())

        new_users = [User(**{
            'username' : 'new',
            'hash' : '',
            'id' : '4',
            'type' : 'staff'
        }),
        User(**{
            'username' : 'new',
            'hash' : '',
            'id' : '-3',
            'type' : 'staff'
        })
        ]
        for staff in new_users:
            self.db.create_staff(staff)
        self.assertEqual(4, self.db.get_largest_id())


    def test_delete_user(self):
        to_delete = list(filter(lambda x : x.username == 'test3', self.db.get_staff())).pop()

        valid_staff = [User(**{
            'username' : 'trest1',
            'id' : '1r',
            'hash' : '',
            'type' : 'staff'
        }),
        User(**{
            'username' : 'test2',
            'id' : '2',
            'hash' : '',
            'type' : 'staff'
        })]

        self.db.delete_user(to_delete)
        resultant_staff = self.db.get_staff()

        self.assertCountEqual(resultant_staff, valid_staff)
        self.assertEqual(len(resultant_staff), 2)
        self.assertEqual(self.db.get_largest_id(), 2)
