data = [int(line.rstrip('\n')) for line in open(f"data/input.txt")]

increases = len([idx for idx, val in enumerate(data[1:]) if val > data[idx]])
print(f"Part 1\t {increases} increases")

increases = len([idx for idx, _ in enumerate(data[1:-2]) if data[idx] < data[3+idx]])
print(f"Part 2\t {increases} increases")
