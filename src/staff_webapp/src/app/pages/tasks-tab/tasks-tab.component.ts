import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ITask } from '../../models/tasks.model';
import { TaskModalComponent } from '../../components/task-modal/task-modal.component';
import { mockTasks } from './mock-tasks';  // mock data

@Component({
  selector: 'app-tasks-tab',
  standalone: true,
  imports: [CommonModule, TaskModalComponent],  
  templateUrl: './tasks-tab.component.html'
})
export class TasksTabComponent {
  tasks: ITask[] = mockTasks;

  // Pagination variables
  currentPage = 1;
  tasksPerPage = 10; // Show 10 tasks per page

  get totalPages(): number {
    return Math.ceil(this.tasks.length / this.tasksPerPage);
  }

  get paginatedTasks(): ITask[] {
    const start = (this.currentPage - 1) * this.tasksPerPage;
    const end = start + this.tasksPerPage;
    return this.tasks.slice(start, end);
  }

  // Modal control
  isModalOpen = false;
  selectedTask: ITask | null = null;

  // Method to limit description length in table
  getDescriptionPreview(description: string, maxLength: number = 100): string {
    if (description.length <= maxLength) {
      return description;  // Return full description if it's short enough
    }
    // Trim trailing spaces
    let trimmedDescription = description.slice(0, maxLength).trimEnd();  
    
    // Check if the last character is a punctuation mark and remove it
    const lastChar = trimmedDescription.charAt(trimmedDescription.length - 1);
    if (['.', ',', ';', ':'].includes(lastChar)) {
      trimmedDescription = trimmedDescription.slice(0, -1);  // Remove the punctuation
    }
  
    return trimmedDescription + '...';
  }

  openModal(task: ITask) {
    this.selectedTask = task;
    this.isModalOpen = true;
  }

  closeModal() {
    this.isModalOpen = false;
  }

  // Pagination controls
  goToPreviousPage() {
    if (this.currentPage > 1) {
      this.currentPage--;
    }
  }

  goToNextPage() {
    if (this.currentPage < this.totalPages) {
      this.currentPage++;
    }
  }

  // Task Actions
  addTask() {
    console.log('Add Task clicked');
  }

  editTask(task: ITask) { 
    console.log('Edit Task:', task);
  }

  claimUnclaimTask(task: ITask) { 
    if (task.assignee) {
      task.assignee = null;
      task.status = 'Pending';
    } else {
      task.assignee = 'Current User';
      task.status = 'In Progress';
    }
  }
}
