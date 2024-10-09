describe('integration test for amenities', () => {
  //Integration test that tests the connectin between the staff ui and amenities
  
  beforeEach(()=> {
    cy.visit('localhost:8082/login')
    //When permissions are added should have to login here
    cy.get('button').click()
    cy.url().should('include', '/dashboard/home')
    cy.get('.sidebar-item')
      .contains('Amenities')
      .click()
    
    //We cannot reset the system for each test
    //Instead we ensure initial state is okay
    cy.get('.amenity-item')
      .should(($items) => {
        expect($items).to.have.length(4)
        expect($items[0]).to.contain.text('Pool')
        expect($items[1]).to.contain.text('Gym')
        expect($items[2]).to.contain.text('Breakfast')
        expect($items[3]).to.contain.text('Bar')
      })
  });
  
  it('Get and view all amenities', () => {
      cy.get('.amenity-item')
      .should(($items) => {
        expect($items).to.have.length(4)
        expect($items[0]).to.contain.text('Pool')
        expect($items[1]).to.contain.text('Gym')
        expect($items[2]).to.contain.text('Breakfast')
        expect($items[3]).to.contain.text('Bar')
      }) 
  });
  
  it('Create new amenity', () => {
    cy.contains('Add Amenity').click();
    cy.get('#name').clear().type('testAmenity');
    cy.get('#description').clear().type('this is a test');
    cy.get('#startTime').clear().type('06:00:00');
    cy.get('#endTime').clear().type('07:00:00');
    cy.contains('Submit').click();
    cy.get('.amenity-item')
      .should(($items) => {
        expect($items).to.have.length(5)
        expect($items[4]).to.contain.text('testAmenity')
      })
    //cleanup
    cy.contains('th', 'testAmenity')
      .parents('tr')
      .find('td')
      .contains('Delete')
      .click()
  });
  
  it('Delete amenity', () => {
    cy.contains('th', 'Bar')
      .parents('tr')
      .find('td')
      .contains('Delete')
      .click()
    cy.get('.amenity-item')
      .should(($items) => {
        expect($items).to.have.length(3)
      })
    //cleanup
    cy.contains('Add Amenity').click();
    cy.get('#name').clear().type('Bar');
    cy.get('#description').clear().type('Serves alcohol and food');
    cy.get('#startTime').clear().type('02:00:00');
    cy.get('#endTime').clear().type('03:00:00');
    cy.contains('Submit').click();
  });
  
  it('Edit amenity', ()=>{
    cy.contains('th', 'Bar')
      .parents('tr')
      .find('td')
      .contains('Edit')
      .click();
    cy.get('#name').clear().type('Test');
    cy.get('#description').clear().type('This is a test');
    cy.get('#startTime').clear().type('12:34:00');
    cy.get('#endTime').clear().type('13:57:00');
    cy.contains('Submit').click();
    
    cy.contains('th', 'Test')
      .parents('tr')
      .find('td')
      .should(($tds)=>{
        expect($tds[0]).to.contain.text('This is a test');
        expect($tds[1]).to.contain.text('12:34:00 - 13:57:00');
      })
      
    //cleanup
    cy.contains('th', 'Test')
      .parents('tr')
      .find('td')
      .contains('Edit')
      .click();
    cy.get('#name').clear().type('Bar');
    cy.get('#description').clear().type('Serves alcohol and food');
    cy.get('#startTime').clear().type('02:00:00');
    cy.get('#endTime').clear().type('03:00:00');
    cy.contains('Submit').click();
  })
  
  afterEach(()=>{
    cy.get('a[href="/login"]')
      .click()
  });
})
