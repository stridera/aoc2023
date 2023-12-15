
def parseData(data):
    data = data.lstrip().rstrip().split(',')
    return data

def hash(data):
    sum = 0
    for i in data:
        sum += ord(i)
        sum *= 17
        sum %= 256
    return sum


def part1(data):
    sum = 0

    for i in data:
        sum += hash(i)

    return sum

def part2(data):
    boxes = [{} for i in range(256)]
    for i in data:
        if '-' in i:
            i = i[:-1]
            box = hash(i)
            if i in boxes[box]:
                del boxes[box][i]
        else:
            box, val = i.split('=')[0], i.split('=')[1]
            boxes[hash(box)][box] = val

    power = 0
    for i, box in enumerate(boxes):
        if len(box) > 0:
            for j, key in enumerate(box):
                power += (i+1) * (j+1) * int(box[key])
            
    return power

def test():
    testdata = '''rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7'''
    data = parseData(testdata)
    sum = part1(data)
    print(f'Test Part 1: {sum}  Expects: 1320')
    sum = part2(data)
    print(f'Test Part 2: {sum}, Expects: 145')


def main():
    data = open("input.txt", "r").read()
    data = parseData(data)
    sum = part1(data)
    print(f'Part 1: {sum}')
    sum = part2(data)
    print(f'Part 2: {sum}')


if __name__ == "__main__":
    test()
    main()
