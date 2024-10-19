import unittest
from app import app
from app.database.DatabaseController import DatabaseController
from app.authentication.AuthenticationManager import AuthenticationManager
from app.database.Mockdata import Mockdata
from app.dto.UserObject import UserObject as User

class TestAuthenticationManager(unittest.TestCase):
    TEST_DATA = [
        {#password: testPassword1
            'username' : 'test1',
            'id' : '1',
            'hash' : '2ab34e3ec1de9d16996e303582da30758f88712c5196212d7e07914a859cc21b',
            'type' : 'staff'
        },
        {#password: testPassword2
            'username' : 'test2',
            'id' : '2',
            'hash' : '9a71d63826555499d852069307b764a30b7fabdcd693a357af3fb6c4c6b3cf1b',
            'type' : 'staff'
        },
        {#password: testPassword3
            'username' : 'test3',
            'id' : '3',
            'hash' : '1951eab9b856c9b3539b616a64608f01ea56d87a5de680946d9740289f8a10cb',
            'type' : 'staff'
        },
        {#password: testPassword3
            'username' : 'test3',
            'id' : '3',
            'hash' : '1951eab9b856c9b3539b616a64608f01ea56d87a5de680946d9740289f8a10cb',
            'type' : 'staff'
        },
        {#password: 55555555
            'username' : '5',
            'id' : '',
            'hash' : '2b218a3de6e9c348c3c482caee9ed793b7963c54d3bbe757a8b1ba7f64cdde0a',
            'type' : 'guest'
        },
        {#password: 66666666
            'username' : '6',
            'id' : '',
            'hash' : '185ce8fc660c5607c09afb444d81f918300a1fc7737d780e4a6c0ed5871c6dd6',
            'type' : 'guest'
        },
        {#password: 77777777
            'username' : '7',
            'id' : '',
            'hash' : 'd43403a2c3dae4e4332bf99111e6e066ecda233354d5fa44484dff058e483bb8',
            'type' : 'guest'
        },
        {#password: 77777777
            'username' : '7',
            'id' : '',
            'hash' : 'd43403a2c3dae4e4332bf99111e6e066ecda233354d5fa44484dff058e483bb8',
            'type' : 'guest'
        }
    ]

    def setUp(self):
        self.database = Mockdata()
        self.database.users = self.TEST_DATA
        self.database_controller = DatabaseController(self.database)
        self.am = AuthenticationManager(self.database_controller)

    def test_get_hash(self):
        test_values = [
            {
                'id' : 'test1',
                'password' : 'pass1',
                'hash' : '8c5941135f8132ee54a465f7f8fbe4eeff600ec008ec1810c1bc36854ab4bb54'
            },
            {
                'id' : '1',
                'password' : '111',
                'hash' : '0ffe1abd1a08215353c233d6e009613e95eec4253832a761af28ff37ac5a150c'
            },
            {
                'id' : 'testUserName334',
                'password' : 'testPassword334',
                'hash' : '1ef8982cfe0e745edaa4a3774e111f69df79f8b567446ca1fa1eaa797a6cd412'
            }
        ]
        for test in test_values:
            self.assertEqual(self.am.get_hash(test['id'], test['password']), test['hash'])


    def test_check_hash(self):

        test_values = [
            (
                User(
                    **{
                        'username' : 'TestUser1',
                        'hash' : '2ab34e3ec1de9d16996e303582da30758f88712c5196212d7e07914a859cc21b',
                        'id' : '1',
                        'type' : 'staff'
                    }),
                'testPassword1'
            ),
            (
                User(
                    **{
                        'username' : 'TestUser2',
                        'hash' : '9a71d63826555499d852069307b764a30b7fabdcd693a357af3fb6c4c6b3cf1b',
                        'id' : '2',
                        'type' : 'staff'
                    }),
                'testPassword2'
            ),
            (
                User(
                    **{
                        'username' : '217',
                        'hash' : '33c51863b34e6bb6ad0559fa087587353899b32bd1f9e241ccf9b72369b40e9a',
                        'id' : '',
                        'type' : 'guest'
                    }),
                '99999999'
            ),
            (
                User(
                    **{
                        'username' : '237',
                        'hash' : '8b7d070eefd77032affb9195e043b8d8021259b0b58f93ae020112fda5016de1',
                        'id' : '',
                        'type' : 'guest'
                    }),
                '88888888'
            )
        ]

        false_values = [
            (
                User(
                    **{#rightPassword1
                        'username' : 'BadData1',
                        'hash' : '92bcaf88de09ae720a24c8806a214812af73fa47bea0f63cdba5d8e23ee586e3',
                        'id' : '99',
                        'type' : 'staff'
                    }),
                'wrongpassword1'
            ),
            (
                User(
                    **{
                        'username' : '002',
                        'hash' : '1de2254e1e459fde743a2b829b2f906038f9cb13f98830a4d8d802d06a2c8afc',#rightPassword2
                        'id' : '',
                        'type' : 'guest'
                    }),
                'wrongpassword2'
            )
        ]

        for user, password in test_values:
            self.assertTrue(self.am.check_hash(user, password))

        for user, password in false_values:
            self.assertFalse(self.am.check_hash(user, password))


    def test_authenticate_user_login(self):
        valid_credentials = [
            ('test1','testPassword1'),
            ('test2', 'testPassword2'),
            ('5', '55555555'),
            ('6','66666666')
        ]
        invalid_credentials = [
            ('test1','698443509'),
            ('test3', 'testPassword3'),
            ('5', '5582478972549'),
            ('7','77777777')
        ]

        for username, password in valid_credentials:
            self.assertTrue(self.am.authenticate_user_login(username, password))
        for username, password in invalid_credentials:
            self.assertFalse(self.am.authenticate_user_login(username,password))
            
    def broken_test(self):
        self.assertFalse(True);
