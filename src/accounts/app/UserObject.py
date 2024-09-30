import Configs as cfg

class UserObject:
    def __init__(self, **kwargs):
        for key in cfg.USER_ATTRIBUTES:
            if key in kwargs.keys():
                setattr(self, key, kwargs[key])
            else:
                setattr(self, key, None)

    