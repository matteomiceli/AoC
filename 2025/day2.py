from utils import get_data

def parse_ranges() -> list[tuple[int, int]]:
    ranges = []
    data = get_data(2)
    range_sets = data.split(",")
    for rset in range_sets:
        [start, end] = rset.split("-")
        ranges.append((int(start), int(end)))

    return ranges

def part1():
    ranges = parse_ranges()
    invalid_ids = []
    for r in ranges:
        for i in range(r[0], r[1]+1):
            str_i = str(i)
            if len(str_i) % 2 == 0:
                # split number in half
                first = str_i[0:len(str_i)//2]
                last= str_i[(len(str_i)//2):]
                if first == last:
                    invalid_ids.append(i)

    print("part 1:", sum(invalid_ids))

def num_repeats(num_str: str) -> bool:
    biggest_chunk = len(num_str) // 2
    for chunk_size in range(1, biggest_chunk + 1):
        if len(num_str) % chunk_size != 0:
            continue
        chunks = [num_str[i: i+chunk_size] for i in range(0, len(num_str), chunk_size)]
        chunk_map = {}
        for chunk in chunks:
            chunk_map[chunk] = True

        if len(chunk_map) == 1:
            return True

    return False

def part2():
    ranges = parse_ranges()
    invalid_ids = []
    for r in ranges:
        for i in range(r[0], r[1]+1):
            str_i = str(i)
            has_repeating_pattern = num_repeats(str_i)
            if has_repeating_pattern:
                invalid_ids.append(i)


    print("part 2:", sum(invalid_ids))


def main():
    #part1()
    part2()

if __name__ == "__main__":
    main()
