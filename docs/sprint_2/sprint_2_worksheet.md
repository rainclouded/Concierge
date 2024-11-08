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

Profiling 

Since we use a microservice architecture, it makes sense to profile the endpoints and identify slow or problematic endpoints. However, after running the profiler, we found that all endpoints were sufficiently performant. Out of all of the services, incident reports took the longest. This does not need to be fixed as it is not an unreasonable amount of time, however we believe it is slower than the rest due to database lookups and similar actions. One way which we could speed this up is by using a clinet side cache to store some of the data in order to reduce the amount of database interaciton as well as data transferred. Each request takes less than 0.1 seconds. Here is the result of our internal profiler: 

``` 

Sending get to http://localhost:50001/accounts 

Response took 0.03427639999972598 seconds: Request OK 

  

Sending post to http://localhost:50001/accounts 

Response took 0.029192799998781993 seconds: Request OK 

  

Sending post to http://localhost:50001/accounts 

Response took 0.03012719999969704 seconds: Request OK 

  

Sending post to http://localhost:50001/accounts 

Response took 0.029902899999797228 seconds: Request OK 

  

Sending post to http://localhost:50001/accounts/login_attempt 

Response took 0.033338999999614316 seconds: Request OK 

  

Sending post to http://localhost:50001/accounts/delete 

Response took 0.029298499999640626 seconds: Request OK 

  

Sending put to http://localhost:50001/accounts/update 

Response took 0.029190200000812183 seconds: Request OK 

  

Sending get to http://localhost:8089/amenities 

Response took 0.03728699999919627 seconds: Request OK 

  

Sending get to http://localhost:8089/amenities/2 

Response took 0.028500799999164883 seconds: Request OK 

  

Sending post to http://localhost:8089/amenities 

Response took 0.014620599999034312 seconds: Request OK 

  

Sending put to http://localhost:8089/amenities/2 

Response took 0.031769200000780984 seconds: Request OK 

  

Sending delete to http://localhost:8089/amenities/1 

Response took 0.028471499999795924 seconds: Request OK 

  

Sending get to http://localhost:8081/ 

Response took 0.02457909999975527 seconds: Request OK 

  

Sending get to http://localhost:8089/incident_reports 

Response took 0.06998700000076497 seconds: Request OK 

  

Sending get to http://localhost:8089/incident_reports/1 

Response took 0.028436799999326468 seconds: Request OK 

  

Sending post to http://localhost:8089/incident_reports 

Response took 0.06165379999947618 seconds: Request OK 

  

Sending put to http://localhost:8089/incident_reports/1 

Response took 0.02909160000126576 seconds: Request OK 

  

Sending delete to http://localhost:8089/incident_reports/2 

Response took 0.0295298000000912 seconds: Request OK 

  

Sending post to http://localhost:8089/sessions 

Response took 0.024636599999212194 seconds: Request OK 

  

Sending get to http://localhost:8089/sessions/me 

Response took 0.03289279999989958 seconds: Request OK 

  

Sending get to http://localhost:8089/sessions/public-key 

Response took 0.02870240000083868 seconds: Request OK 

  

Sending get to http://localhost:8089/permissions/healthcheck 

Response took 0.02876300000025367 seconds: Request OK 

  

Sending get to http://localhost:8089/permissions 

Response took 0.03126490000067861 seconds: Request OK 

  

Sending get to http://localhost:8089/permissions/1 

Response took 0.03031129999908444 seconds: Request OK 

  

Sending post to http://localhost:8089/permissions 

Response took 0.031055100000230595 seconds: Request OK 

  

Sending get to http://localhost:8089/permission-groups 

Response took 0.03221119999943767 seconds: Request OK 

  

Sending get to http://localhost:8089/permission-groups/1 

Response took 0.030704999999215943 seconds: Request OK 

  

Sending post to http://localhost:8089/permission-groups 

Response took 0.012689599998338963 seconds: Request OK 

  

Sending patch to http://localhost:8089/permission-groups/1 

Response took 0.030088199999227072 seconds: Request OK 

  

Sending get to http://localhost:8082 

Response took 0.020607299999028328 seconds: Request OK 

  

Sending get to http://localhost:8089/tasks 

Response took 0.015648599999622093 seconds: Request OK 

  

Sending get to http://localhost:8089/tasks 

Response took 0.011821800000689109 seconds: Request OK 

  

Sending get to http://localhost:8089/tasks/1 

Response took 0.023139400000218302 seconds: Request OK 

  

Sending post to http://localhost:8089/tasks/1 

Response took 0.006332599999950617 seconds: Request OK 

  

Sending put to http://localhost:8089/tasks/1 

Response took 0.012232099999891943 seconds: Request OK 

  

Sending patch to http://localhost:8089/tasks/1 

Response took 0.023242700000992045 seconds: Request OK 

  

Sending delete to http://localhost:8089/tasks/2 

Response took 0.03469489999952202 seconds: Request OK 

``` 

Validated against chrome dev tools profiler and network call time: 
  ![](/docs/sprint_2/images/profile.png)


## Last dash:
As we approach the final sprint, there are several challenges that may arise due to both technical complexity and resource limitations. Here are the key areas of concern: 

UI Enhancements and Refactoring without Breaking Systems:

Enhancing the UI and refactoring code to improve aesthetics and usability is important, but even small changes risk breaking existing functionality or creating inconsistencies. Careful testing will be essential to ensure that visual improvements don’t compromise stability. 

 
Limited Team Resources for Golang Permissions Service:

Only one team member is proficient in Golang, which could cause delays if that person faces challenges or becomes unavailable. This dependency creates a potential bottleneck, making knowledge-sharing essential to reduce risk and ensure progress. 

 
Integrating Watchman into the CD Pipeline:

Integrating Watchman for improved file-watching may introduce compatibility issues within our CD pipeline, potentially leading to deployment delays or misconfigurations. Testing in a staging environment and thorough documentation will help mitigate these risks. 

 
Load Balancer Integration:

Configuring a load balancer is critical for scalability but can lead to session persistence issues or uneven traffic distribution if not set up correctly. Load testing and monitoring will be necessary to confirm stability and reliability. 

 
Managing Scope Creep and Hardening the System:

The final sprint might bring new ideas from the team that can lead to scope creep, risking overloading our workload and compromising stability. Staying focused on hardening the system rather than adding features will be essential to deliver a polished result within the timeline. 

 
While each of these potential issues presents unique challenges, proactive planning, thorough testing, and effective communication within the team will help mitigate these risks. By staying focused on stability, performance, and clear priorities, we can achieve a successful and sustainable final sprint. 



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

### Mykola
[My favourite code](https://github.com/rainclouded/Concierge/tree/main/src/staff_webapp/src/app/components/task-modal)
I focused on designing the Task Modal component, which provides a user-friendly and robust interface for managing individual tasks. This component combines complex functionality, robust state management, and a user-focused design, all while maintaining a simple, polished interface.
