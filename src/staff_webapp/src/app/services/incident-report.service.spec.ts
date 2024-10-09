import { TestBed } from '@angular/core/testing';

import { IncidentReportService } from './incident-report.service';
import { HttpClientTestingModule } from '@angular/common/http/testing';

describe('IncidentReportService', () => {
  let service: IncidentReportService;

  beforeEach(() => {
    TestBed.configureTestingModule(
      {
        imports: [HttpClientTestingModule],
      }
    );
    service = TestBed.inject(IncidentReportService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
