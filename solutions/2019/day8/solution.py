from collections import Counter
from sys import maxsize
from matplotlib.pyplot import savefig, subplots
from matplotlib.cm import Greys

WIDTH = 25
HEIGHT = 6

TRANSPARENT = 2


def main():
    image = _read_input()
    
    layer_size = WIDTH * HEIGHT
    number_layers = len(image)//layer_size
    
    min_0s = maxsize
    min_0s_layer_counts = None

    for layer_num in range(number_layers):
        layer_counts = Counter(image[(layer_num*layer_size):(layer_num*layer_size)+layer_size])
        if layer_counts[0] < min_0s:
            min_0s = layer_counts[0]
            min_0s_layer_counts = layer_counts
    print(min_0s_layer_counts[1]*min_0s_layer_counts[2])

    final_image = []
    for row in range(HEIGHT):
        image_row = []
        for col in range(WIDTH):
            color = get_color_for_position(row*WIDTH + col, image, number_layers, layer_size)
            image_row.append(color)
        final_image.append(image_row)
    generate_plot(final_image)


def get_color_for_position(position, image, number_layers, layer_size):
    for layer in range(number_layers):
        color = image[(layer*layer_size)+position]
        if color != TRANSPARENT:
            return color


def generate_plot(array):
    fig, ax = subplots()
    imgplot = ax.imshow(array, cmap=Greys)
    savefig('message.png')


def _read_input():
    return [int(val) for val in [line.rstrip('\n') for line in open('data/input.txt')][0]]


if __name__ == "__main__":
    main()
