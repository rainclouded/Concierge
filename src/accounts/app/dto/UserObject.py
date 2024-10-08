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
