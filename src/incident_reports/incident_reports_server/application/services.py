from incident_reports_server.validators.permission_validator import PermissionValidator
from incident_reports_server.persistence.i_incident_report_persistence import IIncidentReportPersistence
from incident_reports_server.persistence.incident_report_persistence_mongo import IncidentReportPersistenceMongo
from incident_reports_server.persistence.incident_report_persistence_stub import IncidentReportPersistenceStub
import os

class Services:
    _incident_report_persistence = None 
    _permission_validator = None
    
    @staticmethod
    def get_incident_report_persistence() -> IIncidentReportPersistence:
        """
        Retrieves the singleton instance of the Incident Report Persistence class.
        If it does not exist, it constructs it using the appropriate persistence layer (MongoDB or Mock).
        
        Returns:
            IIncidentReportPersistence: The persistence instance (either MongoDB or Mock).
        """
        if Services._incident_report_persistence is None:
            Services._incident_report_persistence = Services._construct_incident_report_persistence()
            
        return Services._incident_report_persistence 
    
    @staticmethod
    def get_permission_validator():
        """
        Retrieves the singleton instance of the Permission Validator.
        If it does not exist, it constructs it using the default constructor.
        
        Returns:
            PermissionValidator: The permission validator instance.
        """
        if Services._permission_validator is None:
            Services._permission_validator = Services._construct_permission_validator()
        
        return Services._permission_validator
    
    @staticmethod
    def _construct_incident_report_persistence() -> IIncidentReportPersistence:
        """
        Constructs the appropriate persistence instance based on the DB_IMPLEMENTATION environment variable.
        If MongoDB is specified, it attempts to establish a connection using the provided connection string.
        If Mock is specified, it uses a mock/stub implementation.
        
        Returns:
            IIncidentReportPersistence: The constructed persistence instance (either MongoDB or Mock).
        """
        # Construct a persistence instance, whether it constructs mock or a mongodb connection depends on the env variable 'DB_IMPLEMENTATION' initialized by the docker file
        db_implementation = os.getenv("DB_IMPLEMENTATION") 
        
        if db_implementation == "MONGODB":
            """
            Attempt to connect to MongoDB using the connection string.
            """
            print(f"Attempting to connect to MongoDB with connection string: {Services.db_connection_string()}")
            try:
                # Initialize MongoDB persistence with the connection string and DB name
                db_implementation = IncidentReportPersistenceMongo(Services.db_connection_string())
                print("Successfully connected to MongoDB")
            except Exception as e:
                # If the connection fails, log the exception and use a stub instead
                db_implementation = None
                print(f'MongoDB failed to connect: {e}')
        
        elif db_implementation == "MOCK":
            """
            Use a mock/stub implementation for testing or alternative environments.
            """
            db_implementation = IncidentReportPersistenceStub()
        
        if not db_implementation:
            db_implementation = IncidentReportPersistenceStub()
        
        return db_implementation 
    
    @staticmethod
    def _construct_permission_validator():
        """
        Constructs and returns the PermissionValidator instance.
        
        Returns:
            PermissionValidator: The constructed permission validator instance.
        """
        return PermissionValidator()
    
    @staticmethod
    def db_connection_string() -> str:
        """
        Constructs the MongoDB connection string using environment variables (DB_HOST, DB_PORT, DB_USERNAME, DB_PASSWORD).
        This string is used to connect to the MongoDB instance for persistent storage.
        
        Returns:
            str: The MongoDB connection string.
        """
        # If a mongodb connection is to be established, pull connection details from the docker file
        mongo_host = os.getenv('DB_HOST', 'mongo')
        mongo_port = int(os.getenv('DB_PORT', 27017))
        mongo_username = os.getenv('DB_USERNAME', 'mongo_db_user')
        mongo_password = os.getenv('DB_PASSWORD', 'password')
        
        # Format the connection string
        return f'mongodb://{mongo_username}:{mongo_password}@{mongo_host}:{mongo_port}'

    @staticmethod
    def clear() -> None:
        """
        Clears the singleton instance of the Incident Report Persistence and releases any resources.
        This is useful for cleaning up and resetting the service state.
        """
        # If instance exists, clear the data and release the singleton
        if Services._incident_report_persistence:
            Services._incident_report_persistence.clear()  # Clear the persistence data
            Services._incident_report_persistence = None
