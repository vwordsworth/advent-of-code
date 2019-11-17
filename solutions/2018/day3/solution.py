from claim import Claim
from coordinate import Coordinate


def main():
    claims = read_input()
    overlapping_squares = count_overlapping_squares(claims)
    print("Part 1:\n\tOverlapping Squares: {0}".format(str(overlapping_squares)))
    claim_id = find_fully_isolated_claim(claims)
    print("Part 2:\n\tIsolated Claim: {0}".format(claim_id))


def read_input():
    return [line.rstrip('\n') for line in open('data/input.txt')]


def count_overlapping_squares(claims):
    claimed_fabric = {}
    duplicated_squares = set()

    for claim_description in claims:
        claim = Claim(claim_description)
        for x in claim.get_x_range():
            for y in claim.get_y_range():
                coordinate = Coordinate(x,y).get_value()
                if coordinate in claimed_fabric:
                    claimed_fabric[coordinate] = 'X'
                    duplicated_squares.add(coordinate)
                else:
                    claimed_fabric[coordinate] = claim.id
    
    return len(duplicated_squares)


def find_fully_isolated_claim(claims):
    claimed_fabric = {}
    intersected_claims = set()
    all_claims = set()

    for claim_description in claims:
        claim = Claim(claim_description)
        for x in claim.get_x_range():
            for y in claim.get_y_range():
                coordinate = Coordinate(x,y).get_value()
                if coordinate in claimed_fabric:
                    claimed_fabric[coordinate].append(claim.id)
                    for claim_id in claimed_fabric[coordinate]:
                        intersected_claims.add(claim_id)
                else:
                    claimed_fabric[coordinate] = [claim.id]
                    all_claims.add(claim.id)
    
    return (all_claims - intersected_claims).pop()


if __name__ == '__main__':
    main()