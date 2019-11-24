from copy import deepcopy
from instruction import Instruction
from node import Node

def main():
    instructions = read_input()
    option_set = build_map(instructions)
    traversal = get_traversal_sequence(deepcopy(option_set))
    print("Part 1:\n\tTraversal sequence: {0}".format(traversal))
    total_time = get_total_time_for_five_workers(deepcopy(option_set))
    print("Part 2:\n\tTotal time for five workers to complete: {0}".format(str(total_time)))

def read_input():
    return [Instruction(line.rstrip('\n')) for line in open('data/input.txt')]


def build_map(instructions):
    nodes = {}
    entry_points = set()

    for instruction in instructions:
        upstream = instruction.get_upstream()
        downstream = instruction.get_downstream()
        
        if upstream in nodes:
            parent = nodes[upstream]
        else:
            parent = Node(upstream)
            nodes[upstream] = parent
            entry_points.add(parent)
        
        if downstream in nodes:
            child = nodes[downstream]
            if child in entry_points:
                entry_points.remove(child)
        else:
            child = Node(downstream)
            nodes[downstream] = child

        parent.add_child_node(child)
        child.add_parent_node(parent)
    
    return entry_points


def get_traversal_sequence(option_set):
    traversal = ""
    completed_set = set()
    while option_set:
        next_node = _get_next_node_if_available(option_set, completed_set)
        
        option_set.remove(next_node)
        completed_set.add(next_node)
        traversal += str(next_node)
        
        for child in next_node.get_children():
            option_set.add(child)
    
    return traversal


def get_total_time_for_five_workers(option_set):
    total_time = -1
    completed_set = set()
    available_workers = 5
    active_work = {}

    while option_set or active_work:
        total_time += 1
        
        num_completed = _handle_completed_tasks(option_set, completed_set, active_work)
        available_workers += num_completed

        if available_workers == 0:
            continue

        for _ in range(available_workers): 
            next_node = _get_next_node_if_available(option_set, completed_set)
            if not next_node:
                break
            
            option_set.remove(next_node)
            available_workers -= 1
            active_work[next_node] = next_node.get_total_time()   
    
    return total_time


def _get_next_node_if_available(option_set, completed_set):
    next_node = None
    for node in sorted(option_set):
        all_parents_completed = True
        for parent in node.get_parents():
            if parent not in completed_set:
                all_parents_completed = False
        if all_parents_completed:
            next_node = node
            break
    return next_node


def _handle_completed_tasks(option_set, completed_set, active_work):
    number_completed = 0
    completed_tasks = []
    for task in active_work:
        time_remaining = active_work[task] - 1
        if time_remaining == 0:
            completed_tasks.append(task)
        else:
            active_work[task] = time_remaining
    
    for task in completed_tasks:
        del active_work[task]
        completed_set.add(task)
        number_completed += 1
        for child in task.get_children():
            option_set.add(child)
    return number_completed


if __name__ == "__main__":
    main()