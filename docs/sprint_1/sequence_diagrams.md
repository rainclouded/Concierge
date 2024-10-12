# Sequence Diagrams
## Our Current Features
### Amenities
* Offers guests detailed information about available hotel amenities, including operating hours and description, displayed on a dashboard. Staff can create, update, or delete amenity listings to ensure that guests have up-to-date information.

#### Guest Scenario: View all amenities and view specified amenity 
![](/docs/sprint_1/images/Guest_Amenities_Sequence_Diagram.svg)

#### Staff Scenario: View all amenities and update a specified amenity with new information
![](/docs/sprint_1/images/Staff_Amenities_Sequence_Diagram.svg)

### Incident Reports
* Enables hotel managers to monitor past and current incidents through a dedicated dashboard that displays details such as severity, status and description. Managers can update reports in real time, ensuring accurate information is displayed as an incident case progresses.

#### Staff Scenario: View all incident reports and update a specified report with new information
![](/docs/sprint_1/images/Staff_Incident_Reports_Sequence_Diagram.svg)

### Accounts
* Provides secure access for both staff and guests to the hotel management system. Users log in with their room code and gain access to request services. While staff can create accounts and manage permissions of each accounts, while users have the option to customize their settings.

![Guest UI account diagram](/docs/sprint_1/images/Guest_Accounts_Sequence_diagram.svg)
![Staff UI account diagram](/docs/sprint_1/images/Staff_Accounts_Sequence_diagram.svg)

### Staff UI
* This diagram represents the interaction flow of the staff UI. It shows how the staff logs in, navigates between different dashboard tabs (Amenities and Incident Report), and interacts with services that handle data fetching and CRUD operations via Nginx.

#### Staff interacts with Amenities and Incident Reports

![](/docs/sprint_1/images/Staff_UI_Sequence_Diagram.svg)