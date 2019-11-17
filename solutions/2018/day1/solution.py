from itertools import cycle
from operator import add, sub


def main():
    frequency_changes = read_input()
    result = calculate_resulting_frequency(frequency_changes)
    print("Part 1:\n\tResulting frequency: {0}".format(str(result)))
    repeat = find_first_repeated_frequency(frequency_changes)
    print("Part 2:\n\tFirst repeated frequency: {0}".format(str(repeat)))


def read_input():
    return [line.rstrip('\n') for line in open('data/input.txt')]


def calculate_resulting_frequency(changes, starting_value=0):
    result = starting_value
    for change_description in changes:
        result = get_change_to_current_value(result, change_description)
    return result


def find_first_repeated_frequency(changes, starting_value=0):
    frequency_occurrences = set([0])
    result = starting_value

    for change_description in cycle(changes):
        result = get_change_to_current_value(result, change_description)
        if result in frequency_occurrences:
            break
        frequency_occurrences.add(result)    
    return result


def get_change_to_current_value(current, change_description):
    operator, value = get_operator_and_value_from_description(change_description)
    return operator(current, value)


def get_operator_and_value_from_description(description):
    operator = get_operator_from_operator_string(description[0])
    value = int(description[1:])
    return operator, value


def get_operator_from_operator_string(operator_string):
    operators = {
        "+": add,
        "-": sub
    }
    return operators[operator_string]


if __name__ == '__main__':
    main()
