import React from 'react'
import HomePage from './HomePage'

describe('<HomePage />', () => {
  it('renders', () => {
    // see: https://on.cypress.io/mounting-react
    cy.mount(<HomePage />)
  })
})