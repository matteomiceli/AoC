import sys
sys.path.append("..")

from helpers.helpers import get_data

def one():
  data = get_data('./data.txt')
  elves = data['chunked']

  elf_calories = []

  for i in range(len(elves)):
    elf = elves[i]
    total_calories = 0
    for j in range(len(elf)):
      item = elf[j]
      if item != '':
        total_calories += int(item)

    elf_calories.append(total_calories)
    elf_calories.sort()
    elf_calories.reverse()

  print(max(elf_calories))
  print(elf_calories[0] + elf_calories[1] + elf_calories[2])

one()