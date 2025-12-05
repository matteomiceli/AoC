from utils import get_lines

data = get_lines(day=5)
ranges = []
ids = []
for line in data:
    if "-" in line:
        ranges.append([int(r) for r in line.split("-")])

    elif line == "":
        continue
    else: 
        ids.append(int(line))

def part1():
    fresh = 0
    for id in ids:
        for range in ranges:
            if id >= range[0] and id <= range[1]:
                fresh += 1 
                break
    print(">>>", fresh)


# wip
def part2():
    fresh_ids = {}
    for range in ranges:
        start, end = range[0], range[1]
        print(start, end)


if __name__ == "__main__":
    #part1()
    part2()
