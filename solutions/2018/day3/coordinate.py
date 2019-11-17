
class Coordinate:

    def __init__(self, x, y):
        self.x = x
        self.y = y
    
    def get_value(self):
        return "({x},{y})".format(x=self.x, y=self.y)
