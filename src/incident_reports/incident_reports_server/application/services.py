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
            print("Attempting to connect to MongoDB")
            try:
                db_implementation = IncidentReportPersistenceMongo(Services.db_connection_string())
            except Exception as e:
                db_implementation = None
                print("MongoDB Failed!")
        
        if(not db_implementation):
            db_implementation = IncidentReportPersistenceStub()
        
        return db_implementation
    
    @staticmethod
    def db_connection_string() -> str:
        db_host = os.getenv("DB_HOST")
        db_port = os.getenv("DB_PORT")
        db_username = os.getenv("DB_USERNAME")
        db_password = os.getenv("DB_PASSWORD")
        
        return f"mongodb://{db_username}:{db_password}@{db_host}:{db_port}"

    @staticmethod
    def clear() -> None:
        Services._incident_report_persistence = None
        