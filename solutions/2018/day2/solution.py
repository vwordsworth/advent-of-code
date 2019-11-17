from difflib import SequenceMatcher
from itertools import combinations


def main():
    ids = read_input()
    checksum = calculate_checksum(ids)
    print("Part 1:\n\tChecksum: {0}".format(str(checksum)))
    match = find_matching_sequence(ids)
    print("Part 2:\n\tId Match: {0}".format(match))


def read_input():
    return [line.rstrip('\n') for line in open('data/input.txt')]


def calculate_checksum(ids):
    two_repeats = 0
    three_repeats = 0

    for box_id in ids:
        has_two, has_three = determine_if_id_has_two_or_three_repeats(box_id)
        if has_two:
            two_repeats += 1
        if has_three:
            three_repeats += 1
    return two_repeats * three_repeats


def find_matching_sequence(ids):
    for pair in combinations(ids, 2):
        id_1 = pair[0]
        id_2 = pair[1]
        matches = SequenceMatcher(None, id_1, id_2).get_matching_blocks()

        total_match_size = 0
        total_match_string = ""
        for match in matches[:-1]:
            total_match_size += match.size
            total_match_string += id_1[match.a:match.a + match.size]
        
        if total_match_size == (matches[-1].a - 1):
            return total_match_string
    return "No match found..."


def determine_if_id_has_two_or_three_repeats(box_id):
    letter_counts = {}

    for letter in box_id:
        if letter in letter_counts:
            letter_counts[letter] += 1
        else:
            letter_counts[letter] = 1
    
    counts = set(letter_counts.values())
    return 2 in counts, 3 in counts


if __name__ == '__main__':
    main()
