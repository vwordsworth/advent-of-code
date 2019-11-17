import re

class Claim:
    # Claim descriptions are formatted like: `#123 @ 3,2: 5x4`
    pattern = re.compile(r'\#(?P<id>\d+)\s@\s(?P<x>\d+),(?P<y>\d+):\s(?P<width>\d+)x(?P<height>\d+)')
    
    def __init__(self, description):
        match = self.pattern.match(description)
        self.id = match.group("id")
        self.x = int(match.group("x"))
        self.y = int(match.group("y"))
        self.width = int(match.group("width"))
        self.height = int(match.group("height"))
        self.description = description

    def get_x_range(self):
        return range(self.x, self.x+self.width)
    
    def get_y_range(self):
        return range(self.y, self.y+self.height)

    def __str__(self):
        return self.description
