from helpers.get_data import get_data

def one():
  data = get_data('./day1/data.txt')
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

  print(elf_calories[0] + elf_calories[1] + elf_calories[2])

  

def find_biggest(map: dict):
  highest = 0
  idx = 0

  for k, value in map.items():
    if value > highest:
      highest = value 
      idx = k
  
  return [highest, idx ]


    
