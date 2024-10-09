import { Component, EventEmitter, Input, OnChanges, Output } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { IAmenity } from '../../models/amenity.model';
import { AmenityService } from '../../services/amenity.service';

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
    }
  }

  onSubmit() {
    if (this.amenityForm.valid) {
      if (this.data) {
        this.amenityService
          .updateAmenity(this.data.id as number, this.amenityForm.value)
          .subscribe({
            next: (response: any) => {
              this.resetAmenityForm();
              console.log(response.message);
            }
          })
      } else {
        this.amenityService.addAmenity(this.amenityForm.value)
          .subscribe({
            next: (response: any) => {
              this.resetAmenityForm();
              console.log(response.message);
            }
          });
      }
    } else {
      this.amenityForm.markAllAsTouched(); // determine wether or not to display validation errors when users "touched"
    }
  }

  resetAmenityForm() {
    this.amenityForm.reset();
    this.onClose();
  }
}
