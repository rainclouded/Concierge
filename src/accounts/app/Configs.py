"""
Default Configurations for the account service
"""
import hashlib
from app.permissions.ClientPermissions import ClientPermissionValidator
from app.database.Mockdata import Mockdata
from app.permissions.MockPermissions import MockPermissions
from app.database.MongoConnection import MongoConnection
from os import getenv

#Criteria for a staff password
PASSWORD_HASH_FUNCTION = hashlib.sha256
PASSWORD_MINIMAL_LENGTH = 5
PASSWORD_MUST_CONTAIN_LETTER = True
PASSWORD_MUST_CONTAIN_NUMBER = True
MAX_GUEST_PASSWORD = 100000000

#Encryption algorithm for the jwt
JWT_ENCRYPTION = 'ES256'

#Criteria for a staff username
USERNAME_MINIMAL_LENGTH = 5
USERNAME_MUST_CONTAIN_LETTER = True
USERNAME_MUST_CONTAIN_NUMBER = True

#The Different fields a user can have
USER_ATTRIBUTES = ['username', 'hash','type', 'id']
GUEST_TYPE = 'guest'
STAFF_TYPE = 'staff'

#The database to use
if getenv('DEPLOYMENT') == 'production':
    print("Running in Prod")
    DATABASE = MongoConnection()
    PERMISSIONS = ClientPermissionValidator()
elif getenv('DEPLOYMENT') == 'development':
    print("Running in dev")
    DATABASE = MongoConnection()
    PERMISSIONS = MockPermissions()
else:
    print("Running in dev")
    DATABASE = Mockdata()
    PERMISSIONS = MockPermissions()

def create_database():
    return DATABASE

def create_permissions():
    return PERMISSIONS