import { Component, EventEmitter, Output } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-add-task-modal',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './add-task-modal.component.html',
})
export class AddTaskModalComponent {
  newRoomNumber: number | null = null;
  newTypeOfService: string = '';
  newDescription: string = '';
  errorMessage: string = ''; // To display validation error messages

  @Output() close = new EventEmitter<void>();
  @Output() save = new EventEmitter<{
    roomNumber: string;
    typeOfService: string;
    description: string;
  }>();

  // Static service options for now, can be made dynamic later
  getServiceOptions(): string[] {
    return ['Cleaning', 'Maintenance', 'Room Service'];
  }

  // Validate and Save Task
  saveTask() {
    if (!this.newRoomNumber || !this.newTypeOfService || !this.newDescription) {
      this.errorMessage = 'All fields are required.';
      return;
    }

    // Emit the save event with the task data
    this.save.emit({
      roomNumber: this.newRoomNumber.toString(),
      typeOfService: this.newTypeOfService,
      description: this.newDescription,
    });

    this.resetForm();
  }

  // Close the modal
  closeModal() {
    this.close.emit();
    this.resetForm();
  }

  // Reset the form fields
  resetForm() {
    this.newRoomNumber = null;
    this.newTypeOfService = '';
    this.newDescription = '';
    this.errorMessage = '';
  }
}
