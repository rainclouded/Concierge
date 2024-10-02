import json
from incident_reports_server.model.models import IncidentReport

class IncidentReportFactory:

    #converts json object to incident report object
    @staticmethod
    def create_incident_report(incident_report_JSON : json) -> IncidentReport:
        required_keys = ["severity","status","title","description","filing_person_id","reviewer_Id"]
        
        if incident_report_JSON is not json:
            raise ValueError("No JSON object was passed!")
        
        missing_keys = [key for key in required_keys if key not in incident_report_JSON]
        if missing_keys:
            raise KeyError(f"Values are missing in incident report: {','.join(missing_keys)}")
                
        return IncidentReport(severity = incident_report_JSON.severity, status = incident_report_JSON.status, title = incident_report_JSON.title, 
                              description = incident_report_JSON.description, filing_person_Id = incident_report_JSON.filing_person_id, reviewer_Id = incident_report_JSON.reviewer_Id)

        
 
        
        
        
        
     