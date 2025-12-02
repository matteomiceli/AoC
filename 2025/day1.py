from utils import get_lines
import math

def main():
    rotations = get_lines(1)
    dial = 50
    dial_zero = 0

    for r in rotations:
        direction = r[0]
        amount = int(r[1:])
        for i in range(amount):
            if direction == "R":
                dial += 1
            else:
                dial -= 1
            if dial % 100 == 0:
                dial_zero += 1

        dial = dial % 100

    print(dial_zero, dial)


if __name__ == "__main__":
    main()
