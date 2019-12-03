from .add import Add
from .mul import Mul


class InstructionFactory:

    def __init__(self, program, instruction_ptr):
        self.program = program
        self.instruction_ptr = instruction_ptr
        self.opcode = self.program[self.instruction_ptr]
    
    def get_instruction(self):
        instruction = None

        if self.opcode == 1:
            return Add(self.program, self.instruction_ptr)
        elif self.opcode == 2:
            return Mul(self.program, self.instruction_ptr)
        elif self.opcode == 99:
            raise HltEncounteredException
        else:
            raise UnexpectedInstructionException

        return instruction


class HltEncounteredException(Exception):
    pass


class UnexpectedInstructionException(Exception):
    pass
