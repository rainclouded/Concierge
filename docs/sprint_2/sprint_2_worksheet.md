## Regression Testing


1.

So far we have been able to maintain all of our tests on every commit to both dev and main. This inclused all unit tests, integration tests, and system tests. The system tests have began to take a while to run, but not long enough to cause any issues.

We are running the regression testing through the CI/CD pipeline. We are utilizing GitHub Actions in order to run our testing. To perform each test, we are using different tools: 

 

Python - [unittest](https://github.com/python/cpython/blob/3.13/Lib/unittest/__init__.py) 

C#/ASP.NET [dotnet test](https://learn.microsoft.com/en-us/dotnet/core/tools/dotnet-test) 

For front-end integreation tests - [Cypress](https://www.cypress.io/) 

 

On pull request to main or dev, all of our tests run (integration, unit, and system) are run. We have not had to reduce the number of unit tests or integration tests over time as they are still consistently running under our metric of 10 minutes. Furthermore, during development we are constantly manually running unit tests for the components on which we work to ensure smooth development. 

2.

List of test files (corresponding images below each): 

[Build test](https://github.com/rainclouded/Concierge/blob/main/.github/workflows/docker-image.yml) 

  ![Build Test](/docs/sprint_2/images/build.png)



[Node and cypress integration and system tests](https://github.com/rainclouded/Concierge/blob/main/.github/workflows/integration_system_tests.yml) 

  ![Integration and System Tests](/docs/sprint_2/images/e2e_int_cypress.png)


[Python integration tests](https://github.com/rainclouded/Concierge/blob/main/.github/workflows/python-integration-tests.yml) 


  ![Intehration tests on python 3.10](/docs/sprint_2/images/python310_int.png)

  ![Integration tests on python 3.11](/docs/sprint_2/images/python311_int.png)


[Amenities, Accounts, and other unit tests](https://github.com/rainclouded/Concierge/blob/main/.github/workflows/run-unit-tests.yml) 



  ![Accounts on python 3.9](/docs/sprint_2/images/accounts39_unit.png)

  ![Accounts on python 3.10](/docs/sprint_2/images/accounts310_unit.png)

  ![Accounts on python 3.11](/docs/sprint_2/images/accounts_311_unit.png)

  ![Amenities on python 3.9](/docs/sprint_2/images/amenities_39_unit.png)

  ![Amenities on python 3.10](/docs/sprint_2/images/amenities_310_unit.png)

  ![Amenities on python 3.11](/docs/sprint_2/images/amenities_311_unit.png)

  ![Task system](/docs/sprint_2/images/task_unit.png)


[Task system integration tests](https://github.com/rainclouded/Concierge/blob/main/.github/workflows/task-system-postgres-integration-tests.yml) 


  ![Task system](/docs/sprint_2/images/task_system_int.png)


## Testing slowdown

As mentioned in question 1, we have been able to keep all of our testing. Currently, some of our tests which involve the UI are approaching 8 minutes, however we aim to keep each test suite under 10 minutes and thus have not needed to remove any tests. Furthermore, our turnaround time for approving and merging pull requests is about 1 day, giving us plenty of time for the test to complete before any major change is made to our main branch. 


In terms of test-plan, we have not created more than one. We run all our tests in all scenarios in order to ensure expected and desired performance. However, in future, if the integration tests begin to increase at the rate they have been, we will look into reducing the amount run during development.
## Not testing


## Profiler:



## Last dash:



## Showoff

### Leeroy
My favourite code:

https://github.com/rainclouded/Concierge/tree/main/src/amenities/amenities_server

I am most proud of the way I designed the amenities microservice and how everything was separated into their own classes and interfaces by their own special delegated tasks. This separation reduced coupling and made the system more accessible for other group members, enabling smoother implementation, easier additions, quicker bug fixes, and better overall maintenance.

### @rainclouded
[My favourite code](https://github.com/rainclouded/Concierge/blob/main/src/system_tests/api_profiling/api_profiler.py)


This is an api profiler @rainclouded wrote in a very pythonic fashion.

### Nathan (NhatAnh)

[My favourite code: Task System](https://github.com/rainclouded/Concierge/tree/main/src/task_system)

The task system is my favorite project—designing it was incredibly rewarding as I learned .NET and applied the repository pattern to create a robust service structure. I enjoyed building it with well-defined layers, which added flexibility and organization to the code, making each component more efficient and maintainable.
