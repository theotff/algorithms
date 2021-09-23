import random

fin = open('race.in', 'r')
fout = open('race.out', 'w')

data = fin.read().strip().split('\n')
people = dict()

for i in data[1:]:
    info = i.split()
    if info[0] in people:
        people[info[0]].append(info[1])
    else:
        people[info[0]] = [info[1]]

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

countries = quicksort(list(people.keys()))

string = ''
for j in countries:
    if string == '':
        string += ('=== ' + j + ' ===\n')
    else:
        string += ('\n=== ' + j + ' ===\n')
    string += '\n'.join(people[j])

print(string, file = fout)