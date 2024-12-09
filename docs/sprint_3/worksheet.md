## Load Testing

[Link to our JMeter testing file](./jmeter/Concierge.jmx)

### Environment
For our load testing, we will be using JMeter, and our test cases will be the following: 

#### Scenario 1: Staff member uses all staff functionalities of Concierge 

1. Login to Concierge 

2. Get all amenities 

3. Add a new amenity 

4. Edit amenity 

5. Delete amenity 

6. Get all incident reports 

7. Add an incident report 

8. Edit incident report 

9. Delete incident report 

10. Get all tasks 

11. Add a task 

12. Edit a task 

13. Delete a task 

#### Scenario 2: Guest member uses all guest functionalities of Concierge 

1. Login 

2. Get all amenities 

3. Submit incident report 

4. Submit a task 

5. View all tasks under specified guest 

### Test Reports
#### Scenario 1: 100 staff members concurrently making 130 requests each. 
![](./images/S1100130.png)

#### Scenario 1: 500 staff members concurrently making 130 requests each. 
![](./images/S1500130.png)

#### Scenario 2: 100 guests concurrently making 100 requests each. 
![](./images/S2100100.png)

#### Scenario 2: 500 guests concurrently making 100 requests each. 
![](./images/S2500100.png)

### Bottlenecks
Some of the bottlenecks we identified are the GET and ADD endpoints for amenities, tasks and incident reports. We found as more and more of these amenities/tasks/incident reports are being added to the system concurrently, it slowed down the corresponding endpoints due to the sheer amount being added and the size to retrieve and send to the client. But as we followed a microservice architecture, these bottlenecks only impacted their corresponding service.

### Testing Our Non-Functional Requirement
Our non-functional requirement: 500 users are able to make 1000 concurrent requests, here are the results as follows:

#### 500 staff members concurrently making 1001 requests each. 
![](./images/S1NFR.png)

#### 500 guests concurrently making 1000 requests each. 
![](./images/S2NFR.png)

As shown in our test results, we could not meet our non-functional requirement. The sheer amount of concurrent users and their numerous requests proved to much for Concierge to handle. With enough money, we could maybe delegate our software to a third party service such as AWS to handle such a feat, but perhaps improvements could be made to our system to better handle such situations by improving overall efficiency. 

## Thoughts
What we would do to change our design of our project given what we know now, we would definitely use third party services to implement some of our services. As an example, authenticating user and user actions took a huge part of development time in sprint 2, and we felt like if we allocated this to a third-party service, we would allocate that time to flesh out our system more by implementing more features, such as a notification service that delegates inter communication between microserivces. By using third-party services we would also be gaining experience in how to integrate them into an existing codebase, which would be valuable in a workplace.

## Other Thoughts
After discussion with the group we thought if we went over this all over again, we would not use Flask to implement some of our microservices. This is because Flask can not be deployed due to security issues. What we would do instead is find other alternatives to Flask to implement our microservices using Python.