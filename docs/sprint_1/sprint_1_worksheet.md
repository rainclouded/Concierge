## Testing Plan

[Link to our testing plan](Concierge/docs/testing_plan.md)

## Coverage:

### Account Service
![Account Service coverage report. ](images/account-coverage.png)

### Amenities Service

### Incident Report Service

### Guest WebApp Service

### Staff WebApp Service

## Most important unit tests:

[Unit Test 1](https://github.com/rainclouded/Concierge/blob/658552c30cd19159889297f31fc4a4a50e678359/src/accounts/tests/ValidationManager_unit_test.py#L101)
 - This test verifies that new users being created meet the username and password criteria. This is important as being able to maintain the user base is a key feature of the app and the code being tested is a major contributor in managing that.
   
[Unit Test 2](https://github.com/rainclouded/Concierge/blob/658552c30cd19159889297f31fc4a4a50e678359/src/accounts/tests/AuthenticationManager_unit_test.py#L163)
 - This test verifies that correct username, password pairs are authenticated successfully. This is important as user accounts need to be secure and this ensures the integrity.

## Most important integration tests:
[Integration tests 1](https://github.com/rainclouded/Concierge/blob/main/src/guest_webapp/cypress/e2e/amenities_integration.cy.js)
- This tests the integration between the guest front end and the amenities service
[Integration test 2](https://github.com/rainclouded/Concierge/blob/main/src/staff_webapp/cypress/e2e/amenities_integration.cy.ts)
- This tests the integration between the staff front end and the amenities service


## Running team 7's project:



