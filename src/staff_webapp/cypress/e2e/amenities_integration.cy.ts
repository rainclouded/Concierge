describe('integration test for amenities', () => {
  //Integration test that tests the connectin between the staff ui and amenities

  beforeEach(() => {
    cy.visit('localhost:8089/staff/login');
    cy.get('#room-num-input').clear().type('admin');
    cy.get('#pass-code-input').clear().type('admin');
    cy.get('button').click()
    cy.url().should('include', '/dashboard/home')
    cy.get('.sidebar-item')
      .eq(2) // Amenities is at index 2 of the sidebar list
      .click()

    //We cannot reset the system for each test
    //Instead we ensure initial state is okay
    const expectedAmenities = ['Pool', 'Gym', 'Breakfast', 'Bar']

    cy.get('.amenity-item')
      .should(($items) => {
        expect($items).to.have.length(expectedAmenities.length)
        expectedAmenities.forEach((amenity, index) => {
          expect($items[index]).to.contain.text(amenity)
        })
      })
  });

  it('Get and view all amenities', () => {
    const expectedAmenities = ['Pool', 'Gym', 'Breakfast', 'Bar']

    cy.get('.amenity-item')
      .should(($items) => {
        expect($items).to.have.length(expectedAmenities.length)
        expectedAmenities.forEach((amenity, index) => {
          expect($items[index]).to.contain.text(amenity)
        })
      })
  });

  it('Create new amenity', () => {
    cy.contains('Add Amenity').click();
    cy.get('#name').clear().type('testAmenity');
    cy.get('#description').clear().type('this is a test');
    cy.get('#startTime').clear().type('18:00');
    cy.get('#endTime').clear().type('19:00');
    cy.contains('Submit').click();

    cy.get('.amenity-item')
      .should(($items) => {
        expect($items).to.have.length(5)
        const hasTestAmenity = Cypress.$.makeArray($items)
          .some(item => Cypress.$(item).text().includes('testAmenity'))
        expect(hasTestAmenity).to.be.true
      })
    
    //cleanup
    cy.contains('th', 'testAmenity')
      .parents('tr')
      .find('td')
      .find('.delete-amenity')
      .click()

    cy.get('app-confirmation-dialog').should('be.visible');
    cy.get('app-confirmation-dialog')
      .contains('button', 'Confirm')
      .click();
  });

  it('Delete amenity', () => {
    cy.contains('th', 'Bar')
      .parents('tr')
      .find('td')
      .find('.delete-amenity')
      .click()

    cy.get('app-confirmation-dialog').should('be.visible');
    cy.get('app-confirmation-dialog')
      .contains('button', 'Confirm')
      .click();

    cy.get('.amenity-item')
      .should(($items) => {
        expect($items).to.have.length(3)
      })
    //cleanup
    cy.contains('Add Amenity').click();
    cy.get('#name').clear().type('Bar');
    cy.get('#description').clear().type('Serves alcohol and food');
    cy.get('#startTime').clear().type('14:00');
    cy.get('#endTime').clear().type('15:00');
    cy.contains('Submit').click();
  });

  it('Edit amenity', () => {
    cy.contains('th', 'Bar')
      .parents('tr')
      .find('td')
      .find('.edit-amenity')
      .click();
    cy.get('#name').clear().type('Test');
    cy.get('#description').clear().type('This is a test');
    cy.get('#startTime').clear().type('12:34');
    cy.get('#endTime').clear().type('13:57');
    cy.contains('Submit').click();

    cy.contains('th', 'Test')
      .parents('tr')
      .find('td')
      .should(($tds) => {
        expect($tds[0]).to.contain.text('This is a test');
        expect($tds[1]).to.contain.text('12:34 PM - 01:57 PM');
      })

    // close toast since it blocks the button
    cy.get('.toast-close-button').should('be.visible').click();

    //cleanup
    cy.contains('th', 'Test')
      .parents('tr')
      .find('td')
      .find('.edit-amenity')
      .click();
    cy.get('#name').clear().type('Bar');
    cy.get('#description').clear().type('Serves alcohol and food');
    cy.get('#startTime').clear().type('14:00');
    cy.get('#endTime').clear().type('15:00');
    cy.contains('Submit').click();
  })

  afterEach(() => {
    cy.get('#logout')
      .click()
  });
})
