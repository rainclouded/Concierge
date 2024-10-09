import { Component, EventEmitter, Input, Output } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { IIncidentReport } from '../../models/incident-report.model';
import { IncidentReportService } from '../../services/incident-report.service';


@Component({
  selector: 'app-incident-report-form',
  standalone: true,
  imports: [ReactiveFormsModule],
  templateUrl: './incident-report-form.component.html',
})
export class IncidentReportFormComponent {
  // custom event enabling communication from a child component (report form) to a parent component (incident reports page/tab)
  @Input() data: IIncidentReport | null = null;
  @Output() onCloseWindow = new EventEmitter();
  reportForm: FormGroup;

  constructor (
    private fb: FormBuilder,
    private reportService: IncidentReportService
  ) {
    this.reportForm = this.fb.group({
      title: new FormControl<string>('', [Validators.required]),
      description: new FormControl<string>('', [Validators.required]),
      severity: new FormControl<string>('', [Validators.required]),
      status: new FormControl<string>('', [Validators.required]),
      filing_person_id: new FormControl<number>(301), // placeholder for now
      reviewer_id: new FormControl<number>(404), // placeholder for now
    });
  }

  onClose() {
    this.reportForm.reset();
    this.data = null;
    this.onCloseWindow.emit(false);
  }

  ngOnChanges(): void {
    if (this.data) {
      this.reportForm.patchValue({
        title: this.data.title,
        description: this.data.description,
        severity: this.data.severity,
        status: this.data.status,
        filing_person_id: this.data.filing_person_id,
        reviewer_id: this.data.reviewer_id,
      });
    } else {
      this.reportForm.reset();
    }
  }

  onSubmit() {
    if (this.reportForm.valid) {
      const timestamp = new Date().toISOString();
      
      if (this.data) {
        const updateReport = {
          ...this.reportForm.value,
          updated_at: timestamp,
        }

        this.reportService
        .updateReport(this.data.id as number, updateReport)
        .subscribe({
          next: (response) => {
            this.onClose();
            console.log(response.message);
          }
        });
      } else {
        const newReport = {
          ...this.reportForm.value,
          filing_person_id: 301, // placeholder for now
          reviewer_id: 404, // placeholder for now
        }

        this.reportService.addReport(newReport)
        .subscribe({
          next: (response) => {
            this.onClose();
            console.log(response.message);
          }
        });
      }
    } else {
      this.reportForm.markAllAsTouched();
    }
  }
}
