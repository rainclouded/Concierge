import { Component, Input, Output, EventEmitter } from '@angular/core';
import { ITask } from '../../models/tasks.model';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { ConfirmDeletionModalComponent } from './confirm-deletion-modal.component';
import {
  TaskStatus,
  formatTaskType,
  formatStatus,
} from '../../models/task-enums';

@Component({
  selector: 'app-task-modal',
  templateUrl: './task-modal.component.html',
  standalone: true,
  imports: [FormsModule, CommonModule, ConfirmDeletionModalComponent],
})
export class TaskModalComponent {
  @Input() isOpen = false;
  @Input() task: ITask | null = null;
  @Output() close = new EventEmitter<void>();

  isEditing = false; // To toggle between edit and view mode
  editedDescription: string = ''; // For holding the description in edit mode
  isDeleteConfirmModalOpen = false; // Control for the confirm deletion modal

  // Method to format TaskTypey
  formatTaskType = formatTaskType;
  // Method to format TaskStatus
  formatTaskStatus = formatStatus;

  // Close modal
  closeModal() {
    this.isEditing = false; // Reset editing mode
    this.editedDescription = ''; // Clear edited description
    this.close.emit();
  }

  // Toggle edit mode
  toggleEdit() {
    if (this.task) {
      this.isEditing = true; // Enable editing mode
      this.editedDescription = this.task.description; // Set initial value
    }
  }

  // Save the edited description
  saveDescription() {
    if (this.task) {
      this.task.description = this.editedDescription; // Update the task description
      this.isEditing = false; // Exit edit mode
    }
  }

  // Cancel the edit and revert to original description
  cancelEdit() {
    this.isEditing = false; // Exit edit mode without saving
  }

  assignTask() {
    if (this.task) {
      this.task.assignee = 'John Doe'; // Mock user for assignment
      this.task.status = TaskStatus.InProgress;
    }
  }

  unassignTask() {
    if (this.task) {
      this.task.assignee = null;
      this.task.status = TaskStatus.Pending;
    }
  }

  toggleComplete() {
    if (this.task) {
      this.task.status =
        this.task.status === TaskStatus.Completed
          ? TaskStatus.InProgress
          : TaskStatus.Completed;
    }
  }

  // Open the confirm deletion modal
  openDeleteConfirmModal() {
    this.isDeleteConfirmModalOpen = true;
  }

  // Close the confirm deletion modal
  closeDeleteConfirmModal() {
    this.isDeleteConfirmModalOpen = false;
  }

  // Simulate task deletion
  deleteTask() {
    console.log(`Task with ID ${this.task?.id} has been deleted.`);
    this.isDeleteConfirmModalOpen = false; // Close the confirm modal
    this.close.emit(); // Also close the task details modal
  }
}
