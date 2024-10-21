export interface ITask {
    id?: number;  // Optional, can be used for database or unique identification
    title: string;
    roomNumber: string;
    description: string;
    priority: 'Low' | 'Medium' | 'High';
    status: 'Pending' | 'In Progress' | 'Completed';
    assignee: string | null;  // Name of the assignee or null if unassigned
  }
  