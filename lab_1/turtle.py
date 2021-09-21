with open('turtle.in', 'r') as filein:
    data = filein.read()
lines = data.split('\n')
h, w = [int(i) for i in lines[0].strip().split(' ')]
matrix = [[int(i) for i in line.strip().split(' ') if i.isdigit()] for line in lines[1:]]

for i in range(h - 1, -1, -1):
    for j in range(w):
        if i == h - 1:
            if j != 0:
                matrix[i][j] = matrix[i][j] + matrix[i][j - 1]
        else:
            if j == 0:
                matrix[i][j] = matrix[i][j] + matrix[i + 1][j]
            else:
                matrix[i][j] = matrix[i][j] + max(matrix[i][j - 1], matrix[i + 1][j])

with open('turtle.out', 'w') as fileout:
    fileout.write(str(matrix[0][w - 1]))