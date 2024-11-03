import { Component, EventEmitter, Input, OnChanges, Output } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { IAmenity } from '../../models/amenity.model';
import { AmenityService } from '../../services/amenity.service';
import { HttpErrorResponse } from '@angular/common/http';


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
    private amenityService: AmenityService
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
      //format time input into timespan acceptable format
      this.amenityForm.value.startTime! = this.formatTime(this.amenityForm.value.startTime!);
      this.amenityForm.value.endTime! = this.formatTime(this.amenityForm.value.endTime!);

      if (this.data) {
        this.amenityService
          .updateAmenity(this.data.id as number, this.amenityForm.value)
          .subscribe({
            next: (response: any) => {
              this.onClose();
              console.log(response.message);
              if (response.status === 400) alert("You have entered invalid data for the amenity!");
            },
            error: (error: HttpErrorResponse) => {
              this.onClose();
              console.error(`Error Status: ${error.status}`);
              console.error(error.message);
              
              //show error to client
              //if (error.status === 400) alert("You have entered invalid data for the amenity!");
            }
          })
      } else {
        this.amenityService.addAmenity(this.amenityForm.value)
          .subscribe({
            next: (response: any) => {
              this.onClose();
              console.log(response.message);
              if (response.status === 400) alert("You have entered invalid data for the amenity!");
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
      this.amenityForm.markAllAsTouched(); // determine wether or not to display validation errors when users "touched"
    }
  }

  private formatTime(time: string): string {
    const [hours, minutes] = time.split(':');
    return `${hours}:${minutes}:00`;
  }
}
