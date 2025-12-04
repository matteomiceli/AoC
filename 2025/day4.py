from utils import get_lines


def check_neighbors(rows: list[list[str]], y:int, x:int) -> list:
    y_limit = len(rows) - 1
    x_limit = len(rows[0]) - 1
    # [y, x]
    neighbors = [[-1, -1], [-1, 0], [-1, 1], [0, -1], [0, 1], [1, -1], [1, 0], [1, 1]]
    neighbor_rolls = []
    for neighbor in neighbors:
        ymod, xmod = neighbor[0], neighbor[1]
        if y + ymod < 0 or y + ymod > y_limit:
            continue
        if x + xmod < 0 or x + xmod > x_limit:
            continue

        if rows[y+ymod][x+xmod] == "@":
            neighbor_rolls.append([y + ymod, x + xmod])

    return neighbor_rolls
    

rows = [list(row) for row in get_lines(day=4)]
def part1():
    accessible_rolls = 0
    for y in range(len(rows)):
        for x in range(len(rows[y])):
            if rows[y][x] != "@": 
                continue
            adjacent_rolls = check_neighbors(rows, y, x)
            if len(adjacent_rolls) < 4:
                accessible_rolls += 1

    print(">>>", accessible_rolls)

def part2():
    def remove_and_traverse(accessible_rolls: list[list[int]], rm = 0):
        removed = rm
        for roll in accessible_rolls:
            y, x = roll[0], roll[1]
            rows[y][x] = "."
            removed += 1
            accessible_rolls = []
        for y in range(len(rows)):
            for x in range(len(rows[y])):
                if rows[y][x] != "@": 
                    continue
                adjacent_rolls = check_neighbors(rows, y, x)
                if len(adjacent_rolls) < 4:
                    accessible_rolls.append([y, x])
        if accessible_rolls:
            return remove_and_traverse(accessible_rolls, rm=removed)
        return removed
        
    removed = remove_and_traverse(accessible_rolls=[])
    print(">>>", removed)


if __name__ == "__main__":
    #part1()
    part2()
