
class Coordinate:

    def __init__(self, description):
        values = description.split(",")
        self.x = int(values[0])
        self.y = int(values[1])
    
    def get_x(self):
        return self.x

    def get_y(self):
        return self.y

    def get_manhattan_distance(self, x, y):
        return abs(self.x - x) + abs(self.y - y)

    def get_value(self):
        return "({x},{y})".format(x=str(self.x), y=str(self.y))

    def __lt__(self, other):
        return self.x < other.x or (self.x == other.x and self.y < other.y)

    def __gt__(self, other):
        return self.x > other.x or (self.x == other.x and self.y > other.y)

    def __le__(self, other):
        return self.x <= other.x or (self.x == other.x and self.y <= other.y)

    def __ge__(self, other):
        return self.x >= other.x or (self.x == other.x and self.y >= other.y)

    def __eq__(self, other):
        return self.x == other.x and self.y == other.y

    def __ne__(self, other):
        return self.x != other.x or self.y != other.y
