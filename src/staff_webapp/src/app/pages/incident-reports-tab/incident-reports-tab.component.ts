import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { IncidentReportService } from '../../services/incident-report.service';
import { IIncidentReport } from '../../models/incident-report.model';
import { WindowComponent } from '../../components/window/window.component'; 
import { IncidentReportFormComponent } from '../../components/incident-report-form/incident-report-form.component';
import { ConfirmationDialogComponent } from "../../components/confirmation-dialog/confirmation-dialog.component";

@Component({
  selector: 'app-incident-reports-tab',
  standalone: true,
  imports: [WindowComponent, IncidentReportFormComponent, CommonModule, ConfirmationDialogComponent],
  templateUrl: './incident-reports-tab.component.html',
})
export class IncidentReportsTabComponent {
  isReportWindowOpen = false;
  isConfirmWindowOpen = false;
  reportToDelete: number | null = null;
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
    this.openReportWindow();
  }

  deleteReport() {
    if (this.reportToDelete !== null) {
      this.incidentReportService.deleteReport(this.reportToDelete)
      .subscribe({
        next: (response) => {
          this.getAllReports();
          console.log(response.message);
          this.closeConfirmWindow();
        },
        error: (error) => {
          console.error('Error deleting incident report:', error);
        }
      });
    } else {
      console.log('Incident Report ID is null, cannot delete.');
    }
  }

  openReportWindow() {
    this.isReportWindowOpen = true;
  }

  closeReportWindow() {
    this.isReportWindowOpen = false;
    this.getAllReports();
  }

  openConfirmWindow(reportId: number) {
    this.isConfirmWindowOpen = true;
    this.reportToDelete = reportId;
  }

  closeConfirmWindow() {
    this.isConfirmWindowOpen = false;
    this.getAllReports();
  }

  categorizeReports(): void {
    this.todoReports = this.incidentReports.filter(report => report.status == 'OPEN');
    this.inProgressReports = this.incidentReports.filter(report => report.status == 'IN_PROGRESS');
    this.resolvedReports = this.incidentReports.filter(report => report.status == 'RESOLVED');
    this.closedReports = this.incidentReports.filter(report => report.status == 'CLOSED')
  }
}
