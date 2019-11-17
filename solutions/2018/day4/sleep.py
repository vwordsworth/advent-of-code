
class Sleep:

    def __init__(self, start_minute, end_minute):
        self.start_minute = start_minute
        self.end_minute = end_minute
    
    def get_duration(self):
        return self.end_minute - self.start_minute
    
    def __str__(self):
        return "[{start},{end}]".format(self.start_minute,self.end_minute)
