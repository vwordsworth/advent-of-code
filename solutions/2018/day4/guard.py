from sleep import Sleep

class Guard:

    def __init__(self, id):
        self.id = id
        self.sleeps = []
        self.minute_frequency = {}
        self.most_frequent_minute = -1
        self.most_frequent_count = -1
        self.total_sleep_time = 0
    
    def add_sleep(self, start_minute, wake_minute):
        sleep = Sleep(start_minute, wake_minute)
        self.total_sleep_time += sleep.get_duration()
        self.sleeps.append(sleep)

        for minute in range(start_minute, wake_minute):
            if minute in self.minute_frequency:
                self.minute_frequency[minute] += 1
            else:
                self.minute_frequency[minute] = 1
            
            if self.minute_frequency[minute] > self.most_frequent_count:
                self.most_frequent_count = self.minute_frequency[minute]
                self.most_frequent_minute = minute

    def __str__(self):
        return "Id: {id}\tTotal Sleep Time: {time}".format(id=self.id, time=self.total_sleep_time)
