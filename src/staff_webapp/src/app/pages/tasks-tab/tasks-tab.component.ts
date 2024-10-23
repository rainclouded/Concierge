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
  tasksPerPage = 15;

  // Sorting state
  currentSortField: string = 'roomNumber'; // Default sorting field
  sortDirection: 'asc' | 'desc' = 'asc';  // Default sorting direction

  get totalPages(): number {
    return Math.ceil(this.tasks.length / this.tasksPerPage);
  }

  get paginatedTasks(): ITask[] {
    let sortedTasks = this.sortTasks(this.tasks, this.currentSortField, this.sortDirection);
    const start = (this.currentPage - 1) * this.tasksPerPage;
    const end = start + this.tasksPerPage;
    return sortedTasks.slice(start, end);
  }

  // Sorting logic
  sortTasks(tasks: ITask[], field: string, direction: 'asc' | 'desc'): ITask[] {
    return tasks.slice().sort((a, b) => {
      let valueA = a[field as keyof ITask];
      let valueB = b[field as keyof ITask];
  
      // Handle date comparison (for the timeCreated field)
      if (field === 'timeCreated') {
        valueA = new Date(a.timeCreated).getTime();
        valueB = new Date(b.timeCreated).getTime();
      }

      // Convert roomNumber to number for sorting if the field is 'roomNumber'
      if (field === 'roomNumber') {
        valueA = parseInt(valueA as unknown as string, 10);
        valueB = parseInt(valueB as unknown as string, 10);
      }
  
      // Check for null or undefined values and treat them as lowest possible values
      if (valueA == null) return direction === 'asc' ? -1 : 1;
      if (valueB == null) return direction === 'asc' ? 1 : -1;
  
      // Perform the actual comparison
      if (valueA < valueB) {
        return direction === 'asc' ? -1 : 1;
      } else if (valueA > valueB) {
        return direction === 'asc' ? 1 : -1;
      } else {
        return 0;
      }
    });
  }

  toggleSort(field: string) {
    if (this.currentSortField === field) {
      // Toggle sort direction if already sorting by this field
      this.sortDirection = this.sortDirection === 'asc' ? 'desc' : 'asc';
    } else {
      // Set new sort field and default to ascending order
      this.currentSortField = field;
      this.sortDirection = 'asc';
    }
  }

  // Modal control
  isModalOpen = false;
  selectedTask: ITask | null = null;
  
  // Method to limit description length in table
  getDescriptionPreview(description: string, maxLength: number = 70): string {
    return description.length <= maxLength ? description : description.slice(0, maxLength).trimEnd() + '...';
  }

  openModal(task: ITask) {
    this.selectedTask = task;
    this.isModalOpen = true;
  }

  closeModal() {
    this.isModalOpen = false;
    this.selectedTask = null;
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
