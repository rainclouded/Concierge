db = db.getSiblingDB('concierge');

db.createUser({
    user: 'mongo_db_user',
    pwd: 'password',
    roles: [
        { role: 'readWrite', db: 'concierge' },
        { role: 'dbAdmin', db: 'concierge' }
    ]
});