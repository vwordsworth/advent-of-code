data = [tuple(line.rstrip('\n').split(" ")) for line in open(f"data/input.txt")]

depth = 0
horiz = 0

for instr, amount in data:
    amount = int(amount)
    if instr == "forward":
        horiz += amount
    elif instr == "down":
        depth += amount
    elif instr == "up":
        depth -= amount

product = depth * horiz
print(f"Part 1\t {depth} * {horiz} = {product}")


depth = 0
horiz = 0
aim = 0

for instr, amount in data:
    amount = int(amount)
    if instr == "forward":
        horiz += amount
        depth += aim*amount
    elif instr == "down":
        aim += amount
    elif instr == "up":
        aim -= amount

product = depth * horiz
print(f"Part 2\t {depth} * {horiz} = {product}")
