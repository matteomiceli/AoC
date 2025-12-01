def get_data(day: int):
    with open(f"./data/day{day}.txt") as f:
        return f.readlines()
