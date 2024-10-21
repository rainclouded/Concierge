import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ITask } from '../../models/tasks.model';
import { TaskModalComponent } from '../../components/task-modal/task-modal.component';

@Component({
  selector: 'app-tasks-tab',
  standalone: true,
  imports: [CommonModule, TaskModalComponent],  
  templateUrl: './tasks-tab.component.html'
})
export class TasksTabComponent {
  tasks: ITask[] = [
    {
      id: 1,
      roomNumber: '102',
      typeOfService: 'Room Cleaning',
      description: 'Guest requested a quick clean of the entire room.',
      timeCreated: new Date('2024-10-21T08:00:00'),
      assignee: null,
      status: 'Pending'
    },
    {
      id: 2,
      roomNumber: '104',
      typeOfService: 'Extra Towels',
      description: 'Request for additional towels for the room.',
      timeCreated: new Date('2024-10-20T09:30:00'),
      assignee: 'John Doe',
      status: 'Completed'
    },
    {
      id: 3,
      roomNumber: '104',
      typeOfService: 'Food Service',
      description: 'Dinner requested to be served at 7 PM in the room.',
      timeCreated: new Date('2024-10-21T11:00:00'),
      assignee: 'Jane Smith',
      status: 'In Progress'
    },
    {
      id: 4,
      roomNumber: '106',
      typeOfService: 'Maintenance Request',
      description: 'Issue with bathroom faucet leaking since yesterday.',
      timeCreated: new Date('2024-10-20T15:45:00'),
      assignee: null,
      status: 'Pending'
    },
    {
      id: 5,
      roomNumber: '204',
      typeOfService: 'Room Cleaning',
      description: 'Deep clean requested for the carpets and furniture.',
      timeCreated: new Date('2024-10-21T09:15:00'),
      assignee: 'Sarah Lee',
      status: 'In Progress'
    },
    {
      id: 6,
      roomNumber: '207',
      typeOfService: 'Food Service',
      description: 'Guest requested a vegan meal to be served in the room.',
      timeCreated: new Date('2024-10-21T10:30:00'),
      assignee: 'Alex K.',
      status: 'In Progress'
    },
    {
      id: 7,
      roomNumber: '207',
      typeOfService: 'Wake-up Call',
      description: 'Set wake-up call for 7 AM, guest wants to be notified.',
      timeCreated: new Date('2024-10-21T08:30:00'),
      assignee: null,
      status: 'Pending'
    }
  ];

  isModalOpen = false;
  selectedTask: ITask | null = null;  // Store selected task to pass to modal

  openModal(task: ITask) {
    console.log('Task clicked:', task); 
    this.selectedTask = task;  // Set the selected task when opening modal
    this.isModalOpen = true;
  }

  closeModal() {
    this.isModalOpen = false;
  }

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
