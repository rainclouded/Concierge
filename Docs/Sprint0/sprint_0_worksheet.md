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

![Diagram](/Docs/Sprint0/block_diagram.jpg)


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

### Feature: Task System

| **Key Features** | **User Story** | **Acceptance Criteria**  |
|-------------------|-------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Create Tasks (Guest) | As a hotel guest, I want to create service requests (e.g., room service, extra towels) so that my needs are promptly addressed. | **Scenario:** Guest creates a service request. <br> Given I’m a guest, when I open the guest task dashboard and choose to create a new service request, I can enter the request details (type of service, description, priority). When I submit the request, it appears in my task list with a status of "pending," and the relevant staff members are notified. |
| View Task Status (Guest) | As a hotel guest, I want to view the status of my service requests so I can know when my needs will be met. | **Scenario:** Guest views the status of service requests. <br> Given I’m a guest, when I open the guest task dashboard, I can see a list of my active requests with their status ("pending," "in progress," "completed"). |
| Incomplete Tasks (Staff) | As a hotel staff member, I want to see a list of all incomplete guest requests so I can prioritize my work. | **Scenario:** Staff views a list of incomplete tasks. <br> Given I’m a staff member, when I open the staff task dashboard, I can see a list of all incomplete tasks with details such as request type, room number, and priority. |
| Claim a Task (Staff) | As a hotel staff member, I want to claim a task so that I can take responsibility for completing it. | **Scenario:** Staff claims a task. <br> Given I’m a staff member, when I open the staff task dashboard and click to claim a task, I am assigned as the responsible staff member, and the task status changes to "in progress." |
| Mark Task Complete (Staff) | As a hotel staff member, I want to mark a task as complete once it's finished so that I can keep my work queue up to date. | **Scenario:** Staff completes a task. <br> Given I’m a staff member, when I view an existing task I have completed on the staff dashboard, I can click the “Completed” button, and the system updates the task in real-time, giving a confirmation alert. |

---

### Feature: Accounts

| **Key Features** | **User Story** | **Acceptance Criteria** |
|-------------------------------|-------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Login (Staff & Guest) | As a user, I need to be able to access the tools relevant to me so that I can make use of the system. | **Scenario:** User logs in. <br> Given I’m a user with valid credentials, when I enter my username and password and click the "Login" button, I am directed to my role-specific dashboard (guest, staff, or manager). |
| Creation (Staff & Room Accounts) | As a hotel manager, I need to be able to give access to guests and employees so they can do their work. | **Scenario:** Creating a new account. <br> Given I’m a staff member with permission to create accounts, when I fill in the necessary fields for user creation and submit, a new account is created. |
| Permissions | As a manager, I need to be able to restrict access to various tools so that my services remain secure. | **Scenario:** Editing permissions. <br> Given I’m a logged-in staff member with permission to edit permissions, when I create a new group, configure its permissions, and assign accounts to it, those accounts gain access to those tools. |
| Settings | As a user, I want to be able to tailor the tools to my needs so that I can engage with the system as I prefer. | **Scenario:** Modifying preferences. <br> Given I’m a logged-in staff member, when I update settings (e.g., default dashboard views, notification preferences, initial landing pages), the system saves the changes and reflects them in my next login. |

---

### Feature: Amenities

| **Key Features** | **User Story** | **Acceptance Criteria** |
|-----------------------------|----------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| View | As a guest, I want to access hotel amenity information so I can find out when amenities are available (e.g., the hotel swimming pool is open). | **Scenario:** Guest views hotel amenity information. <br> Given I’m a guest, when I open the amenities dashboard, I can see a list of amenities offered by the hotel, with operating hours, descriptions, and titles. |
| Create, Update, Delete | As a staff member, I want to either create, delete, or update hotel amenities so guests are always up to date with our amenity information. | **Scenario:** Staff updates an amenity. <br> Given I’m a staff member, when I update the amenity details on the staff dashboard and click "Save," the system updates the report in real-time and displays a confirmation message. |

---

### Feature: Incident Reports

| **Key Features** | **User Story** | **Acceptance Criteria** |
|-----------------------------|----------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| View | As a hotel manager, I want a way to view past and current incident reports created by my staff to better monitor hotel incidents. | **Scenario:** Hotel manager views past and current incident reports. <br> Given I’m a hotel manager, when I access the incident reports dashboard, I can see a list of past and current incident reports, including details like severity, status, and description. |
| Update | As a hotel manager, I want a way to update incident report details as their related case progresses in real-time. | **Scenario:** Hotel manager updates an incident report. <br> Given I’m a hotel manager, when I update incident details and click “Save,” the system updates the report in real-time and shows a confirmation message. |

---

### Feature: Business Analytics (TRIMMABLE)

| **Key Features** | **User Story** | **Acceptance Criteria** |
|-----------------------------|----------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| View Guest Charges | As a staff member, I want to view and analyze guest charges during their stay so I can ensure accurate billing and identify trends. | **Scenario:** Viewing guest charges. <br> Given I’m a staff member, when I access the analytics dashboard, I can see a list of guest charges, their details, and any relevant trends over time. |

---

### Feature: Booking Service (TRIMMABLE)

| **Key Features** | **User Story** | **Acceptance Criteria** |
|-----------------------------|----------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| See Available Dates, Rooms & Prices | As a guest, I want to check room availability and view pricing details so I can plan and budget my hotel stay. | **Scenario:** Checking room availability and pricing. <br> Given I’m a guest, when I enter my check-in and check-out dates, I can see a list of available rooms, their prices, and any special offers. |




## Work division:

Each team member will be appointed a '(co-)leader' of each main architectural pillar.
Database, Backend, Pipeline, Repository/Kanban, First Front-end, and Second Front-end.
The leader of each piece is not required to create everything for that piece but rather oversee it to ensure development is on time and complete. That is, they will look for missing tests, documentation, and other issues while monitoring progress on that particular area. Each member can work on any part of the software as needed, work division and tracking will be done using a kanban board to ensure balanced distruibution. The board is the source of truth, if there are any discrepancies, the board resolves them. Additionally, there is an expectation each member will develop a complete feature utilizing the entirety of the tech stack.

