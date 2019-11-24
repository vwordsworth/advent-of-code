
class Node:

    def __init__(self, child_count, metadata_count):
        self.child_count = int(child_count)
        self.children_processed = 0
        self.metadata_count = int(metadata_count)
        self.metadata_processed = 0
        self.metadata = []
        self.metadata_sum = 0

    def has_children(self):
        return self.child_count > 0

    def has_processed_all_children(self):
        return self.children_processed == self.child_count

    def has_processed_all_metadata(self):
        return self.metadata_processed == self.metadata_count

    def get_metadata_count(self):
        return self.metadata_count

    def get_metadata(self):
        return self.metadata

    def get_metadata_sum(self):
        return self.metadata_sum
    
    def add_metadata(self, metadata):
        self.metadata.append(int(metadata))
        self.metadata_sum += int(metadata)
        self.metadata_processed += 1

    def increment_processed_children(self):
        self.children_processed += 1

    def __str__(self):
        return "Children: {0} ({1}), Metadata {2} ({3})".format(self.child_count, self.children_processed, self.metadata_count, self.metadata_processed)