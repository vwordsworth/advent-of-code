from .instruction import Instruction


class Load(Instruction):

    def __init__(self, program, instruction_ptr, params):
        super().__init__(program, instruction_ptr, 2, params)

    def set_result(self):
        self.program[self.instruction_params[0]] = 5
