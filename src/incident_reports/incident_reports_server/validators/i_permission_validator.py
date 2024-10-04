from abc import ABC, abstractmethod
from typing import List

class IPermissionValidator(ABC):
    @abstractmethod
    def validate_permissions(self) -> bool:
        pass