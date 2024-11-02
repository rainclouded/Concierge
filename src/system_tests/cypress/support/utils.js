// Shared Utility functions between the integration tests

export const admin_login = () => {
    cy.visit('localhost:8082/login')
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
    cy.visit('localhost:8081')
    cy.get('input')
        .type('404')
    cy.get('button').click()

}

export const guest_logout = () => {
    cy.get('button').find('svg[data-icon="user"]')
        .click();
    cy.contains('a', 'Log Out').should('exist')
        .click();
}