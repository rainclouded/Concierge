describe('HomePage Integration Tests', () => {
  beforeEach(()=>{
    cy.visit('localhost:8081')
    cy.get('input')
      .type('404')
    cy.get('button').click()
    cy.get('.tile-item')
    .contains('Incident Report')
    .click()
  })
  
  it('POST incident repot', ()=>{
    cy.get('form')
      .find('input')
      .clear()
      .type('TestingIncidentCreation')
    cy.get('form')
      .find('textarea')
      .clear()
      .type('TestingIncidentCreationDescription')
    cy.contains('button', 'Submit').click()
    cy.contains('p', 'Incident report submitted successfully').should('exist');
  })
  
  afterEach(()=>{
    cy.contains('button', 'Back')
      .click();
    cy.get('button').find('svg[data-icon="user"]')
      .click();
    cy.contains('a', 'Log Out')
      .click();
  })
});
 