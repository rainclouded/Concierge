import { ITask } from '../../models/tasks.model';

export const mockTasks: ITask[] = [
    {
      id: 1,
      roomNumber: '102',
      typeOfService: 'Room Cleaning',
      description: 'Guest requested a quick clean of the entire room, including dusting all surfaces, vacuuming the carpets, wiping down mirrors and windows, and changing the bed linens. Additionally, the guest asked for the trash to be emptied and the bathroom to be thoroughly sanitized. They also mentioned needing more toiletries, including shampoo, conditioner, and body lotion.',
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
    },
    {
      id: 8,
      roomNumber: '303',
      typeOfService: 'Room Service',
      description: 'Order a bottle of wine and cheese platter.',
      timeCreated: new Date('2024-10-22T14:30:00'),
      assignee: 'Alice W.',
      status: 'In Progress'
    },
    {
      id: 9,
      roomNumber: '305',
      typeOfService: 'Extra Pillows',
      description: 'Request for extra pillows.',
      timeCreated: new Date('2024-10-22T10:00:00'),
      assignee: null,
      status: 'Pending'
    },
    {
      id: 10,
      roomNumber: '402',
      typeOfService: 'Laundry Service',
      description: 'Pick up laundry by 9 AM.',
      timeCreated: new Date('2024-10-22T08:30:00'),
      assignee: 'Mark T.',
      status: 'Completed'
    },
    {
      id: 11,
      roomNumber: '408',
      typeOfService: 'Room Cleaning',
      description: 'Request for a thorough cleaning before guest arrival.',
      timeCreated: new Date('2024-10-21T15:00:00'),
      assignee: 'Mia L.',
      status: 'In Progress'
    },
    {
      id: 12,
      roomNumber: '410',
      typeOfService: 'Technical Support',
      description: 'Guest requested help with TV setup.',
      timeCreated: new Date('2024-10-21T11:30:00'),
      assignee: 'Tom H.',
      status: 'In Progress'
    },
    {
      id: 13,
      roomNumber: '412',
      typeOfService: 'Minibar Restock',
      description: 'Restock minibar in the room.',
      timeCreated: new Date('2024-10-22T13:00:00'),
      assignee: null,
      status: 'Pending'
    },
    {
      id: 14,
      roomNumber: '503',
      typeOfService: 'Room Cleaning',
      description: 'Quick cleaning requested after checkout.',
      timeCreated: new Date('2024-10-22T09:00:00'),
      assignee: 'Mary S.',
      status: 'Completed'
    },
    {
      id: 15,
      roomNumber: '505',
      typeOfService: 'Maintenance',
      description: 'AC unit not working properly.',
      timeCreated: new Date('2024-10-22T12:30:00'),
      assignee: null,
      status: 'Pending'
    },
    {
      id: 16,
      roomNumber: '507',
      typeOfService: 'Extra Blankets',
      description: 'Request for additional blankets.',
      timeCreated: new Date('2024-10-22T14:00:00'),
      assignee: 'John Doe',
      status: 'Completed'
    },
    {
      id: 17,
      roomNumber: '602',
      typeOfService: 'Room Cleaning',
      description: 'Scheduled cleaning for the evening.',
      timeCreated: new Date('2024-10-21T18:00:00'),
      assignee: 'Anna P.',
      status: 'In Progress'
    },
    {
      id: 18,
      roomNumber: '604',
      typeOfService: 'Dinner Reservation',
      description: 'Guest requested a dinner reservation at a nearby restaurant.',
      timeCreated: new Date('2024-10-21T17:00:00'),
      assignee: 'Lisa M.',
      status: 'Completed'
    },
    {
      id: 19,
      roomNumber: '606',
      typeOfService: 'Wake-up Call',
      description: 'Set wake-up call for 6 AM.',
      timeCreated: new Date('2024-10-22T06:00:00'),
      assignee: null,
      status: 'Pending'
    },
    {
      id: 20,
      roomNumber: '608',
      typeOfService: 'Room Cleaning',
      description: 'Post-checkout cleaning required.',
      timeCreated: new Date('2024-10-21T16:00:00'),
      assignee: 'Kathy B.',
      status: 'In Progress'
    },
    {
      id: 21,
      roomNumber: '702',
      typeOfService: 'Minibar Restock',
      description: 'Restock minibar in the room.',
      timeCreated: new Date('2024-10-22T11:30:00'),
      assignee: null,
      status: 'Pending'
    },
    {
      id: 22,
      roomNumber: '704',
      typeOfService: 'Food Service',
      description: 'Guest ordered a light breakfast in the room.',
      timeCreated: new Date('2024-10-22T08:00:00'),
      assignee: 'Paul A.',
      status: 'Completed'
    },
    {
      id: 23,
      roomNumber: '706',
      typeOfService: 'Laundry Service',
      description: 'Pick up laundry for express service.',
      timeCreated: new Date('2024-10-22T09:00:00'),
      assignee: 'Nina V.',
      status: 'In Progress'
    },
    {
      id: 24,
      roomNumber: '802',
      typeOfService: 'Room Cleaning',
      description: 'Deep cleaning requested for long-term guest.',
      timeCreated: new Date('2024-10-22T10:30:00'),
      assignee: 'George R.',
      status: 'In Progress'
    },
    {
      id: 25,
      roomNumber: '804',
      typeOfService: 'Room Service',
      description: 'Guest requested dinner for two in the room.',
      timeCreated: new Date('2024-10-22T19:00:00'),
      assignee: null,
      status: 'Pending'
    },
    {
      id: 26,
      roomNumber: '805',
      typeOfService: 'Room Cleaning',
      description: 'Scheduled room cleaning for tomorrow morning.',
      timeCreated: new Date('2024-10-21T07:30:00'),
      assignee: 'Anna P.',
      status: 'Pending'
    },
    {
      id: 27,
      roomNumber: '807',
      typeOfService: 'Maintenance Request',
      description: 'Fix the leaking showerhead.',
      timeCreated: new Date('2024-10-22T13:45:00'),
      assignee: null,
      status: 'Pending'
    },
    {
      id: 28,
      roomNumber: '902',
      typeOfService: 'Extra Towels',
      description: 'Request for more towels in the room.',
      timeCreated: new Date('2024-10-22T12:00:00'),
      assignee: 'Jane Smith',
      status: 'In Progress'
    },
    {
      id: 29,
      roomNumber: '904',
      typeOfService: 'Food Service',
      description: 'Guest ordered room service breakfast.',
      timeCreated: new Date('2024-10-22T07:45:00'),
      assignee: 'Mark T.',
      status: 'Completed'
    },
    {
      id: 30,
      roomNumber: '906',
      typeOfService: 'Room Cleaning',
      description: 'Regular room cleaning scheduled for tomorrow.',
      timeCreated: new Date('2024-10-22T16:00:00'),
      assignee: 'Sarah Lee',
      status: 'In Progress'
    },
    {
      id: 31,
      roomNumber: '908',
      typeOfService: 'Maintenance Request',
      description: 'Fix the air conditioning unit in the room.',
      timeCreated: new Date('2024-10-21T10:00:00'),
      assignee: null,
      status: 'Pending'
    },
    {
      id: 32,
      roomNumber: '1002',
      typeOfService: 'Food Service',
      description: 'Guest requested a late-night snack.',
      timeCreated: new Date('2024-10-21T23:00:00'),
      assignee: 'John Doe',
      status: 'Completed'
    },
    {
      id: 33,
      roomNumber: '1004',
      typeOfService: 'Room Cleaning',
      description: 'General cleaning before new guest arrives.',
      timeCreated: new Date('2024-10-22T11:00:00'),
      assignee: 'Mary S.',
      status: 'In Progress'
    },
    {
      id: 34,
      roomNumber: '1006',
      typeOfService: 'Room Service',
      description: 'Room service dinner for the guest at 8 PM.',
      timeCreated: new Date('2024-10-22T20:00:00'),
      assignee: null,
      status: 'Pending'
    },
    {
      id: 35,
      roomNumber: '1008',
      typeOfService: 'Room Cleaning',
      description: 'Post-checkout cleaning required for the room.',
      timeCreated: new Date('2024-10-22T13:00:00'),
      assignee: 'Paul A.',
      status: 'In Progress'
    }
];
