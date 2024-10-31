import { TaskType, TaskStatus } from './task-enums';

export interface ITask {
  assigneeId?: number;
  createdAt: Date;
  description: string;
  id?: number;
  requesterId?: number,
  roomId: number;
  status: TaskStatus;
  taskType: TaskType;
  assignee: string | null; // null if unassigned
}
