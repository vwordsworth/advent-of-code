from copy import deepcopy
from instructions.instruction_factory import InstructionFactory, HltEncounteredException, UnexpectedInstructionException
from itertools import product


def main():
    program = _read_input()
    
    noun = 12
    verb = 2
    result = run_program(deepcopy(program), noun, verb)
    print("Part 1:\n\tprogram[0] = {0}".format(str(result)))

    desired_result = 19690720
    noun, verb = find_noun_and_verb_to_generate_result(program, desired_result)
    result = (100*noun)+verb
    print("Part 2:\n\t(100*{0})+{1} = {2}".format(str(noun), str(verb), str(result)))


def run_program(program, noun, verb):
    instruction_ptr = 0
    prog_len = len(program)
    
    program[1] = noun
    program[2] = verb

    while instruction_ptr < prog_len:
        try:
            instruction = InstructionFactory(program, instruction_ptr).get_instruction()
        except HltEncounteredException:
            break
        instruction.set_result()
        instruction_ptr = instruction.get_next_instruction_pointer()
    return program[0]


def find_noun_and_verb_to_generate_result(program, desired_result, viable_range=100):
    for noun, verb in product(range(viable_range), repeat=2):
        result = run_program(deepcopy(program), noun, verb)
        if result == desired_result:
            return noun, verb


def _read_input():
    return [int(val) for val in [line.rstrip('\n') for line in open('data/input.txt')][0].split(",")]


if __name__ == "__main__":
    main()
