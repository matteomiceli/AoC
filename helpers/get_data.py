
def get_data(path):
  f = open(path, 'r')
  stringified = f.read()
  list = stringified.split('\n')
  double_newline = stringified.split('\n\n')
  chunked = []
  pairs = []
  
  # build chunked data (split by two \n\n)
  for chunk in double_newline:
    newlist = chunk.split('\n')
    chunked.append(newlist)
  
  # build list of pairs
  for item in list:
    pairs.append(item.split())
  
  return {
    'raw': stringified,
    'list': list,
    'chunked': chunked,
    'pairs': pairs
  }



  