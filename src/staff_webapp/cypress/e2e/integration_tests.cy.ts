describe('integration test for amenities', () => {
  //Integration test that tests the connectin between the staff ui and amenities
  it('passes', () => {
    cy.visit('localhost:8082/login')
    //When permissions are added should have to login here
    cy.get('button').click()

    cy.url().should('include', '/dashboard/home')
    cy.get('.sidebar-item')
      .contains('Amenities')
      .click()

      cy.get('.amenity-item')
      .should(($items) => {
        expect($items).to.have.length(4)
        expect($items[0]).to.contain.text('Pool')
        expect($items[1]).to.contain.text('Gym')
        expect($items[2]).to.contain.text('Breakfast')
        expect($items[3]).to.contain.text('Bar')

      }) 
  })
})