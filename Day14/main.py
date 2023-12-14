from enum import Enum


class Direction(Enum):
    UP = (-1, 0)
    DOWN = (1, 0)
    LEFT = (0, -1)
    RIGHT = (0, 1)


def parseData(data):
    data = data.lstrip().rstrip().split('\n')
    data = [list(i) for i in data]
    return data


CACHE = {}


def tilt(data):
    global CACHE
    key = (str(data))
    if key in CACHE:
        return CACHE[key]

    dir = Direction.UP

    for i in range(len(data)):
        for j in range(len(data[i])):
            if data[i][j] == 'O':
                ii, jj = i, j
                while True:
                    ti, tj = ii + dir.value[0], jj + dir.value[1]
                    if ti < 0 or ti >= len(data) or tj < 0 or tj >= len(data[i]):
                        break
                    if data[ti][tj] != '.':
                        break
                    ii, jj = ti, tj
                if ii == i and jj == j:
                    continue

                data[i][j] = '.'
                data[ii][jj] = 'O'

    CACHE[key] = data
    return data


def calc(data):
    sum = 0
    for i, row in enumerate(data):
        for j in row:
            if j == 'O':
                sum += len(data) - i
    return sum


def rotate(cur):
    rotate = []
    width = len(cur)
    for j in range(len(cur[0])):
        col = ['.']*width
        for i in range(len(cur)):
            col[len(cur)-i-1] = cur[i][j]
        rotate.append(col)
    return rotate


def pp(data):
    for i in data:
        print(''.join(i), end='')
    print()


def part2(data, length=1000000000):
    memory = []

    def p2Cycle(data):
        data = tilt(data)
        data = rotate(data)
        data = tilt(data)
        data = rotate(data)
        data = tilt(data)
        data = rotate(data)
        data = tilt(data)
        data = rotate(data)
        # pp(data)

        return data

    for i in range(length):
        # pp(data)
        data = p2Cycle(data)
        if data in memory:
            # print(f'Found period at {i}')
            break
        memory.append(data)

    cycle_start = memory.index(data)
    cycle_length = i - cycle_start
    idx = (length - cycle_start) % cycle_length + cycle_start - 1
    data = memory[idx]
    sum = calc(data)
    return sum


def test():
    testdata = '''O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....'''
    data = parseData(testdata)
    sum = calc(tilt(data))
    print(f'Test Part 1: {sum}  Expects: 136')
    data = parseData(testdata)
    sum = part2(data)
    print(f'Test Part 2: {sum}, Expects: 64')


def main():
    data = open("input.txt", "r").read()
    data = parseData(data)
    data = tilt(data)
    sum = calc(data)
    print(f'Part 1: {sum}')
    sum = part2(data)
    print(f'Part 2: {sum}')


if __name__ == "__main__":
    test()
    main()
