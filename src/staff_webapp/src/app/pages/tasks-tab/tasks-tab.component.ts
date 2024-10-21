import { Component } from '@angular/core';
import { ITask } from '../../models/tasks.model';

@Component({
  selector: 'app-tasks-tab',
  templateUrl: './tasks-tab.component.html'
})
export class TasksTabComponent {
  tasks: ITask[] = [
    { title: 'Clean Room', roomNumber: '102', description: 'Full room cleaning requested', priority: 'High', status: 'Pending', assignee: null },
    { title: 'Deliver Towels', roomNumber: '103', description: 'Request for extra towels', priority: 'Medium', status: 'In Progress', assignee: 'John Doe' },
    { title: 'Prepare Dinner', roomNumber: '104', description: 'Vegan dinner preparation', priority: 'Low', status: 'Completed', assignee: 'Jane Smith' }
  ];

  addTask() {
    console.log('Add Task clicked');
  }

  editTask(task: ITask) {  // Use the ITask type here
    console.log('Edit Task:', task);
  }

  claimUnclaimTask(task: ITask) {  // Use the ITask type here
    if (task.assignee) {
      task.assignee = null;
      task.status = 'Pending';
    } else {
      task.assignee = 'Current User';
      task.status = 'In Progress';
    }
  }
}
