describe('Integration test for Home page', () => {
	// Integration test that tests the connection between the staff UI and home

	beforeEach(() => {
		cy.viewport(1280, 720);
		cy.visit('localhost:8089/staff/login');
		cy.get('#room-num-input').clear().type('admin');
		cy.get('#pass-code-input').clear().type('admin');
		cy.get('button').click();
		cy.url().should('include', '/dashboard/home');
		cy.get('.sidebar-item')
			.eq(0) // Home is at index 0 of the sidebar list
			.click();
	});

	it('View amenities', () => {
		cy.get('.left-upper-section')
			.children('h2')
			.should('contain', 'Amenities');

		cy.get('.left-upper-section table')
			.find('thead tr')
			.children()
			.should(($cols) => {
				expect($cols).to.have.length(2);
				expect($cols.eq(0)).to.contain.text('Title');
				expect($cols.eq(1)).to.contain.text('Time');
			});

		const expectedAmenities = ['Pool', 'Gym', 'Breakfast', 'Bar']

		cy.get('th.whitespace-nowrap')
			.should(($items) => {
				expect($items).to.have.length(expectedAmenities.length)
				expectedAmenities.forEach((amenity, index) => {
					expect($items[index]).to.contain.text(amenity)
				})
			})
	});

	it('View incident reports', () => {
		cy.get('.right-upper-section')
			.children('h2')
			.should('contain', 'Incident Reports');

		// Verify the counts in the "Incident Reports" subsections
		const reportSections = ['To Do', 'In Progress', 'Resolved', 'Closed'];
		reportSections.forEach((section) => {
			cy.get('h3.text-xl.font-semibold.uppercase').contains(section).should('be.visible');
			cy.get('h3')
				.contains(section)
				.parents('.bg-lightPrimary')
				.within(() => {
					cy.get('span.text-5xl.font-bold').should('not.be.empty');
				});
		});
	});

	it('View unclaimed tasks', () => {
		cy.get('.bottom-section')
			.children('h2')
			.should('contain', 'Recent Unclaimed Tasks');

		cy.get('.bottom-section table')
			.find('thead tr')
			.children()
			.should(($cols) => {
				expect($cols).to.have.length(4);
				expect($cols.eq(0)).to.contain.text('Room No');
				expect($cols.eq(1)).to.contain.text('Type of Service');
				expect($cols.eq(2)).to.contain.text('Description');
				expect($cols.eq(3)).to.contain.text('Time Created');
			});

		cy.get('.bottom-section table')
			.find('tbody tr')
			.each(($row) => {
				cy.wrap($row).within(() => {
					// Ensure all table cells in the row are not empty
					cy.get('td').each(($cell) => {
						cy.wrap($cell).should('not.be.empty');
					});
				});
			});
	});

	afterEach(() => {
		cy.get('#logout').click();
	});
});
