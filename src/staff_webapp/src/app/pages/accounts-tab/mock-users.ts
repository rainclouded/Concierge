export const mockUsers = [
  {
    username: '123456',
    type: 'guest',
    password: '12345678',
  },
  {
    username: '654321',
    type: 'guest',
    password: '87654321',
  },
  {
    username: '987654',
    type: 'guest',
    password: '45678912',
  },

  // Staff (alphanumeric username > 5 characters, any password)
  {
    username: 'staff123',
    type: 'staff',
    password: 'password123',
  },
  {
    username: 'admin2023',
    type: 'staff',
    password: 'securePass2023!',
  },
  {
    username: 'tech99',
    type: 'staff',
    password: 'techSupport42',
  },
  {
    username: 'manager007',
    type: 'staff',
    password: 'topSecret!',
  },
];
