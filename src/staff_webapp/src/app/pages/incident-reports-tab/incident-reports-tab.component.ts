import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { IncidentReportService } from '../../services/incident-report.service';
import { IIncidentReport } from '../../models/incident-report.model';
import { WindowComponent } from '../../components/window/window.component'; 
import { IncidentReportFormComponent } from '../../components/incident-report-form/incident-report-form.component';
import { ConfirmationDialogComponent } from "../../components/confirmation-dialog/confirmation-dialog.component";
import { ToastrService } from 'ngx-toastr';
import { SessionService } from '../../services/session.service';

@Component({
  selector: 'app-incident-reports-tab',
  standalone: true,
  imports: [WindowComponent, IncidentReportFormComponent, CommonModule, ConfirmationDialogComponent],
  templateUrl: './incident-reports-tab.component.html',
})
export class IncidentReportsTabComponent implements OnInit {

  isReportWindowOpen = false;
  isConfirmWindowOpen = false;
  reportToDelete: number | null = null;
  incidentReport !: IIncidentReport;
  incidentReports : IIncidentReport[] = [];
  
  todoReports: IIncidentReport[] = [];
  inProgressReports: IIncidentReport[] = [];
  resolvedReports: IIncidentReport[] = [];
  closedReports: IIncidentReport[] = [];

  sessionPermissionList: string[] | null = null;
  canCreate: boolean = false;
  canEdit: boolean = false;
  canDelete: boolean = false;

  constructor(
    private incidentReportService: IncidentReportService,
    private sessionService: SessionService,
    private toastr: ToastrService
  ) { }

  ngOnInit():void {
    this.getAllReports();
    this.sessionService.getSessionMe().subscribe(() => {
      this.sessionPermissionList = this.sessionService.sessionPermissionList;
      this.checkPermissions();
      console.log();
    });
  }

  checkPermissions():void {
    if (this.sessionPermissionList) {
      console.log('reached here')
      this.canCreate = this.sessionPermissionList.includes('canCreateIncidentReports');
      this.canEdit = this.sessionPermissionList.includes('canEditIncidentReports');
      this.canDelete = this.sessionPermissionList.includes('canDeleteIncidentReports');
    }
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
  loadReport(incidentReport: IIncidentReport) {
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
          this.toastr.success('Incident report deleted successfully!', 'Delete Successful');
          this.closeConfirmWindow();
        },
        error: (error) => {
          console.error('Error deleting incident report:', error);
          this.toastr.error('Error deleting incident report!', 'Delete Failed');
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
