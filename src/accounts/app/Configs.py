import hashlib

#Criteria for a staff password
PASSWORD_HASH_FUNCTION = hashlib.sha256
PASSWORD_MINIMAL_LENGTH = 5
PASSWORD_MUST_CONTAIN_LETTER = True
PASSWORD_MUST_CONTAIN_NUMBER = True
MAX_GUEST_PASSWORD = 100000000

#Criteria for a staff username
USERNAME_MINIMAL_LENGTH = 5
USERNAME_MUST_CONTAIN_LETTER = True
USERNAME_MUST_CONTAIN_NUMBER = True

#The Different fields a user can have
USER_ATTRIBUTES = ['username', 'password', 'hash','type', 'id']
