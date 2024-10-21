export interface ITask {
    id?: number;  
    roomNumber: string;
    typeOfService: string;
    description: string;
    timeCreated: Date;
    assignee: string | null;  // null if unassigned
    status: 'Pending' | 'In Progress' | 'Completed';
  }
  