import sys
sys.path.append("..")

from helpers.get_data import get_data

data = get_data('./data.txt')

# 49-49,10-50
def get_ranges(line):
  pairs = line.split(',')
  organized_pairs = []

  for pair in pairs:
    organized_pairs.append(pair.split('-'))

  return organized_pairs

def is_first_inside_second(first, second):
  is_inside = True 
  if not int(first[0]) >= int(second[0]):
    is_inside = False 
  if not int(first[1]) <= int(second[1]):
    is_inside = False 

  return is_inside


def has_range_overlap(first, second):
  has_overlap = False 
  min = int(second[0])
  max = int(second[1])
  
  if int(first[0]) >= min and int(first[0]) <= max:
    has_overlap = True
  if int(first[1]) >= min and int(first[1]) <= max:
    has_overlap = True
  
  if min >= int(first[0]) and min <= int(first[1]):
    has_overlap = True
  if max >= int(first[0]) and max <= int(first[1]):
    has_overlap = True
  return has_overlap



is_inside = {}
has_overlap = {}
for i, line in enumerate(data['list']):
  # [['6','6'], ['4', '6']]
  ranged_pair = get_ranges(line)

  first_is_inside =  is_first_inside_second(ranged_pair[0], ranged_pair[1])
  if first_is_inside:
    is_inside[i] = True
  second_is_inside =  is_first_inside_second(ranged_pair[1], ranged_pair[0])
  if second_is_inside:
    is_inside[i] = True
  
  overlap = has_range_overlap(ranged_pair[0], ranged_pair[1])
  if overlap:
    has_overlap[i] = True
  
print(len(has_overlap))
