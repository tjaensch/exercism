class Clock(object):

    """docstring for Clock"""

    def __init__(self, hours, minutes):

        self.minutes = minutes % 60
        self.hours = (hours + minutes // 60) % 24

    def __str__(self):
        return "%02d:%02d" % (self.hours, self.minutes)

    def __eq__(self, clock_instance):
        return self.hours == clock_instance.hours and \
            self.minutes == clock_instance.minutes

    def add(self, minutes):
        min = self.minutes + minutes
        self.minutes = min % 60
        self.hours = (self.hours + min // 60) % 24
        return self