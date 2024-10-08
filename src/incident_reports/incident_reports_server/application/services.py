from incident_reports_server.persistence.i_incident_report_persistence import IIncidentReportPersistence
from incident_reports_server.persistence.incident_report_persistence_stub import IncidentReportPersistenceStub
from incident_reports_server.validators.permission_validator import PermissionValidator 
from incident_reports_server.validators.i_permission_validator import IPermissionValidator

class Services:
    _incident_report_persistence = None
    _permission_validator = None
    
    @staticmethod
    def get_incident_report_persistence() -> IIncidentReportPersistence:
        if Services._incident_report_persistence is None:
            Services._incident_report_persistence = IncidentReportPersistenceStub()
            
        return Services._incident_report_persistence
    
    @staticmethod
    def get_permission_validator() -> IPermissionValidator:
        if Services._permission_validator is None:
            Services._permission_validator = PermissionValidator()
            
        return Services._permission_validator
    
    @staticmethod
    def clear() -> None:
        Services._incident_report_persistence = None
        Services._permission_validator = None
        