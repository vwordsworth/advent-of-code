from abc import abstractmethod, ABCMeta
from operator import add, mul

class Instruction(metaclass=ABCMeta):

    def __init__(self, program, pointer, size):
        self.program = program
        self.pointer = pointer
        self.size = size
        self.instruction_params = program[pointer+1:pointer+size]
    
    def get_next_instruction_pointer(self):
        return self.pointer + self.size

    @abstractmethod
    def set_result(self):
        return
