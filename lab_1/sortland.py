fin = open('sortland.in', 'r')
fout = open('sortland.out', 'w')

people = int(fin.readline().strip())
numbers = [float(i) for i in fin.readline().split()]
sortlist = numbers.copy()

for i in range(1, people):
    for j in range(i, 0, -1):
        if sortlist[j] <= sortlist[j - 1]:
            sortlist[j], sortlist[j - 1] = sortlist[j - 1], sortlist[j]

print(str(numbers.index(sortlist[0]) + 1) + ' ' + str(numbers.index(sortlist[people//2]) + 1) + ' ' + str(numbers.index(sortlist[people - 1]) + 1), file = fout)