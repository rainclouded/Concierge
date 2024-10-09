import { Component } from '@angular/core';
import { IncidentReportService } from '../../services/incident-report.service';
import { IIncidentReport } from '../../models/incident-report.model';
import { WindowComponent } from '../../components/window/window.component'; 
import { IncidentReportFormComponent } from '../../components/incident-report-form/incident-report-form.component';

@Component({
  selector: 'app-incident-reports-tab',
  standalone: true,
  imports: [WindowComponent, IncidentReportFormComponent],
  templateUrl: './incident-reports-tab.component.html',
  styleUrl: './incident-reports-tab.component.css'
})
export class IncidentReportsTabComponent {
  isOpenWindow = false;
  incidentReport !: IIncidentReport;
  incidentReports : IIncidentReport[] = [];
  
  todoReports: IIncidentReport[] = [];
  inProgressReports: IIncidentReport[] = [];
  resolvedReports: IIncidentReport[] = [];
  closedReports: IIncidentReport[] = [];

  constructor(
    private incidentReportService: IncidentReportService
  ) { }

  ngOnInit():void {
    this.getAllReports();
  }

  getAllReports() {
    this.incidentReportService.getAllReports()
    .subscribe({
      next: (response) => {
        if (response.data) {
          this.incidentReports = response.data;
          this.categorizeReports();
        }
      } 
    });
  }

  // load current incident report onto IncidentReportForm to edit
  loadAmenity(incidentReport: IIncidentReport) {
    this.incidentReport = incidentReport;
    this.openWindow();
  }

  deleteReport(id: number) {
    this.incidentReportService.deleteReport(id)
    .subscribe({
      next: (response) => {
        this.getAllReports();
        console.log(response.message);
      }
    });
  }

  openWindow() {
    this.isOpenWindow = true;
  }

  closeWindow() {
    this.isOpenWindow = false;
    this.getAllReports();
  }

  categorizeReports(): void {
    this.todoReports = this.incidentReports.filter(report => report.status == 'OPEN');
    this.inProgressReports = this.incidentReports.filter(report => report.status == 'IN_PROGRESS');
    this.resolvedReports = this.incidentReports.filter(report => report.status == 'RESOLVED');
    this.closedReports = this.incidentReports.filter(report => report.status == 'CLOSED')
  }
}
