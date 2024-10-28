import { TaskType, TaskStatus } from './task-enums';

export interface ITask {
  id?: number;
  roomId: number;
  taskType: TaskType;
  description: string;
  timeCreated: Date;
  assignee: string | null; // null if unassigned
  status: TaskStatus;
}
