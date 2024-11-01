import unittest
import os 
from flask import json
from incident_reports_server.validators.mock_permission_validator import MockPermissionValidator
from incident_reports_server.factory.incident_report_factory import IncidentReportFactory
from incident_reports_server.application.services import Services
from incident_reports_server.controllers.incident_reports_controller import create_app
from incident_reports_server.models.models import Severity, Status, IncidentReport

class incident_report_unit_test(unittest.TestCase):    
    @staticmethod
    def from_byte_string(byte_string):
        return json.loads(byte_string.decode('utf-8'))[0]['data']
    
    def setUp(self):
        os.environ['DB_IMPLEMENTATION'] = 'MOCK' 
        
        self.valid_incident_report = IncidentReport(
            severity=Severity.HIGH,  
            status=Status.OPEN,       
            title="Fire Alarm Malfunction",
            description="The fire alarm in the lobby is continuously ringing but there is no fire. Immediate attention is required.",
            filing_person_id=303,    
            reviewer_id=403          
        )

        self.updated_incident_report_json = {
            "severity": "HIGH",
            "status": "Open",
            "title": "Fire Alarm Malfunction",
            "description": "The fire alarm in the lobby is continuously ringing but there is no fire. Immediate attention is required.",
            "filing_person_id": 303,
            "reviewer_id": 403
        }
        
        self.invalid_incident_report_json = {
            "severity": "High",
            "status": "Open",
            "title": "Fire Alarm Malfunction",
            "description": "The fire alarm in the lobby is continuously ringing but there is no fire. Immediate attention is required.",
            "filing_person_id": 303,
            "reviewer_id": 403
        }
        
        Services.clear()
        self._incident_report_persistence = Services.get_incident_report_persistence()
        self._permission_validator = MockPermissionValidator()

        self.app = create_app(self._incident_report_persistence, self._permission_validator).test_client()
        self.app.testing = True

    def test_get_incident_reports_filter_severity(self):
        response = self.app.get('/incident_reports/?severity=HIGH')
        
        self.assertEqual(response.status_code, 200)

    def test_get_incident_reports_filter_status(self):
        response = self.app.get('/incident_reports/?status=Closed')
        
        self.assertEqual(response.status_code, 200)
    
    def test_get_incident_reports_filter_by_dates(self):
        response = self.app.get('/incident_reports/?beforeDate=2024-01-03&afterDate=2024-01-01')
        
        self.assertEqual(response.status_code, 200)

    def test_get_incident_reports_no_filters(self):
        response = self.app.get('/incident_reports/')
        
        self.assertEqual(response.status_code, 200)
        
    def test_get_incident_reports_filter_invalid_severity(self):
        response = self.app.get('/incident_reports/?severity=SUPERHIGH')
        
        self.assertEqual(response.status_code, 400)

    def test_get_incident_reports_filter_invalid_status(self):
        response = self.app.get('/incident_reports/?status=YUH')
        
        self.assertEqual(response.status_code, 400)
    
    def test_get_incident_reports_filter_by_invalid_dates(self):
        response = self.app.get('/incident_reports/?beforeDate=123&afterDate=231')
        
        self.assertEqual(response.status_code, 400)

    def test_get_incident_report_by_id_valid_id_should_not_return_none(self):
        incident_report = self._incident_report_persistence.create_incident_report(self.valid_incident_report)
        
        response = self.app.get(f"/incident_reports/{incident_report.id}")
        self.assertEqual(response.status_code, 200)
        self.assertIsNotNone(response.json['data'])
        self.assertDictEqual(response.json['data'], incident_report.to_dict())

    def test_get_incident_report_by_id_invalid_id_returns_not_found(self):
        response = self.app.get("/incident_reports/-1")
        self.assertEqual(response.status_code, 404)

    def test_add_incident_report_valid_incident_report_is_successful(self):
        response = self.app.post("/incident_reports/", data=json.dumps(self.valid_incident_report.to_dict()), content_type='application/json')
        self.assertEqual(response.status_code, 201) 

    def test_add_incident_report_valid_incident_report_able_to_be_fetched(self):
        response = self.app.post("/incident_reports/", data=json.dumps(self.valid_incident_report.to_dict()), content_type='application/json')
        
        self.assertEqual(response.status_code, 201) 
        
        incident_report = IncidentReportFactory.create_incident_report(response.json['data'])
        
        response = self.app.get(f"/incident_reports/{incident_report.id}")
                
        self.assertEqual(response.status_code, 200) 

    def test_add_incident_report_invalid_incident_report_title_fails(self):
        self.invalid_incident_report_json["title"] = ""
        response = self.app.post("/incident_reports/", data=json.dumps(self.invalid_incident_report_json), content_type='application/json')
        self.assertEqual(response.status_code, 400) 

    def test_add_incident_report_null_incident_report_fails(self):
        response = self.app.post("/incident_reports/", data=json.dumps(None), content_type='application/json')
        self.assertEqual(response.status_code, 400) 

    def test_add_incident_report_invalid_incident_report_desc_fails(self):
        self.invalid_incident_report_json["description"] = ""
        response = self.app.post("/incident_reports/", data=json.dumps(self.invalid_incident_report_json), content_type='application/json')
        self.assertEqual(response.status_code, 400) 
        
    def test_add_incident_report_invalid_incident_report_filing_person_id_fails(self):
        self.invalid_incident_report_json["filing_person_id"] = ""
        response = self.app.post("/incident_reports/", data=json.dumps(self.invalid_incident_report_json), content_type='application/json')
        self.assertEqual(response.status_code, 400) 
        
    def test_add_incident_report_invalid_incident_report_reviewer_id_fails(self):
        self.invalid_incident_report_json["reviewer_id"] = ""
        response = self.app.post("/incident_reports/", data=json.dumps(self.invalid_incident_report_json), content_type='application/json')
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
        self.invalid_incident_report_json["title"] = ""
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