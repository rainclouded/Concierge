// Shared Utility functions between the integration tests

export const admin_login = () => {
    cy.viewport(1280, 720);
    cy.visit('localhost:8089/staff/login')
    cy.get('#room-num-input').clear().type('admin');
    cy.get('#pass-code-input').clear().type('admin');
    cy.get('button').click()
    cy.url().should('include', '/dashboard/home')

}

export const admin_logout = () => {
    cy.get('#logout')
        .click()
}

export const guest_login = () => {
    cy.viewport(390, 844);
    cy.visit("localhost:8089");
    cy.get("#roomNum").type("11111");
    cy.get("#roomKey").type("password");
    cy.get("button").click();
}

export const guest_logout = () => {
    cy.get('button').find('svg[data-icon="user"]')
        .click();
    cy.contains('a', 'Log Out').should('exist')
        .click();
}

export const getReportCard = (category, name) => {
    return cy.contains('div', category)
      .parent()
      .contains('p', name);
  };