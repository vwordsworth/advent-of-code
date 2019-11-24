
class Node:

    def __init__(self, name):
        self.name = name
        self.additional_time = ord(name) - 64
        self.parents = []
        self.children = []
    
    def get_name(self):
        return self.name
    
    def get_total_time(self):
        return self.additional_time + 60

    def get_additional_time(self):
        return self.additional_time

    def get_children(self):
        return self.children

    def get_parents(self):
        return self.parents

    def add_child_node(self, child):
        self.children.append(child)

    def add_parent_node(self, parent):
        self.parents.append(parent)

    def remove_child_node(self, child):
        self.children.remove(child)
    
    def get_description_of_all_dependents(self):
        description = self.name if not self.children else ""
        for child in self.children:
            description += self.name + "-->" + child.get_description_of_all_dependents() + " "
        return description
    
    def __rep__(self):
        return repr(self.name)
    
    def __str__(self):
        return self.name

    def __lt__(self, other):
        return self.name < other.name

    def __gt__(self, other):
        return self.name > other.name

    def __le__(self, other):
        return self.name <= other.name

    def __ge__(self, other):
        return self.name >= other.name

    def __eq__(self, other):
        return (self.name, self.children, self.parents) == (other.name, other.children, other.parents)

    def __ne__(self, other):
        return not(self == other)

    def __hash__(self):
        return hash((self.name, self.additional_time))
