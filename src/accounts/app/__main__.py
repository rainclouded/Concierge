import app.server.Accounts as account_service
from app.core.Services import Services as dependencies
if __name__ == "__main__":
    dependencies.set_up()
    print(dependencies._user_service)
    account_service.start_service()
    