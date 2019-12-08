from .instruction import Instruction


class Load(Instruction):

    def __init__(self, program, instruction_ptr, params, value):
        super().__init__(program, instruction_ptr, 2, params)
        self.input_value = value
    
    def set_result(self):
        self.program[self.instruction_params[0]] = self.input_value
