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
            
        return Services._incident_report_persistence
    
    @staticmethod
    def _construct_incident_report_persistence() -> IIncidentReportPersistence:
        db_implementation = os.getenv("DB_IMPLEMENTATION")
        
        if(db_implementation == "MONGODB"):
            print(f"Attempting to connect to MongoDB with connection string: {Services.db_connection_string()}")
            try:
                db_implementation = IncidentReportPersistenceMongo(Services.db_connection_string(), os.getenv("DB_NAME", "test_concierge"))
                print("Successfully connected to MongoDB")
            except Exception as e:
                db_implementation = None
                print(f'MongoDB failed to connect: {e}')
        elif(db_implementation == "MOCK"):
            db_implementation = IncidentReportPersistenceStub()
        
        if(not db_implementation):
            db_implementation = IncidentReportPersistenceStub()
        
        return db_implementation
    
    @staticmethod
    def db_connection_string() -> str:
        mongo_host = os.getenv('DB_HOST', 'mongo')
        mongo_db_name = os.getenv("DB_NAME", "test_concierge")
        mongo_port = int(os.getenv('DB_PORT', 27017))
        mongo_username = os.getenv('DB_USERNAME', 'mongo_db_user')
        mongo_password = os.getenv('DB_PASSWORD', 'password')
        
        return f'mongodb://{mongo_username}:{mongo_password}@{mongo_host}:{mongo_port}/{mongo_db_name}'

    @staticmethod
    def clear() -> None:
        if Services._incident_report_persistence:
            Services._incident_report_persistence.clear()
            Services._incident_report_persistence = None
        