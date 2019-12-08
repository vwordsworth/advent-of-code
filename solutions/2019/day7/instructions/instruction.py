from abc import abstractmethod, ABCMeta
from operator import add, mul, pow


class Instruction(metaclass=ABCMeta):

    def __init__(self, program, pointer, size, params):
        self.program = program
        self.pointer = pointer
        self.override_pointer = None
        self.size = size
        self.instruction_params = program[pointer+1:pointer+size]
        self.params = params
    
    def get_next_instruction_pointer(self):
        return self.override_pointer if self.override_pointer else self.pointer + self.size

    def get_nth_parameter_value(self, n):
        return (self.params // pow(10,n)) % 10

    @abstractmethod
    def set_result(self):
        return
