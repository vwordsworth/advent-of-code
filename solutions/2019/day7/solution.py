from copy import deepcopy
from instructions.instruction_factory import InstructionFactory, HltEncounteredException
from instructions.load import Load
from instructions.output import Output
from itertools import permutations


num_phases = 5
outputs = []
phase_inputs = []
is_phase_halted = []
active_phase = 0
pointers = []
programs = []


def main():
    program = _read_input()
    best_output = 0

    possible_phases = list(permutations(range(5,10)))
    for combination in possible_phases:
        global phase_inputs, is_phase_halted, outputs, active_phase, pointers, programs
        
        programs = [deepcopy(program)]*num_phases
        pointers = [0]*num_phases
        outputs = [0]*num_phases
        is_phase_halted = [False]*num_phases
        phase_inputs = list(combination)
        active_phase = 0
        
        while phases_are_executing():
            execute_phase_if_not_halted(active_phase)
    
        best_output = max(best_output, outputs[-1])

    print("Max output after feedback loop: {0}".format(str(best_output)))


def phases_are_executing():
    return not all(is_phase_halted)


def increment_current_phase():
    global active_phase
    active_phase = (active_phase+1) % num_phases


def execute_phase_if_not_halted(phase):
    global pointers, is_phase_halted

    if not is_phase_halted[phase]:
        try:
            pointers[phase] = execute_instruction(programs[phase], pointers[phase], phase)
        except HltEncounteredException:
            is_phase_halted[phase] = True
    else:
        increment_current_phase()


def execute_instruction(program, instruction_ptr, phase):
    global phase_inputs, outputs, active_phase
    
    if phase_inputs[phase]:
        input_value = phase_inputs[phase]
    else:
        input_value = outputs[(phase-1)%num_phases]
    instruction = InstructionFactory(program, instruction_ptr, input_value).get_instruction()
        
    instruction.set_result()
    instruction_ptr = instruction.get_next_instruction_pointer()

    if isinstance(instruction, Load):
        phase_inputs[phase] = None
    if isinstance(instruction, Output):
        outputs[phase] = instruction.get_output_value()
        increment_current_phase()

    return instruction_ptr


def _read_input():
    return [int(val) for val in [line.rstrip('\n') for line in open('data/input.txt')][0].split(",")]


if __name__ == "__main__":
    main()
