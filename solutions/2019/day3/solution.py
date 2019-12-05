from sys import maxsize


def main():
    wires = _read_input()

    w_1 = wires[0].split(",")
    w_2 = wires[1].split(",")

    w = {(0,0):0}
    current_x, current_y, current_steps = (0, 0, 0)

    for point in w_1:
        direction = point[0]
        amount = int(point[1:])
        i = 0
        while i < amount:
            current_steps += 1
            current_x, current_y = get_updated_x_y_position(current_x, current_y, direction)
            add_point_to_map(w, current_x, current_y, current_steps)
            i += 1

    min_distance, min_combined_steps = (maxsize, maxsize)
    current_x, current_y, current_steps = (0, 0, 0)

    for point in w_2:
        direction = point[0]
        amount = int(point[1:])
        i = 0
        while i < amount:
            current_steps += 1
            current_x, current_y = get_updated_x_y_position(current_x, current_y, direction)
            min_distance = update_minimum_manhattan_distance_if_intersection(w, current_x, current_y, min_distance)
            min_combined_steps = update_minimum_steps_if_intersection(w, current_x, current_y, current_steps, min_combined_steps)
            i += 1
    
    print("Part 1:\n\tMinimum Manhattan distance = {0}".format(str(min_distance)))
    print("Part 2:\n\tMinimum steps distance = {0}".format(str(min_combined_steps)))


def _read_input():
    return [line.rstrip('\n') for line in open('data/input.txt')]


def get_updated_x_y_position(current_x, current_y, direction):
    if direction == "R":
        current_x += 1
    elif direction == "L":
        current_x -= 1
    elif direction == "U":
        current_y += 1
    elif direction == "D":
        current_y -= 1
    return current_x, current_y


def add_point_to_map(w, current_x, current_y, current_steps):
    if (current_x, current_y) not in w:
        w[(current_x, current_y)] = current_steps


def update_minimum_manhattan_distance_if_intersection(w, current_x, current_y, min_distance):
    if (current_x,current_y) in w:
        min_distance = min(min_distance, get_manhattan_distance(current_x, current_y))  
    return min_distance   


def update_minimum_steps_if_intersection(w, current_x, current_y, current_steps, min_steps):
    if (current_x,current_y) in w:
        min_steps = min(min_steps, get_steps_to_current_intersection(w, current_x, current_y, current_steps))  
    return min_steps  


def get_steps_to_current_intersection(w, current_x, current_y, current_steps):
    return w[(current_x, current_y)] + current_steps


def get_manhattan_distance(current_x, current_y):
    return abs(current_x) + abs(current_y)


if __name__ == "__main__":
    main()
