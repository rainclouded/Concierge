from incident_reports.incident_reports_server.persistence.IIncidentReportPersistence import IIncidentReportPersistence, IncidentReportPersistenceStub

class Services:
    def __init__(self) -> None:
        self._incident_report_persistence = None

    def get_incident_report_persistence(self) -> IIncidentReportPersistence:
        if (self._incident_report_persistence == None):
            _incident_report_persistence = IncidentReportPersistenceStub()
        return _incident_report_persistence
    
    def clear(self) -> None:
        self._incident_report_persistence = None;

        