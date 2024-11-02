import { admin_login, admin_logout, guest_login, guest_logout } from '../support/utils.js';
describe('System tests', () => {
  //System tests for the entire App



  it('Test amenities end to end', () => {

    guest_login();
    cy.get('.tile-item')
      .contains('Amenities')
      .click()
    //Verify initial amenities
    cy.get('.amenity-row-item')
      .should(($items) => {
        expect($items).to.have.length(4)
        expect($items[0]).to.contain.text('Pool')
        expect($items[1]).to.contain.text('Gym')
        expect($items[2]).to.contain.text('Breakfast')
        expect($items[3]).to.contain.text('Bar')
      })
    cy.contains('a', 'Back to Homepage')
      .click();
    guest_logout();

    admin_login();
    cy.get('.sidebar-item')
      .contains('Amenities')
      .click()
    //Add an amenity
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
      });
    admin_logout();
    //Verify on guest side
    guest_login();
    cy.get('.tile-item')
      .contains('Amenities')
      .click()

    cy.get('.amenity-row-item')
      .should(($items) => {
        expect($items).to.have.length(5)
        expect($items[0]).to.contain.text('Pool')
        expect($items[1]).to.contain.text('Gym')
        expect($items[2]).to.contain.text('Breakfast')
        expect($items[3]).to.contain.text('Bar')
        expect($items[4]).to.contain.text('testAmenity')
      })

    cy.contains('a', 'Back to Homepage')
      .click();
    guest_logout();

    admin_login();
    cy.get('.sidebar-item')
      .contains('Amenities')
      .click()
    //cleanup
    cy.contains('th', 'testAmenity')
      .parents('tr')
      .find('td')
      .contains('Delete')
      .click()

    //Delete an amenity
    cy.contains('th', 'Bar')
      .parents('tr')
      .find('td')
      .contains('Delete')
      .click()
    cy.get('.amenity-item')
      .should(($items) => {
        expect($items).to.have.length(3)
      })

    admin_logout();

    //Validate update on guest side
    guest_login();
    cy.get('.tile-item')
      .contains('Amenities')
      .click()

    cy.get('.amenity-row-item')
      .should(($items) => {
        expect($items).to.have.length(3)
        expect($items[0]).to.contain.text('Pool')
        expect($items[1]).to.contain.text('Gym')
        expect($items[2]).to.contain.text('Breakfast')

      })

    cy.contains('a', 'Back to Homepage')
      .click();
    guest_logout();

    admin_login();
    cy.get('.sidebar-item')
      .contains('Amenities')
      .click()
    //Cleanup
    cy.contains('Add Amenity').click();
    cy.get('#name').clear().type('Bar');
    cy.get('#description').clear().type('Serves alcohol and food');
    cy.get('#startTime').clear().type('02:00:00');
    cy.get('#endTime').clear().type('03:00:00');
    cy.contains('Submit').click();

    //Edit an amenity
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
      .should(($tds) => {
        expect($tds[0]).to.contain.text('This is a test');
        expect($tds[1]).to.contain.text('12:34:00 - 13:57:00');
      })

    admin_logout();

    //Validate on guest side
    guest_login();
    cy.get('.tile-item')
      .contains('Amenities')
      .click()

    cy.get('.amenity-row-item')
      .should(($items) => {
        expect($items).to.have.length(4)
        expect($items[0]).to.contain.text('Pool')
        expect($items[1]).to.contain.text('Gym')
        expect($items[2]).to.contain.text('Breakfast')
        expect($items[3]).to.contain.text('Test')
        expect($items[3]).to.contain.text('This is a test')
        expect($items[3]).to.contain.text('12:34:00')
        expect($items[3]).to.contain.text('13:57:00')

      })

    cy.contains('a', 'Back to Homepage')
      .click();
    guest_logout();
    admin_login();
    cy.get('.sidebar-item')
      .contains('Amenities')
      .click()
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
    admin_logout();
  });


})
