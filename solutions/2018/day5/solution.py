import re
from string import ascii_lowercase


def main():
    polymer = read_input()
    result = remove_reactions(polymer)
    print("Part 1:\n\tLength of resulting polymer: {0}".format(str(len(result))))
    result = produce_shortest_possible_polymer(polymer)
    print("Part 2:\n\tLength of shortest possible polymer: {0}".format(str(len(result))))


def read_input():
    return [line.rstrip('\n') for line in open('data/input.txt')][0]


def remove_reactions(polymer):
    end = len(polymer) - 1
    i = 0

    while i < end:
        current_unit = polymer[i]
        next_unit = polymer[i+1]

        if units_will_react(current_unit, next_unit):
            polymer = polymer[:i] + polymer[i+2:]
            end -= 2
            i = max(i-1, 0)
        else:
            i += 1
    
    return polymer


def produce_shortest_possible_polymer(polymer):
    minimum_polymer = polymer
    minimum_length = len(polymer)

    for letter in ascii_lowercase:
        new_polymer = re.sub('[{0}{1}]'.format(letter, letter.upper()), "", polymer)
        new_polymer_without_reactions = remove_reactions(new_polymer)
        length_without_reactions = len(new_polymer_without_reactions)
        
        if length_without_reactions < minimum_length:
            minimum_length = length_without_reactions
            minimum_polymer = new_polymer_without_reactions
    
    return minimum_polymer


def units_will_react(current_unit, next_unit):
    return current_unit != next_unit and (current_unit == next_unit.lower() or current_unit.lower() == next_unit)


if __name__ == "__main__":
    main()