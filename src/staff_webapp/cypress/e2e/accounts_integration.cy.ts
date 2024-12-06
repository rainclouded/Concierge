describe('Integration Test for Accounts Page', () => {
  beforeEach(() => {
    // Set viewport and log in
    cy.viewport(1280, 720);
    cy.visit('localhost:8082/login');
    cy.get('#room-num-input').clear().type('admin');
    cy.get('#pass-code-input').clear().type('admin');
    cy.get('button').click();
    cy.url().should('include', '/dashboard/home');

    // Navigate to the Accounts page
    cy.get('.sidebar-item').contains('Accounts').click();
  });

  it('View All Accounts', () => {
    // Check that initial accounts are loaded correctly
    cy.get('table').find('tr').should('have.length.greaterThan', 1); // Includes header row

    // Check that the pagination controls exist
    cy.get('button').find('.fa-arrow-left').should('exist'); // Left arrow
    cy.get('button').find('.fa-arrow-right').should('exist'); // Right arrow
  });

  it('Search Accounts', () => {
    // Search for an account by username
    cy.get('input[placeholder="Search by username"]').type('admin{enter}');
    cy.get('table').contains('td', 'admin').should('exist');

    // Clear the search and reset
    cy.get('input[placeholder="Search by username"]').clear().type('{enter}');
    cy.get('table').find('tr').should('have.length.greaterThan', 1); // Check accounts are reloaded
  });

  it('Add New Account', () => {
    // Open "Add Account" modal
    cy.contains('Add Account').click();

    // Add a new guest account
    cy.get('select').select('Guest');
    cy.get('input[placeholder="Enter username"]').type('12345');
    cy.contains('button', 'Save').click();

    // Verify the new account appears in the table
    cy.get('table').contains('td', '12345').should('exist');

    // Cleanup: delete the newly added account
    cy.contains('td', '12345').click(); // Open modal for the new account
    cy.get('.fa-trash').click(); // Click delete icon
    cy.contains('button', 'Delete').click(); // Confirm deletion
    cy.get('table').contains('td', '12345').should('not.exist'); // Verify deletion
  });

  it('Delete Account', () => {
    // Add a new account to delete
    cy.contains('Add Account').click();
    cy.get('select').select('Guest');
    cy.get('input[placeholder="Enter username"]').type('123223');
    cy.contains('button', 'Save').click();

    // Verify the new account appears in the table
    cy.get('table').contains('td', '123223').should('exist');

    // Open modal for the new account
    cy.contains('td', '123223').click();

    // Delete the account
    cy.get('.fa-trash').click(); // Click delete icon
    cy.contains('button', 'Delete').click(); // Confirm deletion

    // Verify the account is removed from the table
    cy.get('table').contains('td', '123223').should('not.exist');
  });

  afterEach(() => {
    // Logout after each test
    cy.get('#logout').click();
  });
});
