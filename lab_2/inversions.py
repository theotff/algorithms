fin = open('inversions.in', 'r')
fout = open('inversions.out', 'w')

numbers = [int(i) for i in fin.read().split('\n')[1].split()]

def mergesort(array):
    if len(array) < 2:
        return array, 0
    else:
        half = len(array) // 2
        left, l_inversions = mergesort(array[:half])
        right, r_inversions = mergesort(array[half:])
        result = []
        inversions = 0 + l_inversions + r_inversions

        l_index, r_index = 0, 0

        while l_index < len(left) and r_index < len(right):
            if left[l_index] <= right[r_index]:
                result.append(left[l_index])
                l_index += 1
            else:
                result.append(right[r_index])
                r_index += 1
                inversions += (len(left) - l_index)

        result += left[l_index:]
        result += right[r_index:]

    return result, inversions

print(mergesort(numbers)[1], file = fout)