from operator import itemgetter
from point_map import PointMap

def main():
    point_descriptions = read_input()
    point_map = create_point_map(point_descriptions)
    closest_areas = point_map.calculate_closest_points()
    largest_area = get_largest_area(closest_areas)
    print("Part 1:\n\tSize of largest area: {0}".format(str(largest_area)))
    size = point_map.calculate_size_of_region_less_than_10000_to_all_coordinates()
    print("Part 2:\n\tSize of area with points < 10000 to all coordinates: {0}".format(str(size)))

def read_input():
    return [line.rstrip('\n') for line in open('data/input.txt')]

def create_point_map(points):
    point_map = PointMap()
    for point in points:
        point_map.add_point(point)
    return point_map

def get_largest_area(closest_areas):
    return max(closest_areas.items(), key=itemgetter(1))[1]

if __name__ == "__main__":
    main()
