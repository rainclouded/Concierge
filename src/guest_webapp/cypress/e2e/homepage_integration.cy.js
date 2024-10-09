import React from 'react'; 
import HomePage from '../../src/pages/HomePage';
import { MemoryRouter } from 'react-router-dom';  // For routing context

describe('HomePage Integration Tests', () => {
  const mockAmenities = [
    {
      id: 1,
      name: 'Pool',
      description: 'Outdoor pool',
      startTime: '09:00:00',
      endTime: '21:00:00',
    },
    {
      id: 2,
      name: 'Gym',
      description: '24/7 access gym',
      startTime: '00:00:00',
      endTime: '23:59:59',
    },
  ];

  beforeEach(() => {
    cy.intercept('GET', '**/amenities', {
      statusCode: 200,
      body: {
        message: 'Amenities retrieved successfully.',
        data: mockAmenities,
      },
    }).as('getAmenities');

    cy.mount(
      <MemoryRouter>
        <HomePage />
      </MemoryRouter>
    );
  });

  it('should display Quick Service and load the amenities data', () => {
    cy.contains('Quick Service').should('be.visible');
    cy.wait('@getAmenities');
    cy.contains('Amenities').should('be.visible');
  });

  it('should navigate to the Amenities page and display data', () => {
    cy.contains('Amenities').click();
    cy.url().should('include', '/amenities');
    cy.get('table').should('exist');
    cy.get('tbody tr').should('have.length', mockAmenities.length);
  });
});
