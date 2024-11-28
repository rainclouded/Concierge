import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SessionService } from '../../services/session.service';
import { IAmenity } from '../../models/amenity.model';
import { AmenityService } from '../../services/amenity.service';
import { IncidentReportService } from '../../services/incident-report.service';
import { ITask } from '../../models/tasks.model';
import { TaskService } from '../../services/task.service';

@Component({
  selector: 'app-home-tab',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './home-tab.component.html',
})
export class HomeTabComponent implements OnInit {
  accountName: string | null = null;
  amenities: IAmenity[] = [];
  amenity !: IAmenity;

  reportCounts = {
    OPEN: 0,
    IN_PROGRESS: 0,
    RESOLVED: 0,
    CLOSED: 0
  };  

  tasks: ITask[] = [];

  constructor(
    private sessionService: SessionService,
    private amenityService: AmenityService,
    private incidentReportService: IncidentReportService,
    private taskService: TaskService
  ) {}

  ngOnInit() {
    this.refreshPage();
    this.sessionService.getSessionMe().subscribe(() => {
      this.accountName = this.sessionService.accountName;
    });
  }

  refreshPage(): void {
    this.getAllAmenities();
    this.getReportCounts();
    this.getPendingTasks();
  }

  // GET all amenities
  getAllAmenities() {
    this.amenityService.getAllAmenities()
      .subscribe({
        next: (response) => {
          if (response.data) {
            this.amenities = response.data;
          }
        },
        error: (error) => {
          console.error('Error fetching reports:', error);
        }
      });
  }

  // GET reports count
  getReportCounts() {
    this.incidentReportService.getAllReports()
      .subscribe({
        next: (response) => {
          if (response.data) {
            const reports = response.data;
            this.reportCounts.OPEN = reports.filter(report => report.status === 'OPEN').length;
            this.reportCounts.IN_PROGRESS = reports.filter(report => report.status === 'IN_PROGRESS').length;
            this.reportCounts.RESOLVED = reports.filter(report => report.status === 'RESOLVED').length;
            this.reportCounts.CLOSED = reports.filter(report => report.status === 'CLOSED').length;
          }
        },
        error: (error) => {
          console.error('Error fetching reports:', error);
        }
      });
  }

  // GET all pending tasks
  getPendingTasks(): void {
    this.taskService.getTasksByStatus('Pending').subscribe({
      next: (response) => {
        if (response.data) {
          this.tasks = response.data;
          this.tasks.reverse(); // Displaying the most recent one on top
          console.log('Pending tasks:', this.tasks);
        }
      },
      error: (error) => {
        console.error('Error fetching pending tasks:', error);
      }
    });
  }

  // Helper: time converter
  formatTimeToAMPM(time: string): string {
    if (!time) return '';
    const [hours, minutes] = time.split(':');
    const date = new Date();
    date.setHours(parseInt(hours, 10), parseInt(minutes, 10));

    return date.toLocaleTimeString('en-US', {
      hour: '2-digit',
      minute: '2-digit',
      hour12: true
    });
  }
}
