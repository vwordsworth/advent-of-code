from log_type import LogTypes
import re


class Log:

    log_pattern = re.compile(r'\[\d+-\d+-\d+\s\d+:(?P<minute>\d+)\]\s(?P<message>.*)$')
    shift_pattern = re.compile(r'Guard #(?P<guard_number>\d+) begins shift')
    
    def __init__(self, message):
        match = self.log_pattern.match(message)
        self.minute = match.group("minute")
        self.message = match.group("message")
        self.log_type = LogTypes.get_log_type(self.message)

    def get_guard_number_from_new_guard_message(self):
        match = self.shift_pattern.match(self.message)
        return match.group("guard_number")
