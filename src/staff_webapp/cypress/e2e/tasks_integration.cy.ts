describe('Integration Test for Task Manager', () => {
  // Utility function to get a task by its room number and description
  const getTaskRow = (roomNumber: number, description: string) => {
    return cy
      .contains('td', roomNumber.toString())
      .parent()
      .contains('td', description);
  };

  beforeEach(() => {
    // Set viewport and log in
    cy.viewport(1280, 720);
    cy.visit('localhost:8089/staff/login');
    cy.get('#room-num-input').clear().type('admin');
    cy.get('#pass-code-input').clear().type('admin');
    cy.get('button').click();
    cy.url().should('include', '/dashboard/home');

    // Navigate to the Task Manager page
    cy.get('.sidebar-item').contains('Tasks').click();
  });

  it('View All Tasks', () => {
    // Check that the table is present
    cy.get('table').should('exist');

    // Check that the pagination controls exist
    cy.get('button').find('.fa-arrow-left').should('exist'); // Left arrow
    cy.get('button').find('.fa-arrow-right').should('exist'); // Right arrow
  });

  it('Add New Task', () => {
    // Open "Add Task" modal
    cy.contains('Add Task').click();

    // Validation check for required fields
    cy.contains('Save').click();
    cy.contains('All fields are required').should('be.visible');

    // Enter task details and submit
    cy.get('input[placeholder="Enter room number"]').type('202');

    // Set the value of the dropdown using JavaScript
    cy.get('#taskType').then(($select) => {
      const selectElement = $select[0] as HTMLSelectElement;
      const targetOption = Array.from(selectElement.options).find(
        (option) => option.text === 'Maintenance'
      );
      if (targetOption) {
        targetOption.selected = true;
        selectElement.dispatchEvent(new Event('change', { bubbles: true }));
      }
    });

    // Enter the task description
    cy.get('textarea[placeholder="Enter task description"]').type(
      'Temporary task for test.'
    );

    // Verify that the fields contain the correct values before saving
    cy.get('input[placeholder="Enter room number"]').should(
      'have.value',
      '202'
    );

    cy.get('textarea[placeholder="Enter task description"]').should(
      'have.value',
      'Temporary task for test.'
    );

    // Click the Save button to add the task
    cy.contains('Save').click();

    // Verify the new task appears in the task list
    cy.contains('td', '202')
      .parent()
      .contains('td', '202')
      .should('exist');
  });

  it('Edit Task', () => {
    // Open task details by clicking on the row containing the task
    getTaskRow(
      202,
      'Temporary task for test.'
    ).click(); // Opens the modal for this task

    // Click on the edit icon
    cy.get('.fa-edit').parent('button').click();

    // Update task description in edit mode
    cy.get('textarea.w-full.mt-1.p-2.border.rounded.h-32')
      .clear()
      .type('This is an edited description.');

    // Save the edited description
    cy.contains('button', 'Save').click();

    // Manually close the modal by clicking the close button (×)
    cy.contains('button', '×').click();

    // Ensure modal is closed by checking that the modal backdrop no longer exists
    cy.get('.fixed.z-10.inset-0.bg-black.bg-opacity-50').should('not.exist');

    // Verify the task list shows the updated description
    getTaskRow(202, 'This is an edited description.').should('exist');

    // Cleanup: revert task to original description
    getTaskRow(202, 'This is an edited description.').click();

    // Click on the edit icon again to revert the description
    cy.get('.fa-edit').parent('button').click();
    cy.get('textarea.w-full.mt-1.p-2.border.rounded.h-32')
      .clear()
      .type('Temporary task for test.');

    // Save the original description
    cy.contains('button', 'Save').click();

    // Manually close the modal again after reverting the description
    cy.contains('button', '×').click();

    // Ensure modal is closed after cleanup
    cy.get('.fixed.z-10.inset-0.bg-black.bg-opacity-50').should('not.exist');
  });


  it('Assign and Unassign Task', () => {
    // Open the task modal by clicking on the task row
    getTaskRow(
      202,
      'Temporary task for test.'
    ).click();

    // Click "Assign to Me" to assign the task
    cy.contains('button', 'Assign to Me').click();

    // Verify the assignment by checking for "John Doe" in the Assignee field
    cy.contains('button', 'Unassign').should('be.visible'); // The button should change to "Unassign"
    cy.contains('p', 'Assignee:').invoke('text').should('include', 'John Doe'); // Check Assignee text

    // Click "Unassign" to revert the assignment
    cy.contains('button', 'Unassign').click();

    // Verify the task is unassigned
    cy.contains('p', 'Assignee:')
      .invoke('text')
      .should('include', 'Unassigned'); // Assignee should revert to "Unassigned"

    // Close the modal by clicking the close button (×)
    cy.get('button').contains('×').click();

    // Verify the modal is closed
    cy.get('.fixed.z-10.inset-0.bg-black.bg-opacity-50').should('not.exist');
  });

  it('Mark Task as Completed', () => {
    // Open task details by clicking on the task row
    getTaskRow(
      202,
      'Temporary task for test.'
    ).click();

    // Click "Assign to Me" to assign the task if it's not already assigned
    cy.contains('button', 'Assign to Me').click();

    // Mark the task as complete
    cy.contains('button', 'Resolved').click();

    // Verify the status change to "Completed"
    cy.contains('p', 'Status:').invoke('text').should('include', 'Completed');

    // Cleanup: Revert the status to "In Progress"
    cy.contains('button', 'Completed').click(); // Clicking again should toggle back to "In Progress"

    // Verify the status reverted to "In Progress"
    cy.contains('p', 'Status:').invoke('text').should('include', 'In Progress');

    // Click "Unassign" to unassign the task
    cy.contains('button', 'Unassign').click();

    // Ensure task is unassigned by checking if "Assignee:" text contains "Unassigned"
    cy.contains('p', 'Assignee:')
      .invoke('text')
      .should('include', 'Unassigned');

    // Close the modal by clicking the close button (×)
    cy.get('button').contains('×').click();

    // Verify the modal is closed
    cy.get('.fixed.z-10.inset-0.bg-black.bg-opacity-50').should('not.exist');
  });

  it('Delete Task', () => {
    // Open task details by clicking on the row containing the new task
    getTaskRow(202, 'Temporary task for test.').click();

    // Click on the delete icon
    cy.get('.fa-trash').parent('button').click();

    // Confirm the deletion in the confirmation modal
    cy.contains('button', 'Confirm').click();

    // Verify both modals are closed after deletion
    cy.get('.fixed.z-10.inset-0.bg-black.bg-opacity-50').should('not.exist');

    // Verify the task no longer exists in the task list
    cy.contains('Temporary task for test.').should('not.exist');
  });

  afterEach(() => {
    // Logout after each test
    cy.get('#logout').click();
  });
});
