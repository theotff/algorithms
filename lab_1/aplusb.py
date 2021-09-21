with open("aplusb.in", 'r') as filein:
    data = filein.read()
with open("aplusb.out", 'w') as fileout:
    numbers = data.split(sep = ' ')
    fileout.write(str(int(numbers[0]) + int(numbers[1])))