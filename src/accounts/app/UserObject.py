import app.Configs as cfg

class UserObject:
    """Class to encapsulate the data corresponding to a user"""

    def __init__(self, **kwargs):
        #Takes in a dictionary, the correct attributes get mapped to values
        for key in cfg.USER_ATTRIBUTES:
            if key in kwargs.keys():
                setattr(self, key, kwargs[key])
            else:
                setattr(self, key, None)

    def __eq__(self, comparable)->bool:
        return type(comparable) == type(self) and all(
            [
                getattr(self, key) ==  getattr(comparable, key) 
                for key in cfg.USER_ATTRIBUTES
            ]
        )
    def __str__(self):
        to_string = 'thing for the tihing\n'
        for key, value in self.__dict__.items():
            to_string += f'{key} : {value}\n'
        return to_string
