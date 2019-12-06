from node import Node
from sys import maxsize


def main():
    data = _read_input()
    nodes = generate_graph(data)
    count = nodes["COM"].get_total_count()
    print("Part 1:\n\tTotal orbit count: {0}".format(str(count)))
    transfers = get_orbital_transfer_count(nodes)
    print("Part 2:\n\tTransfers between me and Santa: {0}".format(str(transfers)))


def get_orbital_transfer_count(nodes):
    my_parent = nodes["YOU"].get_parents()[0]
    santa_parent = nodes["SAN"].get_parents()[0]

    visited = [my_parent]
    to_visit = my_parent.get_children() + my_parent.get_parents()
    distances = {my_parent: (0, None)}
    current_node = to_visit.pop()

    while to_visit:
        to_visit += [item for item in (current_node.get_children() + current_node.get_parents()) if item not in visited]
        visited.append(current_node)

        min_parent_dist, min_parent = _get_closest_neighbor_and_distance(current_node, distances)
        current_distance = maxsize if current_node not in distances else distances[current_node]
        if min_parent_dist + 1 < current_distance:
            distances[current_node] = (min_parent_dist + 1, min_parent)

        current_node = to_visit.pop()

    return distances[santa_parent][0]


def _get_closest_neighbor_and_distance(current_node, distances):
    min_neighbor_dist = maxsize
    min_neighbor = None
    for neighbor in current_node.get_parents() + current_node.get_children():
        try:
            neighbor_dist, _ = distances[neighbor]
            if neighbor_dist < min_neighbor_dist:
                min_neighbor_dist = neighbor_dist
                min_neighbor = neighbor
        except KeyError:
            # If there is a KeyError, we haven't yet visited the neighbor
            pass

    return min_neighbor_dist, min_neighbor

def generate_graph(data):
    nodes = {}
    for orbit in data:
        center = orbit[0]
        orbiter = orbit[1]
        
        center_node = None
        orbiter_node = None

        if center in nodes:
            center_node = nodes[center]
        else:
            center_node = Node(center)
            nodes[center] = center_node

        if orbiter in nodes:
            orbiter_node = nodes[orbiter]
        else:
            orbiter_node = Node(orbiter)
            nodes[orbiter] = orbiter_node

        orbiter_node.add_parent_node(center_node)
        center_node.add_child_node(orbiter_node)

        nodes[center] = center_node
        nodes[orbiter] = orbiter_node
    return nodes


def _read_input():
    return [val.split(")") for val in [line.rstrip('\n') for line in open('data/input.txt')]]


if __name__ == "__main__":
    main()
