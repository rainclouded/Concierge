// task-enums.ts

export enum TaskType {
  RoomCleaning = 'RoomCleaning',
  Maintenance = 'Maintenance',
  FoodDelivery = 'FoodDelivery',
  WakeUpCall = 'WakeUpCall',
  LaundryService = 'LaundryService',
  SpaAndMassage = 'SpaAndMassage',
}

export enum TaskStatus {
  Pending = 'Pending',
  InProgress = 'InProgress',
  Completed = 'Completed',
}

// Formatting function for TaskType values
export function formatTaskType(value: TaskType): string {
  switch (value) {
    case TaskType.RoomCleaning:
      return 'Room Cleaning';
    case TaskType.Maintenance:
      return 'Maintenance';
    case TaskType.FoodDelivery:
      return 'Food Delivery';
    case TaskType.WakeUpCall:
      return 'Wake-Up Call';
    case TaskType.LaundryService:
      return 'Laundry Service';
    case TaskType.SpaAndMassage:
      return 'Spa & Massage';
    default:
      return value;
  }
}

// Formatting function for TaskStatus values
export function formatStatus(value: TaskStatus): string {
  switch (value) {
    case TaskStatus.Pending:
      return 'Pending';
    case TaskStatus.InProgress:
      return 'In Progress';
    case TaskStatus.Completed:
      return 'Completed';
    default:
      return value;
  }
}
