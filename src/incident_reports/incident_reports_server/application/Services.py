from incident_reports.incident_reports_server.persistence.IIncidentReportPersistence import IIncidentReportPersistence, IncidentReportPersistenceStub

class Services:
    _incident_report_persistence = None

    @staticmethod
    def get_incident_report_persistence() -> IIncidentReportPersistence:
        if Services._incident_report_persistence is None:
            Services._incident_report_persistence = IncidentReportPersistenceStub()
        return Services._incident_report_persistence
    
    @staticmethod
    def clear() -> None:
        Services._incident_report_persistence = None
        