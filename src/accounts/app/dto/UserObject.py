"""
Module for the UserObject dto
"""
import app.Configs as cfg

class UserObject:
    """Class to encapsulate the data corresponding to a user"""

    def __init__(self, **kwargs):
        #Takes in a dictionary, the correct attributes get mapped to values
        for key in cfg.USER_ATTRIBUTES:
            if key in kwargs:
                setattr(self, key, kwargs[key])
            else:
                setattr(self, key, None)


    def __eq__(self, comparable)->bool:
        #Check equality based on dictionary values and type
        return (
            type(comparable) is type(self)
            and all(
                getattr(self, key) ==  getattr(comparable, key)
                for key in cfg.USER_ATTRIBUTES
            )
        )
    def __str__(self):
        # Create a dictionary of attribute names and their values
        attribute_dict = {key: getattr(self, key) for key in cfg.USER_ATTRIBUTES}
        
        # Create a string representation, which will include each attribute and its value
        attribute_str = ', '.join(f"{key}={value}" for key, value in attribute_dict.items())
        
        # Return the formatted string
        return f"{self.__class__.__name__}({attribute_str})"