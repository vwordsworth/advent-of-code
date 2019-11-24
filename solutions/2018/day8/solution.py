from collections import deque
from node import Node


def main():
    data = read_input()
    license_sum, root = get_license_sum(data)
    print("Part 1:\n\tTraversal license sum: {0}".format(str(license_sum)))
    root_value = find_value_of_root(root)
    print("Part 2:\n\tValue of root node: {0}".format(str(root_value)))

def get_license_sum(data):
    license_sum = 0
    root = Node(data[0], data[1])
    current_parent = root
    stack = deque()
    stack.append(root)

    i = 2
    while i < len(data):    
        next_location = i+2

        if not current_parent.has_processed_all_children():
            child = Node(data[i], data[i+1])
            current_parent.add_child(child)

            if child.has_children():        
                stack.append(current_parent)
                current_parent = child
            else:
                node_sum, next_location = process_metadata_and_return_sum_and_next_index(child, i+2, data)
                license_sum += node_sum
        elif not current_parent.has_processed_all_metadata():
            node_sum, next_location = process_metadata_and_return_sum_and_next_index(current_parent, i, data)
            license_sum += node_sum

            current_parent = stack.pop()
        
        i = next_location

    return license_sum, root


def find_value_of_root(root):
    value = 0
    if not root.has_children():
        value = root.get_metadata_sum()
    else:
        children = root.get_children()
        for metadata in root.get_metadata():
            if metadata > 0:
                try:
                    value += find_value_of_root(children[metadata-1])
                except IndexError:
                    pass
    return value


def process_metadata_and_return_sum_and_next_index(node, start, data):
    metadata_index = start
    while not node.has_processed_all_metadata():
        metadata = data[metadata_index]
        node.add_metadata(metadata)
        metadata_index += 1
    return node.get_metadata_sum(), metadata_index


def read_input():
    return [line.rstrip('\n') for line in open('data/input.txt')][0].split(" ")


if __name__ == "__main__":
    main()
