describe('HomePage Integration Tests', () => {
  it('passes', () => {
    cy.visit('localhost:8081')
    cy.get('input')
    .type('404')
    cy.get('button').click()

    cy.url().should('include', '/home')
    cy.get('.tile-item')
    .contains('Amenities')
    .click()

    cy.get('.amenity-row-item')
    .should(($items)=> {
      expect($items).to.have.length(4)
      expect($items[0]).to.contain.text('Pool')
      expect($items[1]).to.contain.text('Gym')
      expect($items[2]).to.contain.text('Breakfast')
      expect($items[3]).to.contain.text('Bar')

    })

  })
});