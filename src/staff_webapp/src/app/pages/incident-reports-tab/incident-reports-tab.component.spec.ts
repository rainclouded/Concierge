import { ComponentFixture, TestBed } from '@angular/core/testing';

import { WindowComponent } from '../../components/window/window.component';
import { IncidentReportFormComponent } from '../../components/incident-report-form/incident-report-form.component';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { IncidentReportsTabComponent } from './incident-reports-tab.component';

describe('IncidentReportsTabComponent', () => {
  let component: IncidentReportsTabComponent;
  let fixture: ComponentFixture<IncidentReportsTabComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [WindowComponent, IncidentReportFormComponent, HttpClientTestingModule]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(IncidentReportsTabComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
