
def parseData(data):
    return data.lstrip().rstrip().split('\n\n')

def getReflectionScore(tile, allow_smudge: bool = False):
    tile = tile.split('\n')
    score = 0

    # print('\n'.join(tile))

    for i in range(1, len(tile)):
        difference = 0
        for j in range(0, min(i, len(tile) - i)):
            for k in range(len(tile[i])):
                if tile[i - 1 - j][k] != tile[i + j][k]:
                    difference += 1
                    
        if allow_smudge and difference == 1 or not allow_smudge and difference == 0:
            # print(f"Horizontal reflection found at: {str(i-1)}-{str(i)} (difference: {difference})")
            return i * 100

    for i in range(1, len(tile[0])):
        difference = 0
        for j in range(0, min(i, len(tile[0]) - i)):
            for k in range(len(tile)):
                if tile[k][i - 1 - j] != tile[k][i + j]:
                    difference += 1

        if allow_smudge and difference == 1 or not allow_smudge and difference == 0:
            # print(f"Vertical reflection found at: {str(i-1)}-{str(i)}")
            return i
    
    return score
        


def test():
    testdata = '''#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#'''
    data = parseData(testdata)
    sum = 0
    for i in data:
        sum += getReflectionScore(i)
    print(f'Test Part 1: {sum}')

    sum = 0
    for i in data:
        sum += getReflectionScore(i, True)
    print(f'Test Part 2: {sum}')

def main():
    data = open("input.txt", "r").read()
    data = parseData(data)
    sum = 0
    for i in data:
        sum += getReflectionScore(i)
    print(f'Part 1: {sum}')

    sum = 0
    for i in data:
        sum += getReflectionScore(i, True)
    print(f'Part 2: {sum}')


if __name__ == "__main__":
    test()

    main()