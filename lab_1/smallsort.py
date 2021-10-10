f = open('smallsort.in', 'r')
fout = open('smallsort.out', 'w')

n = int(f.readline().strip())
numbers = [int(i) for i in f.readline().split()]
for i in range(1, n):
    for j in range(i, 0, -1):
        if numbers[j] < numbers[j - 1]:
            numbers[j], numbers[j - 1] = numbers[j - 1], numbers[j]
        else:
            break

print(' '.join([str(i) for i in numbers]), file = fout)