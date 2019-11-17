from enum import Enum


class LogTypes(Enum):
    
    NEW_GUARD = 1
    WAKE_UP = 2
    FALL_ASLEEP = 3

    @staticmethod
    def get_log_type(message):
        if message == "wakes up":
            return LogTypes.WAKE_UP
        elif message == "falls asleep":
            return LogTypes.FALL_ASLEEP
        else:
            return LogTypes.NEW_GUARD
