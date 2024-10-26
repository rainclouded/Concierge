import app.server.Accounts as account_service
from app.core.Services import Services as dependencies
if __name__ == "__main__":
    dependencies.set_up()
    account_service.start_service()
    