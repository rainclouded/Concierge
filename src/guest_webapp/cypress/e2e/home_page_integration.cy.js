describe('Home Page Integration Tests', () => {
    beforeEach(()=>{
      cy.visit('localhost:8081')
      cy.get('input')
        .type('404')
      cy.get('button').click()
    })


    it('Get Quick Service Buttons', ()=>{
        cy.url().should('include', 'home');
        cy.wait(5);
        cy.contains('button', 'Room Service').should('be.visible');
        cy.contains('button', 'Food Delivery').should('be.visible');
        cy.contains('button', 'Wake Up Call').should('be.visible');
        cy.contains('button', 'Laundry Service').should('be.visible');
        cy.contains('button', 'Spa & Massage').should('be.visible');
        cy.contains('button', 'Maintenance').should('be.visible');
    })
    
    it('Submit request buttons', ()=>{
        cy.contains('button', 'Room Cleaning').click();
        cy.contains('button', 'Submit Request').click();
    })

    afterEach(()=>{
      cy.get('button').find('svg[data-icon="user"]')
        .click();
      cy.contains('a', 'Log Out').should('exist')
        .click();
    })
  });