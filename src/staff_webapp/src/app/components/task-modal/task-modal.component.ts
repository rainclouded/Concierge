import { Component, Input, Output, EventEmitter } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ITask } from '../../models/tasks.model';

@Component({
  selector: 'app-task-modal',
  templateUrl: './task-modal.component.html',
  standalone: true,
  imports: [CommonModule]
})
export class TaskModalComponent {
  @Input() isOpen = false;
  @Input() task: ITask | null = null;
  @Output() close = new EventEmitter<void>();

  closeModal() {
    this.close.emit();
  }

  // Mock user data
  currentUser = { firstName: 'John', lastName: 'Doe' };

  // Function to assign the task to the current user
  assignTask() {
    if (this.task) {
      this.task.assignee = `${this.currentUser.firstName} ${this.currentUser.lastName}`;
      this.task.status = 'In Progress';
    }
  }

  // Function to unassign the task
  unassignTask() {
    if (this.task) {
      this.task.assignee = null;
      this.task.status = 'Pending';
    }
  }

  // Function to toggle task completion status
  toggleComplete() {
    if (this.task) {
      if (this.task.status === 'Completed') {
        this.task.status = 'In Progress';
      } else {
        this.task.status = 'Completed';
      }
    }
  }
}
