from utils import get_lines

def part1():
    banks = get_lines(day=3)
    big_nums = []
    for bank in banks:
        biggest_num = 0
        for l in range(len(bank)):
            for r in range(l, len(bank)):
                if l == r:
                    continue
                num = int(bank[l]+bank[r])
                if num > biggest_num:
                    biggest_num = num

        big_nums.append(biggest_num)

    print(">>>", sum(big_nums))


def part2():
    banks = get_lines(day=3)
    big_nums = []
    num_batts = 12
    for bank in banks:
        batt_pos = 0
        big_num = []
        for i in range(num_batts):
            end = (len(bank) - (num_batts - i)) + 1
            check_range = bank[batt_pos:end]
            biggest_num = max(check_range)
            big_num.append(biggest_num)
            batt_pos = bank.index(biggest_num, batt_pos, end) + 1

        big_nums.append(int("".join(big_num)))
    
    print(">>>", sum(big_nums))



if __name__ == "__main__":
    #part1()
    part2()
