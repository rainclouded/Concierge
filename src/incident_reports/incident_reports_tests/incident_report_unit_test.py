import unittest
from flask import json
from incident_reports_server.factory.incident_report_factory import IncidentReportFactory
from incident_reports_server.application.services import Services
from incident_reports_server.controllers.incident_reports_controller import create_app
from incident_reports_server.models.models import Severity, Status, IncidentReport

class incident_report_unit_test(unittest.TestCase):
    valid_incident_report = IncidentReport(
        severity=Severity.HIGH,  
        status=Status.OPEN,       
        title="Fire Alarm Malfunction",
        description="The fire alarm in the lobby is continuously ringing but there is no fire. Immediate attention is required.",
        filing_person_id=303,    
        reviewer_id=403          
    )

    updated_incident_report_json = {
        "severity": "High",
        "status": "Open",
        "title": "Fire Alarm Malfunction",
        "description": "The fire alarm in the lobby is continuously ringing but there is no fire. Immediate attention is required.",
        "filing_person_id": 303,
        "reviewer_id": 403
    }
    
    invalid_incident_report_json = {
        "severity": "Awesome",
        "status": "Open",
        "title": "Fire Alarm Malfunction",
        "description": "The fire alarm in the lobby is continuously ringing but there is no fire. Immediate attention is required.",
        "filing_person_id": 303,
        "reviewer_id": 403
    }
    
    @staticmethod
    def from_byte_string(byte_string):
        return json.loads(byte_string.decode('utf-8'))[0]['data']
    
    def setUp(self):
        Services.clear()
        self._incident_report_persistence = Services.get_incident_report_persistence()

        self.app = create_app(self._incident_report_persistence).test_client()
        self.app.testing = True

    def test_get_incident_reports(self):
        response = self.app.get('/incident_reports')
        
        self.assertEqual(response.status_code, 200)
        self.assertIsNotNone(response.json['data'])
        
    def test_get_incident_report_by_id_valid_id_should_not_return_none(self):
        incident_report = self._incident_report_persistence.create_incident_report(self.valid_incident_report)
        
        response = self.app.get(f"/incident_reports/{incident_report.id}")
        self.assertEqual(response.status_code, 200)
        self.assertIsNotNone(response.json['data'])
        self.assertDictEqual(response.json['data'], incident_report.to_dict())
        pass

    def test_get_incident_report_by_id_invalid_id_returns_not_found(self):
        response = self.app.get("/incident_reports/-1")
        self.assertEqual(response.status_code, 404)

    def test_add_incident_report_valid_incident_report_is_successful(self):
        response = self.app.post("/incident_reports", data=json.dumps(self.valid_incident_report.to_dict()), content_type='application/json')
        self.assertEqual(response.status_code, 201) 

    def test_add_incident_report_valid_incident_report_able_to_be_fetched(self):
        response = self.app.post("/incident_reports", data=json.dumps(self.valid_incident_report.to_dict()), content_type='application/json')
        
        incident_report = IncidentReportFactory.create_incident_report(response.json['data'])
        
        response = self.app.get(f"/incident_reports/{incident_report.id}")
                
        self.assertEqual(response.status_code, 200) 

    def test_add_incident_report_invalid_incident_report_fails(self):
        response = self.app.post("/incident_reports", data=json.dumps(self.invalid_incident_report_json), content_type='application/json')
        self.assertEqual(response.status_code, 400) 

    def test_add_incident_report_null_incident_report_fails(self):
        response = self.app.post("/incident_reports", data=json.dumps(None), content_type='application/json')
        self.assertEqual(response.status_code, 400) 

    def test_update_incident_report_valid_incident_report_is_successful(self):
        incident_report = self._incident_report_persistence.create_incident_report(self.valid_incident_report)

        response = self.app.put(f"/incident_reports/{incident_report.id}", data=json.dumps(self.updated_incident_report_json), content_type='application/json')
        self.assertEqual(response.status_code, 200) 

    def test_update_incident_report_valid_incident_report_updated_incident_report_fetched(self):
        incident_report = self._incident_report_persistence.create_incident_report(self.valid_incident_report)

        self.app.put(f"/incident_reports/{incident_report.id}", data=json.dumps(self.updated_incident_report_json), content_type='application/json')
        
        response = self.app.get(f"/incident_reports/{incident_report.id}")        
        self.assertEqual(response.status_code, 200) 

    def test_update_incident_report_invalid_incident_report_fails(self):
        incident_report = self._incident_report_persistence.create_incident_report(self.valid_incident_report)

        response = self.app.put(f"/incident_reports/{incident_report.id}", data=json.dumps(self.invalid_incident_report_json), content_type='application/json')
        self.assertEqual(response.status_code, 400) 

    def test_update_incident_report_non_existing_incident_report_fails(self):
        response = self.app.put("/incident_reports/999", data=json.dumps(self.updated_incident_report_json), content_type='application/json')
        self.assertEqual(response.status_code, 404) 

    def test_update_incident_report_null_incident_report_fails(self):
        incident_report = self._incident_report_persistence.create_incident_report(self.valid_incident_report)

        response = self.app.put(f"/incident_reports/{incident_report.id}", data=json.dumps(None), content_type='application/json')
        self.assertEqual(response.status_code, 400) 

    def test_delete_incident_report_valid_incident_report_is_successful(self):
        incident_report = self._incident_report_persistence.create_incident_report(self.valid_incident_report)

        response = self.app.delete(f"/incident_reports/{incident_report.id}")
        self.assertEqual(response.status_code, 200) 
        
    def test_delete_incident_report_valid_incident_report_is_fetchable(self):
        incident_report = self._incident_report_persistence.create_incident_report(self.valid_incident_report)

        self.app.delete(f"/incident_reports/{incident_report.id}")
        
        response = self.app.get(f"/incident_reports/{incident_report.id}")        
        self.assertEqual(response.status_code, 404) 


    def test_delete_incident_report_invalid_incident_report_fails(self):
        incident_report = self._incident_report_persistence.create_incident_report(self.valid_incident_report)

        response = self.app.delete(f"/incident_reports/{999}")
        self.assertEqual(response.status_code, 404) 

    def test_delete_incident_report_invalid_incident_report_not_fetchable(self):
        incident_report = self._incident_report_persistence.create_incident_report(self.valid_incident_report)

        self.app.delete(f"/incident_reports/{999}")
        
        response = self.app.get(f"/incident_reports/{999}")        
        self.assertEqual(response.status_code, 404) 

if __name__ == '__main__':
    unittest.main()