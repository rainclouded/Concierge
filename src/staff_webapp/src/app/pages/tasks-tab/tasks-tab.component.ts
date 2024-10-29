import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ITask } from '../../models/tasks.model';
import { TaskModalComponent } from '../../components/task-modal/task-modal.component';
import { AddTaskModalComponent } from '../../components/task-modal/add-task-modal.component'; // Import the new modal
import { mockTasks } from './mock-tasks'; // mock data
import {
  TaskType,
  TaskStatus,
  formatTaskType,
  formatStatus,
} from '../../models/task-enums';
import { TaskService } from '../../services/task.service';

@Component({
  selector: 'app-tasks-tab',
  standalone: true,
  imports: [CommonModule, TaskModalComponent, AddTaskModalComponent],
  templateUrl: './tasks-tab.component.html',
})
export class TasksTabComponent {
  tasks: ITask[] = [];

  // Pagination variables
  currentPage = 1;
  tasksPerPage = 15;

  // Sorting state
  currentSortField: string = 'roomId'; // Default sorting field
  sortDirection: 'asc' | 'desc' = 'asc'; // Default sorting direction

  // Modal control
  isAddTaskModalOpen = false; // Control for the Add Task modal

  // Method to format TaskTypey
  formatTaskType = formatTaskType;
  // Method to format TaskStatus
  formatTaskStatus = formatStatus;

  constructor(private taskService: TaskService) {}

  ngOnInit(): void {
    this.fetchTasks();
  }

  fetchTasks(): void {
    this.taskService.getAllTasks().subscribe(
      (response) => {
        this.tasks = response.data;
        console.log('Fetched tasks:', this.tasks);
      },
      (error) => {
        console.error('Failed to fetch tasks:', error);
      }
    );
  }

  get totalPages(): number {
    return Math.ceil(this.tasks.length / this.tasksPerPage);
  }

  get paginatedTasks(): ITask[] {
    let sortedTasks = this.sortTasks(
      this.tasks,
      this.currentSortField,
      this.sortDirection
    );
    const start = (this.currentPage - 1) * this.tasksPerPage;
    const end = start + this.tasksPerPage;
    return sortedTasks.slice(start, end);
  }

  // Sorting logic
  sortTasks(tasks: ITask[], field: string, direction: 'asc' | 'desc'): ITask[] {
    return tasks.slice().sort((a, b) => {
      let valueA = a[field as keyof ITask];
      let valueB = b[field as keyof ITask];

      // Handle date comparison (for the createdAt field)
      if (field === 'createdAt') {
        valueA = new Date(a.createdAt).getTime();
        valueB = new Date(b.createdAt).getTime();
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
    return description.length <= maxLength
      ? description
      : description.slice(0, maxLength).trimEnd() + '...';
  }

  // Open and close View + Edit task Task modal
  openModal(task: ITask) {
    this.selectedTask = task;
    this.isModalOpen = true;
  }

  closeModal() {
    this.isModalOpen = false;
    this.selectedTask = null;
  }

  // Open and close Add Task modal
  openAddTaskModal() {
    this.isAddTaskModalOpen = true;
  }

  closeAddTaskModal() {
    this.isAddTaskModalOpen = false;
  }

  // Save a new task via TaskService
  saveNewTask(data: {
    roomId: number;
    taskType: TaskType;
    description: string;
  }) {
    const newTask: ITask = {
      roomId: data.roomId,
      taskType: data.taskType,
      description: data.description,
      status: TaskStatus.Pending,
      assignee: null,
      requesterId: 1, // CHANGE
      createdAt: new Date(),
    };

    this.taskService.addTask(newTask).subscribe({
      next: (response) => {
        this.tasks.push(response.data); // Add the newly created task to the local list
        this.closeAddTaskModal();
        console.log('Task added successfully:', response.data);
      },
      error: (error) => {
        console.error('Failed to add task:', error);
      },
    });
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

  claimUnclaimTask(task: ITask) {
    if (task.assignee) {
      task.assignee = null;
      task.status = TaskStatus.Pending;
    } else {
      task.assignee = 'Current User';
      task.status = TaskStatus.InProgress;
    }
  }
}
