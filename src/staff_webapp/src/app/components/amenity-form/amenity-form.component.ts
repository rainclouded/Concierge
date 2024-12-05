import { Component, EventEmitter, Input, OnChanges, Output } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { IAmenity } from '../../models/amenity.model';
import { AmenityService } from '../../services/amenity.service';
import { HttpErrorResponse } from '@angular/common/http';
import { ToastrService } from 'ngx-toastr';

@Component({
  selector: 'app-amenity-form',
  standalone: true,
  imports: [ReactiveFormsModule],
  templateUrl: './amenity-form.component.html',
})
export class AmenityFormComponent implements OnChanges {
  // custom event enabling communication from a child component (amenity form) to a parent component (amenities page/tab)
  @Input() data: IAmenity | null = null;
  @Output() onCloseWindow = new EventEmitter();
  amenityForm: FormGroup;

  constructor(
    private fb: FormBuilder,
    private amenityService: AmenityService,
    private toastr: ToastrService
  ) {
    this.amenityForm = this.fb.group({
      name: new FormControl<string>('', [Validators.required]),
      description: new FormControl<string>('', [Validators.required]),
      startTime: new FormControl<string>('', [Validators.required]),
      endTime: new FormControl<string>('', [Validators.required]),
    });
  }

  onClose() {
    this.amenityForm.reset();
    this.data = null;
    this.onCloseWindow.emit(false);
  }

  ngOnChanges(): void {
    if (this.data) {
      this.amenityForm.patchValue({
        name: this.data.name,
        description: this.data.description,
        startTime: this.data.startTime,
        endTime: this.data.endTime,
      });
    } else {
      this.amenityForm.reset();
    }
  }

  onSubmit() {
    if (this.amenityForm.valid) {
      // Convert startTime and endTime to HH:mm:ss format
      const formattedData = {
        ...this.amenityForm.value,
        startTime: this.formatTime(this.amenityForm.value.startTime),
        endTime: this.formatTime(this.amenityForm.value.endTime)
      };

      if (this.data) {
        this.amenityService
          .updateAmenity(this.data.id as number, formattedData)
          .subscribe({
            next: (response: any) => {
              this.onClose();
              console.log(response.message);
              if (response.status === 400) {
                this.toastr.error('You have entered invalid data for the amenity!', 'Update Failed');
              } else {
                this.toastr.success('Amenity updated successfully!', 'Update Successful');
              }
            },
            error: (error: HttpErrorResponse) => {
              this.onClose();
              console.error(`Error Status: ${error.status}`);
              console.error(error.message);
              
              //show error to client
              //if (error.status === 400) alert("You have entered invalid data for the amenity!");
            }
          });
      } else {
        this.amenityService.addAmenity(formattedData)
          .subscribe({
            next: (response: any) => {
              this.onClose();
              console.log(response.message);
              if (response.status === 400) {
                this.toastr.error('You have entered invalid data for the amenity!', 'Add Failed');
              } else {
                this.toastr.success('Amenity added successfully!', 'Add Successful');
              }
            },
            error: (error: HttpErrorResponse) => {
              this.onClose();
              console.error(`Error Status: ${error.status}`);
              console.error(error.message);
              
               //show error to client
              //if (error.status === 400) alert("You have entered invalid data for the amenity!");
            }
          });
      }
    } else {
      this.amenityForm.markAllAsTouched();
    }
  }

  // Helper function to format time in HH:mm:ss
  private formatTime(time: string): string {
    const [hour, minute] = time.split(':');
    return `${hour.padStart(2, '0')}:${minute.padStart(2, '0')}:00`;
  }
}
