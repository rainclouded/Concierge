describe('Amenities Integration Tests', () => {
  beforeEach(()=>{
    cy.visit('localhost:8081')
    cy.get('input')
      .type('404')
    cy.get('button').click()
    cy.get('.tile-item')
    .contains('Amenities')
    .click()
  })
  
  it('GET amenities', ()=>{
  
    cy.get('.amenity-row-item')
    .should(($items)=> {
      expect($items).to.have.length(4)
      expect($items[0]).to.contain.text('Pool')
      expect($items[1]).to.contain.text('Gym')
      expect($items[2]).to.contain.text('Breakfast')
      expect($items[3]).to.contain.text('Bar')
    })
  })
  
  afterEach(()=>{
    cy.contains('a', 'Back to Homepage')
      .click();
    cy.get('button').find('svg[data-icon="user"]')
      .click();
    cy.contains('a', 'Log Out')
      .click();
  })
});
