from coordinate import Coordinate
from operator import attrgetter
from sys import maxsize

class PointMap:

    def __init__(self):
        self.coordinates = []
        self.max_x = -1
        self.max_y = -1
        self.min_x = maxsize
        self.min_y = maxsize
    
    def add_point(self, point):
        coordinate = Coordinate(point)
        self.coordinates.append(coordinate)
        self._update_map_bounds(coordinate)
    
    def get_coordinates(self):
        return self.coordinates

    def get_number_coordinates(self):
        return len(self.coordinates)

    def calculate_closest_points(self):
        x_range, y_range = self._get_x_y_map_bounds()
        infinite_coordinates = set()
        closest_areas = {}

        for x in x_range:
            for y in y_range:
                closest = self._get_closest_coordinate(x, y)
                if closest == "equidistant":
                    continue
                elif closest in closest_areas:
                    closest_areas[closest] = closest_areas[closest] + 1
                else:
                    closest_areas[closest] = 1
                
                if self._is_area_adjacent_to_infinite(x, y):
                    infinite_coordinates.add(closest)

        return {k:v for k,v in closest_areas.items() if k not in infinite_coordinates}
    
    def calculate_size_of_region_less_than_10000_to_all_coordinates(self):
        x_range, y_range = self._get_x_y_map_bounds()
        points_less_than_10000 = 0

        for x in x_range:
            for y in y_range:
                total_distance = self._get_distance_to_all_coordinates(x, y)
                if total_distance < 10000:
                    points_less_than_10000 += 1
        
        return points_less_than_10000

    def _get_closest_coordinate(self, x, y):
        min_distance = maxsize
        closest_coordinate = []

        for coordinate in self.coordinates:
            distance = coordinate.get_manhattan_distance(x, y)
            if distance == min_distance:
                closest_coordinate.append(coordinate.get_value)
            elif distance < min_distance:
                min_distance = distance
                closest_coordinate = [coordinate.get_value()]
        
        return closest_coordinate[0] if len(closest_coordinate) == 1 else "equidistant"
    
    def _get_distance_to_all_coordinates(self, x, y):
        total_distance = 0

        for coordinate in self.coordinates:
            total_distance += coordinate.get_manhattan_distance(x, y)
        
        return total_distance

    def _is_area_adjacent_to_infinite(self, x, y):
        return x == self.min_x or x == self.max_x or y == self.max_y or y == self.min_y

    def _update_map_bounds(self, coordinate):
        new_x = coordinate.get_x()
        new_y = coordinate.get_y()

        if new_x > self.max_x:
            self.max_x = new_x
        if new_x < self.min_x:
            self.min_x = new_x
        if new_y > self.max_y:
            self.max_y = new_y
        if new_y < self.min_y:
            self.min_y = new_y

    def _get_x_y_map_bounds(self):
        return range(self.min_x, self.max_x+1), range(self.min_y, self.max_y+1)
