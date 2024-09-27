from enum import Enum

class Status(Enum):
    OPEN = "Open"
    CLOSED = "Closed"
    RESOLVED = "Resolved"
    IN_PROGRESS = "In Progress"