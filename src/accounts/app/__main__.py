import app.server.Accounts as account_service
from . import create_database

if __name__ == "__main__":
    account_service.start_service(create_database())