from copy import deepcopy
from instructions.instruction_factory import InstructionFactory, HltEncounteredException


def main():
    run_program(_read_input())


def run_program(program):
    instruction_ptr = 0
    prog_len = len(program)

    while instruction_ptr < prog_len:
        try:
            instruction = InstructionFactory(program, instruction_ptr).get_instruction()
        except HltEncounteredException:
            break
        instruction.set_result()
        instruction_ptr = instruction.get_next_instruction_pointer()
    return program[0]


def _read_input():
    return [int(val) for val in [line.rstrip('\n') for line in open('data/input.txt')][0].split(",")]


if __name__ == "__main__":
    main()
