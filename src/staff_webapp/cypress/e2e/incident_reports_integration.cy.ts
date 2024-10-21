describe('integration test for amenities', () => {
  //Integration test that tests the connectin between the staff ui and amenities
  const getReportCard = (category: string, name: string) => {
    return cy.contains('div', category)
      .parent()
      .contains('p', name);
  };

  beforeEach(()=> {
    cy.viewport(1280, 720);
    cy.visit('localhost:8082/login')
    cy.get('#room-num-input').clear().type('admin');
    cy.get('#pass-code-input').clear().type('admin');
    cy.get('button').click()
    cy.url().should('include', '/dashboard/home')
    cy.get('.sidebar-item')
      .contains('Incident Reports')
      .click()

    //We cannot reset the system for each test
    //Instead we ensure initial state is okay
    getReportCard('To do', 'Room Maintenance Request')
      .should('exist')
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

  it('Edit', () => {
    getReportCard('To do', 'Room Maintenance Request')
      .parent()
      .children()
      .contains('Edit')
      .click()
    cy.get('#title').clear().type('Test');
    cy.get('#description').clear().type('TEST');
    cy.get('#severity').clear().type('HIGH');
    cy.get('#status').clear().type('CLOSED');
    cy.contains('Submit').click();

    getReportCard('Closed', 'TEST')
      .parent()
      .children('p')
      .should(($ps)=>{
        expect($ps[0]).to.contain.text('TEST')
        expect($ps[1]).to.contain.text('HIGH')
      })

    //Cleanup
    getReportCard('Closed', 'TEST')
      .parent()
      .children()
      .contains('Edit')
      .click()
    cy.get('#title').clear().type('Room Maintenance Request');
    cy.get('#description').clear().type('Guest reported a leaky faucet in Room 203.');
    cy.get('#severity').clear().type('LOW');
    cy.get('#status').clear().type('OPEN');
    cy.contains('Submit').click();
  });

  it('Delete', () => {
    getReportCard('To do', 'Room Maintenance Request')
      .parent()
      .children()
      .contains('Delete')
      .click()

    cy.contains('Room Maintenance Request').should('not.exist')

    //Cleanup
    cy.contains('button', 'Report an incident').click();
    cy.get('#title').clear().type('Room Maintenance Request');
    cy.get('#description').clear().type('Guest reported a leaky faucet in Room 203.');
    cy.get('#severity').clear().type('LOW');
    cy.get('#status').clear().type('OPEN');
    cy.contains('Submit').click();
  });

  afterEach(()=>{
    cy.get('#logout').click();
  });
 })
