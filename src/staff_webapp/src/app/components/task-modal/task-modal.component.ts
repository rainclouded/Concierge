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
import { TaskService } from '../../services/task.service';

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
  @Output() taskDeleted = new EventEmitter<number>(); // Emit task ID when deleted
  @Output() taskUpdated = new EventEmitter<ITask>(); // Emit the updated task

  @Input() canEdit: boolean = false;
  @Input() canDelete: boolean = false;

  isEditing = false; // To toggle between edit and view mode
  editedDescription: string = ''; // For holding the description in edit mode
  isDeleteConfirmModalOpen = false; // Control for the confirm deletion modal

  // Method to format TaskTypey
  formatTaskType = formatTaskType;
  // Method to format TaskStatus
  formatTaskStatus = formatStatus;

  constructor(private taskService: TaskService) {} // Inject TaskService

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

   // Save the edited description and send update to backend
   saveDescription() {
    if (this.task) {
      const updatedTask = { ...this.task, description: this.editedDescription };

      this.taskService.updateTask(this.task.id!, updatedTask).subscribe({
        next: (response) => {
          this.task = response.data; // Update the task with response data
          this.isEditing = false; // Exit edit mode
          this.taskUpdated.emit(this.task); // Emit the updated task
          console.log('Task updated successfully:', response.data);
        },
        error: (error) => {
          console.error('Failed to update task:', error);
        },
      });
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
      this.taskService.updateTask(this.task.id!, {...this.task, status: TaskStatus.InProgress}).subscribe({
        next: (response) => {
          console.log(response)
        },
        error: (error) => {
          console.error(error)
        } 
      })
    }
  }

  unassignTask() {
    if (this.task) {
      this.task.assignee = null;
      this.task.status = TaskStatus.Pending;
      this.taskService.updateTask(this.task.id!, {...this.task, status: TaskStatus.Pending}).subscribe({
        next: (response) => {
          console.log(response)
        },
        error: (error) => {
          console.error(error)
        } 
      })
    }
  }

  toggleComplete() {
    if (this.task) {
      this.task.status =
        this.task.status === TaskStatus.Completed
          ? TaskStatus.InProgress
          : TaskStatus.Completed;
        this.taskService.updateTask(this.task.id!, {...this.task, status: this.task.status}).subscribe({
          next: (response) => {
            console.log(response)
          },
          error: (error) => {
            console.error(error)
          } 
        })
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

  // Delete task
  deleteTask() {
    if (this.task) {
      this.taskService.deleteTask(this.task.id!).subscribe({
        next: () => {
          console.log(`Task with ID ${this.task?.id} has been deleted.`);
          this.isDeleteConfirmModalOpen = false; // Close the confirm modal
          this.taskDeleted.emit(this.task!.id); // Emit the deleted task ID
          this.close.emit(); // Also close the task details modal
        },
        error: (error) => {
          console.error('Failed to delete task:', error);
        }
      });
    }
  }
}
