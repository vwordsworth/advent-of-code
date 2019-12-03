from .instruction import Instruction
from operator import mul

class Mul(Instruction):

    def __init__(self, program, instruction_ptr):
        super().__init__(program, instruction_ptr, 4)

    def set_result(self):
        a = self.program[self.instruction_params[0]]
        b = self.program[self.instruction_params[1]]
        self.program[self.instruction_params[2]] = mul(a, b)
