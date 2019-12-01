
def main():
    data = _read_input()
    fuel_mass = calculate_fuel_mass(data)
    print("Part 1:\n\tFuel required: {0}".format(str(fuel_mass)))
    new_fuel_mass = calculate_fuel_mass(data, include_new_fuel=True)
    print("Part 2:\n\tAccounting for new fuel's mass: {0}".format(str(new_fuel_mass)))


def calculate_fuel_mass(data, include_new_fuel=False):
    total_sum = 0
    for mass in data:
        fuel_mass = _calculate_fuel_required_for_mass(int(mass))
        total_sum += fuel_mass + (_get_extra_mass_for_new_fuel(fuel_mass) if include_new_fuel else 0)
    return total_sum


def _get_extra_mass_for_new_fuel(new_fuel_mass):
    extra_mass = 0
    while new_fuel_mass > 0:
        next_new_fuel_mass = _calculate_fuel_required_for_mass(new_fuel_mass)
        extra_mass += next_new_fuel_mass
        new_fuel_mass = next_new_fuel_mass
    return extra_mass


def _calculate_fuel_required_for_mass(mass):
    return max((mass//3)-2, 0)


def _read_input():
    return [line.rstrip('\n') for line in open('data/input.txt')]


if __name__ == "__main__":
    main()
