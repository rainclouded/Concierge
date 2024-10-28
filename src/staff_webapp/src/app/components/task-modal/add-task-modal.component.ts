import { Component, EventEmitter, Output } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { TaskType, formatTaskType } from '../../models/task-enums';

@Component({
  selector: 'app-add-task-modal',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './add-task-modal.component.html',
})
export class AddTaskModalComponent {
  newRoomNumber: number | null = null;
  newTaskType: TaskType | undefined;
  newDescription: string = '';
  errorMessage: string = ''; // To display validation error messages

  @Output() close = new EventEmitter<void>();
  @Output() save = new EventEmitter<{
    roomId: number;
    taskType: TaskType;
    description: string;
  }>();

  // Static service options for now, can be made dynamic later
  getServiceOptions(): { value: TaskType; label: string }[] {
    return Object.values(TaskType).map((taskType) => ({
      value: taskType as TaskType,
      label: formatTaskType(taskType as TaskType),
    }));
  }

  // Validate and Save Task
  saveTask() {
    if (!this.newRoomNumber || !this.newTaskType || !this.newDescription) {
      this.errorMessage = 'All fields are required.';
      return;
    }

    // Emit the save event with the task data
    this.save.emit({
      roomId: this.newRoomNumber,
      taskType: this.newTaskType,
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
    this.newTaskType = undefined;
    this.newDescription = '';
    this.errorMessage = '';
  }
}
