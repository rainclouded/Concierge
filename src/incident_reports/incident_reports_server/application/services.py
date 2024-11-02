from incident_reports_server.validators.permission_validator import PermissionValidator
from incident_reports_server.persistence.i_incident_report_persistence import IIncidentReportPersistence
from incident_reports_server.persistence.incident_report_persistence_mongo import IncidentReportPersistenceMongo
from incident_reports_server.persistence.incident_report_persistence_stub import IncidentReportPersistenceStub
import os

class Services:
    _incident_report_persistence = None 
    
    @staticmethod
    def get_incident_report_persistence() -> IIncidentReportPersistence:
        if Services._incident_report_persistence is None:
            Services._incident_report_persistence = Services._construct_incident_report_persistence()
            
        return Services._incident_report_persistence  # Return the persistence instance
    
    @staticmethod
    def _construct_incident_report_persistence() -> IIncidentReportPersistence:
        # Construct a persistence instance, whether it constructs mock or a mongodb connection depends on the env variable 'DB_IMPLEMENTATION' initialized by the docker file
        db_implementation = os.getenv("DB_IMPLEMENTATION") 
        
        if db_implementation == "MONGODB":
            # Attempt to connect to MongoDB using the connection string
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
            # Use a mock/stub implementation for testing or alternative environments
            db_implementation = IncidentReportPersistenceStub()
        
        if not db_implementation:
            db_implementation = IncidentReportPersistenceStub()
        
        return db_implementation 
    
    @staticmethod
    def db_connection_string() -> str:
        # If a mongodb connection is to established, pull connection details from the docker file
        mongo_host = os.getenv('DB_HOST', 'mongo')
        mongo_port = int(os.getenv('DB_PORT', 27017))
        mongo_username = os.getenv('DB_USERNAME', 'mongo_db_user')
        mongo_password = os.getenv('DB_PASSWORD', 'password')
        
        #Format the connection string
        return f'mongodb://{mongo_username}:{mongo_password}@{mongo_host}:{mongo_port}'

    @staticmethod
    def clear() -> None:
        #If instance exists, clear the data and release the singleton
        if Services._incident_report_persistence:
            Services._incident_report_persistence.clear()  #Clear the persistence data
            Services._incident_report_persistence = None 

        