import random

fin = open('sort.in', 'r')
fout = open('sort.out', 'w')

data = fin.read().strip().split('\n')
numbers_raw = data[1].split()

numbers = [int(i) for i in numbers_raw if i.strip('-').isnumeric()]

def quicksort(array: list) -> list:
    if len(array) < 2:
        return array
    else:
        index = random.randint(0, len(array) - 1)
        pivot = array[index]
        array.pop(index)
        less = [i for i in array if i <= pivot]
        greater = [i for i in array if i > pivot]

        result = quicksort(less) + [pivot] + quicksort(greater)
        return result

r = quicksort(numbers)
print(' '.join([str(i) for i in r]), file = fout)