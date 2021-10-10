with open("aplusbb.in", 'r') as filein:
    data = filein.read()
with open("aplusbb.out", 'w') as fileout:
    numbers = data.split(sep = ' ')
    fileout.write(str(int(numbers[0]) + pow(int(numbers[1]), 2)))