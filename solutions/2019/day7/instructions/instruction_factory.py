from .add import Add
from .equals import Equals
from .jump_false import JumpFalse
from .jump_true import JumpTrue
from .less_than import LessThan
from .load import Load
from .mul import Mul
from .output import Output


class InstructionFactory:

    def __init__(self, program, instruction_ptr, input_value):
        self.program = program
        self.instruction_ptr = instruction_ptr
        self.input_value = input_value
        self.opcode = self.program[self.instruction_ptr] % 100
        self.params = self.program[self.instruction_ptr] // 100

    
    def get_instruction(self):
        instruction = None
        args = [self.program, self.instruction_ptr, self.params]

        if self.opcode == 1:
            return Add(*args)
        elif self.opcode == 2:
            return Mul(*args)
        elif self.opcode == 3:
            return Load(*(args+[self.input_value]))
        elif self.opcode == 4:
            return Output(*args)
        elif self.opcode == 5:
            return JumpTrue(*args)
        elif self.opcode == 6:
            return JumpFalse(*args)
        elif self.opcode == 7:
            return LessThan(*args)
        elif self.opcode == 8:
            return Equals(*args)
        elif self.opcode == 99:
            raise HltEncounteredException
        else:
            raise UnexpectedInstructionException

        return instruction


class HltEncounteredException(Exception):
    pass


class UnexpectedInstructionException(Exception):
    pass
