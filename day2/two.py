import sys
sys.path.append("..")

from helpers.helpers import get_data

data = get_data('./data.txt')

# 1. A - rock ----- X 
# 2. B - paper ---- Y
# 3. C - scissors - Z

# 0 loss
# 3 draw
# 6 win

# -- part 2
# X lose
# Y draw
# Z win

win_map = {
  "X": {
    "C": "win",
    "B": "lose"
  },
  "Y": {
    "A": "win",
    "C": "lose"
  },
  "Z": {
    "B": "win", 
    "A": "lose"
  }
}

#p2 -- what to throw to get the right outcome -- xyz on left column maps to lose,draw,win
missing_map = {
  "A": {
    "Z": "Y",
    "Y": "X",
    "X": "Z"
  },
  "B": {
    "Z": "Z",
    "Y": "Y",
    "X": "X"
  },
  "C": {
    "Z": "X", 
    "Y": "Z",
    "X": "Y"
  }
}

points = 0

def get_throw_value(you: str):
  match you:
    case "X": 
      return 1
    case "Y": 
      return 2
    case "Z": 
      return 3

def check_winner(pair: list):
  opponent = pair[0]
  you = pair[1]

  if opponent not in win_map[you]:
    return 3
  
  match win_map[you][opponent]:
    case "win":
      return 6
    case "lose":
      return 0

def winner_part2(outcome: str):
  match outcome:
    case "X":
      return 0
    case "Y": 
      return 3
    case "Z": 
      return 6

def outcome_based_throw(pair):
  opponent = pair[0]
  outcome = pair[1]


  return missing_map[opponent][outcome]

# for pair in data['pairs']:
#   points += get_throw_value(pair[1])
#   points += check_winner(pair)

#p2
for pair in data['pairs']:
  you = outcome_based_throw(pair)
  points += get_throw_value(you)
  points += winner_part2(pair[1])
  
print(points)