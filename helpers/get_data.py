
def get_data(path):
  f = open(path, 'r')
  stringified = f.read()
  array = stringified.split('\n')
  double_newline = stringified.split('\n\n')
  chunked = []
  
  for chunk in double_newline:
    newlist = chunk.split('\n')
    chunked.append(newlist)
  
  return {
    'raw': stringified,
    'array': array,
    'chunked': chunked,
    'double_newline': double_newline
  }



  