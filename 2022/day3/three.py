import sys
sys.path.append("..")

from helpers.helpers import get_data

def halve_sack_contents(contents):
  first = contents[0 : (len(contents)//2)]
  second = contents[(len(contents)//2) : len(contents)] 
  return [first, second]

def get_item_priority(char):
  if not char:
    return 0
  unicode = ord(char)
  if unicode < 97: # uppercase
    return unicode - 38 # adjust to return 27 - 52 for A-Z
  
  return unicode - 96 # adjust to return 1 - 26 for a - z

data = get_data('./data.txt')

priority_sum = 0
for items in data['list']:
  found = {}
  alike = {}

  [first, second] = halve_sack_contents(items)
  for char in first:
    found[char] = True
  
  for char in second:
    if char in found:
      alike[char] = True

  for char in alike:
    priority_sum += get_item_priority(char)


# part 2
def get_found(contents, found, i):
  for item in contents:
    if item in found:
      found[item][i] = True
    else: 
      found[item] = {i: True}

def get_bage(common):
  for item in common:
    if len(item[1]) == 3:
     return item[0]


badge_sum = 0
for i, items in enumerate(data['list']):
  found = {}
  if i % 3 == 0: 
    get_found(data['list'][i], found, i)
    get_found(data['list'][i + 1], found, i+1)
    get_found(data['list'][i + 2], found, i+2)
  
  badge_sum += get_item_priority(get_bage(found.items()))

print(badge_sum)