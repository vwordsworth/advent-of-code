data = [line.rstrip('\n') for line in open(f"data/input.txt")]

gamma = 0
epsilon = 0
half = len(data) // 2
num_size = len(data[0])

for pos in range(num_size):
    pos_sum = 0
    for num in data:
        pos_sum += int(num[pos])

    gamma *= 2
    epsilon *= 2

    if pos_sum > half:
        gamma += 1
    else:
        epsilon += 1

power = gamma * epsilon
print(f"Part 1\t {gamma} * {epsilon} = {power}")

o2 = data
for pos in range(num_size):
    pos_class = {"0": [], "1": []}
    for num in o2:
        pos_class[num[pos]].append(num)

    if len(pos_class["1"]) >= len(pos_class["0"]):
        o2 = pos_class["1"]
    else:
        o2 = pos_class["0"]

    if len(o2) == 1:
        o2 = int(o2[0], 2)
        break

co2 = data
for pos in range(num_size):
    pos_class = {"0": [], "1": []}
    for num in co2:
        pos_class[num[pos]].append(num)

    if len(pos_class["1"]) >= len(pos_class["0"]):
        co2 = pos_class["0"]
    else:
        co2 = pos_class["1"]

    if len(co2) == 1:
        co2 = int(co2[0], 2)
        break

life = o2 * co2
print(f"Part 1\t {o2} * {co2} = {life}")
