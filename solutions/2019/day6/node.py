
class Node:

    def __init__(self, name):
        self.name = name
        self.parents = []
        self.children = []
    
    def get_name(self):
        return self.name

    def get_children(self):
        return self.children

    def get_child_count(self):
        return len(self.children)

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
    
    def get_total_count(self, dist_to_parent=0):
        count = dist_to_parent
        for child in self.children:
            count += child.get_total_count(dist_to_parent+1)
        return count
    
    def __rep__(self):
        return repr(self.name)
    
    def __str__(self):
        return self.name

    def __hash__(self):
        return hash((self.name))
