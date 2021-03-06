from .instruction import Instruction


class JumpTrue(Instruction):

    def __init__(self, program, instruction_ptr, params):
        super().__init__(program, instruction_ptr, 3, params)

    def set_result(self):
        a = self.program[self.instruction_params[0]] if self.get_nth_parameter_value(0) == 0 else self.instruction_params[0]
        b = self.program[self.instruction_params[1]] if self.get_nth_parameter_value(1) == 0 else self.instruction_params[1]
        if a != 0:
            self.override_pointer = b
