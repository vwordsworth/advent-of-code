from .instruction import Instruction


class Output(Instruction):

    def __init__(self, program, instruction_ptr, params):
        super().__init__(program, instruction_ptr, 2, params)

    def set_result(self):
        value = self.program[self.instruction_params[0]] if self.get_nth_parameter_value(0) == 0 else self.instruction_params[0]
        print("Value: {0}".format(str(value)))
