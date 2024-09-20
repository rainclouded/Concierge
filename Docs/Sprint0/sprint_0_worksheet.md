# Sprint 0 Worksheet 

# [**Concierge**](https://github.com/rainclouded/Concierge)

Presentation available [here](https://umanitoba.sharepoint.com/:p:/r/sites/COMP4350/Shared%20Documents/General/Sprint%200/FullPresentation.pptx?d=w097f0f3995644dca837a1eb97dde907b&csf=1&web=1&e=fyVh6n)

[Project/Kanban Board](https://github.com/users/rainclouded/projects/2)

## Vision Statement:

Concierge is an easy to use hotel administration system which allows guests to communicate with hotel operators from anywhere and provide comprehensive administration tools for hotel employees.

### More detailed summary: 

To enhance the hotel stay experience, providing a platform that allows guests to request and inquire services and amenities and for employees to manage such services and hotel operations, ensuring convenience and consolidation for all members.

 

Concierge is a platform that is catered to both hotel guests and employees. We will provide hotel management suite that provides both hotel administration and guests with a cohesive and easy to use interface for requesting and inquiring hotel services. Additionally, Concierge includes features tailored for employees, such as dashboards to create and manage incident reports, and tools for efficiently assigning and tracking tasks to ensure smooth hotel operations.

 

Concierge stands out by modernizing the hotel experience, aligning it with today’s mobile-driven world. While most of our daily needs can be easily managed through smartphones, many hotels still rely on outdated systems like landline phones for service requests and inquiries. Concierge transforms this process, making it simple and seamless for both guests and employees to request and manage services through their mobile devices. By streamlining operations and minimizing frustrations, the platform not only enhances guest satisfaction but also drives increased sales through a more efficient, user-friendly experience.


## Initial Architecture:

![Architecture Diagram](/Docs/Sprint0/images/arch.png)

### Description
The database will be MongoDB and MySQL.\
The backend will be microservices orchestrated by Kubernetes with Python (and maybe C#) webservers hosted on Azure.\
The first front end will be for desktop and will utilize Angular 18 and TypeScript.\
The second from end will be for mobile and desktop and will utilize React and JavaScript.

#### Separation of layers


We will have a 3 layer architecture all encapsulated within microservices.

All layers will be hosted on azure separated into their own microservice component.

The first layer is the UI layer which consists of our two front ends. The second layer is our business/logic layer. This layer handles the logic/processing of the communications between the differet layers. Lastly is out data layer which consists of all of our databases storing all user data.


### Why this will work well:

We chose MongoDB and MySQL as we have identified Mongo to be more suitable for some of the data we will store compared to a SQL based database while utilizing MySQL to maintain our more classical object structured data. Both of these databases integrate well with our back-end infrastructure allowing us to scale the size as needed.​For our backend, both Python and C# provide comprehensive web server operability, we will leverage the speed of C# as well as the simplicity of Python to support our end point services. We will make use of microservices hosted on azure faciliteted by Kubernetes to permit scalability and accessibility as we would like to scale the application according to the user base.  

​Lastly, to design nice-looking and full-featured front-ends we chose Angular/TypeScript and React/JavaScript. React is a lightweight and popular JavaScript library, known for its efficiency and flexibility. Its wide adoption ensures access to numerous supported libraries and tools, which can significantly reduce development time, especially for designing user interfaces. With its rich ecosystem, creating visually appealing and polished applications is made easier and faster. This will be used for the mobile/guest front end. With React’s numerous supported libraries and tools, creating visually appealing and polished applications will be a breeze. Which is why we chose this framework for our customer/mobile front end. The component-centric structure of Angular will be leveraged in order to develop a clean and modular desktop environment making use of the various cooked-in libraries which allow angular apps to look so sleek and modern while providing full feature support.

## Features

**Core features:**

|||
|--|--|
|**Functional Features**|
|1. User Story: |As a hotelier, I want a way for guests to access hotel amenities.|
|Acceptance Criteria:|The guest dashboard should be able to review all available hotel amenity information.|
|2. User Story:| As a hotel guest, I want to request hotel services to improve my stay.|
|Acceptance Criteria:|The guest dashboard should be able to request all available hotel services.|
|3. User Story:|As an employee, I want to be informed of guest requests to efficiently perform my job.|
|Acceptance Criteria:|The employee dashboard shows all appropriate requests.|
|4. User Story:| As a hotel managaer, I want to provide tasks to my employees.|
|Acceptance Criteria:|The manager dashboard should be able to create recurring and one-time employee tasks.|
|**Non-functional Features**|
|1. User Story: | As a hotel owner, I want 500 guests to be able to make 1000 requests (cumulative) without loss of access.|
|Acceptance Criteria:|500 concurrent users should be able to make 1000 cumulative requests without noticible latency (less than one second).|
|**Trimmable Features** |
|1. User Story: |As a hotel caretaker, I would like to notify guests when I am servicing their room.|
|Acceptance Criteria:|The Employee dashboard should be able to send custom notifications to specified guests.|
|2. User Story: | As a hotel guest, I would like to monitor the charges incurred for my stay.|
|Acceptance Criteria:|The guest app should be able to view a summary of pendng charges.|




## Work division:

Each team member will be appointed a '(co-)leader' of each main architectural pillar.
Database, Backend, Pipeline, Repository/Kanban, First Front-end, and Second Front-end.
The leader of each piece is not required to create everything for that piece but rather oversee it to ensure development is on time and complete. That is, they will look for missing tests, documentation, and other issues while monitoring progress on that particular area. Each member can work on any part of the software as needed, work division and tracking will be done using a kanban board to ensure balanced distruibution. The board is the source of truth, if there are any discrepancies, the board resolves them. Additionally, there is an expectation each member will develop a complete feature utilizing the entirety of the tech stack.

