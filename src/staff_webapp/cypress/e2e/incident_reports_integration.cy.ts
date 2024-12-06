describe('integration test for amenities', () => {
  //Integration test that tests the connectin between the staff ui and amenities
  const getReportCard = (category: string, name: string) => {
    return cy.contains('div', category)
      .parent()
      .contains('p', name);
  };

  beforeEach(() => {
    cy.viewport(1280, 720);
    cy.visit('localhost:8082/login')
    cy.get('#room-num-input').clear().type('admin');
    cy.get('#pass-code-input').clear().type('admin');
    cy.get('button').click()
    cy.url().should('include', '/dashboard/home')
    cy.get('.sidebar-item')
      .eq(3) // Incident reports is at index 3 of the sidebar list
      .click()

    //We cannot reset the system for each test
    //Instead we ensure initial state is okay
    getReportCard('In progress', 'Lost Property')
      .should('exist')
    getReportCard('Resolved', 'Fire Alarm Malfunction')
      .should('exist')
    getReportCard('Closed', 'Food Poisoning Incident')
      .should('exist')
  });

  it('Get and view all Reports', () => {
    //Relevent checks in the "beforeEach" clause
  });

  it('Add new report', () => {
    cy.contains('Report an incident').click();
    cy.get('#title').clear().type('testReport');
    cy.get('#description').clear().type('this is a test');
    cy.get('#severity').select('LOW');
    cy.get('#status').select('OPEN');
    cy.contains('Submit').click();

    // close toast since it blocks the button
    cy.get('.toast-close-button').should('be.visible').click();

    cy.get('.report-title').then(($reports) => {
      // Filter reports that contain "testReport" in their text
      const testReports = $reports.filter((index, element) =>
        Cypress.$(element).text().includes('testReport')
      );

      // Assert that at least one report with "testReport" exists
      expect(testReports.length).to.be.greaterThan(0, 'No reports with "testReport" found');

      // clean up
      if (testReports.length > 0) {
        cy.wrap(testReports[0])
          .closest('.w-full')
          .within(() => {
            cy.get('.delete-report').click();
          });

        cy.get('app-confirmation-dialog').should('be.visible');
        cy.get('app-confirmation-dialog')
          .contains('button', 'Confirm')
          .click();

        cy.get('.report-title')
          .should('not.contain', 'testReport');
      }
    });
  });

  it('Edit report', () => {
    getReportCard('To do', 'LOW')
      .parent()
      .children()
      .find('.edit-report')
      .click()
    cy.get('#title').clear().type('Test');
    cy.get('#description').clear().type('TEST');
    cy.get('#severity').select('HIGH');
    cy.get('#status').select('CLOSED');
    cy.contains('Submit').click();

    // close toast since it blocks the button
    cy.get('.toast-close-button').should('be.visible').click();

    getReportCard('Closed', 'TEST')
      .parent()
      .children('p')
      .should(($ps) => {
        expect($ps[0]).to.contain.text('TEST')
      })

    //Cleanup
    getReportCard('Closed', 'HIGH')
      .parent()
      .children()
      .find('.edit-report')
      .click()
    cy.get('#title').clear().type('Room Maintenance Request');
    cy.get('#description').clear().type('Guest reported a leaky faucet in Room 203.');
    cy.get('#severity').select('LOW');
    cy.get('#status').select('OPEN');
    cy.contains('Submit').click();
  });

  it('Delete report', () => {
    getReportCard('To do', 'LOW')
      .parent()
      .children()
      .find('.delete-report')
      .click()

    cy.get('app-confirmation-dialog').should('be.visible');
    cy.get('app-confirmation-dialog')
      .contains('button', 'Confirm')
      .click();

    // close toast since it blocks the button
    cy.get('.toast-close-button').should('be.visible').click();

    cy.contains('Room Maintenance Request').should('not.exist')

    //Cleanup
    cy.contains('button', 'Report an incident').click();
    cy.get('#title').clear().type('Room Maintenance Request');
    cy.get('#description').clear().type('Guest reported a leaky faucet in Room 203.');
    cy.get('#severity').select('LOW');
    cy.get('#status').select('OPEN');
    cy.contains('Submit').click();
  });

  afterEach(() => {
    cy.get('#logout').click();
  });
})
