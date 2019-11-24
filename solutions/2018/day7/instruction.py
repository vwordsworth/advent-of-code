import re

class Instruction:

    instruction_pattern = re.compile(r'Step (?P<upstream>[A-Z]) must be finished before step (?P<downstream>[A-Z]) can begin.')

    def __init__(self, instruction):
        match = self.instruction_pattern.match(instruction)
        self.upstream = match.group("upstream")
        self.downstream = match.group("downstream")

    def get_upstream(self):
        return self.upstream

    def get_downstream(self):
        return self.downstream

    def __str__(self):
        return "{0}-->{1}".format(self.upstream, self.downstream)
