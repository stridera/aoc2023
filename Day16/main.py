from enum import Enum
from dataclasses import dataclass

class Direction(Enum):
    UP = (0, -1)
    DOWN = (0, 1)
    LEFT = (-1, 0)
    RIGHT = (1, 0)

    def __hash__(self):
        return hash((self.value[0], self.value[1]))
    
    def __repr__(self):
        if self == Direction.UP:
            return '^'
        elif self == Direction.DOWN:
            return 'v'
        elif self == Direction.LEFT:
            return '<'
        elif self == Direction.RIGHT:
            return '>'

@dataclass
class Vector:
    x: int
    y: int
    direction: Direction    

    def __hash__(self):
        return hash((self.x, self.y, self.direction))
    
    def str(self):
        return f'{self.x}, {self.y} {self.direction}'
    
    def __eq__(self, other):
        return self.x == other.x and self.y == other.y and self.direction == other.direction
    
    def __repr__(self):
        return f'{self.x}, {self.y} {self.direction}'
    
    def move(self):
        dx, dy = self.direction.value
        self.x += self.direction.value[0]
        self.y += self.direction.value[1]
        return self
    


def parseData(data):
    data = data.lstrip().rstrip().split('\n')
    return data

def shootLaser(grid, start):
    energized = {}
    paths = [start]

    while len(paths) > 0:
        # print(f"\n{paths}")
        path = paths.pop(0)
        x, y = path.x, path.y
        # print(f'{x}, {y} {path.direction} ', end='')
        if x < 0 or y < 0 or y >= len(grid) or x >= len(grid[0]):
            # print(f"Out of bounds")
            continue

        key = (x, y, path.direction)
        if key in energized:
            continue
        energized[key] = True

        # print(f'{grid[y][x]} ', end='')

        if grid[y][x] == '/':
            if path.direction == Direction.RIGHT:
                path.direction = Direction.UP
            elif path.direction == Direction.LEFT:
                path.direction = Direction.DOWN
            elif path.direction == Direction.UP:
                path.direction = Direction.RIGHT
            elif path.direction == Direction.DOWN:
                path.direction = Direction.LEFT
            path.move()
            paths.append(path)
        elif grid[y][x] == '\\':
            if path.direction == Direction.RIGHT:
                path.direction = Direction.DOWN
            elif path.direction == Direction.LEFT:
                path.direction = Direction.UP
            elif path.direction == Direction.UP:
                path.direction = Direction.LEFT
            elif path.direction == Direction.DOWN:
                path.direction = Direction.RIGHT
            path.move()
            paths.append(path)
        elif grid[y][x] == '-':
            newpath = Vector(path.x, path.y, Direction.RIGHT)
            newpath.move()
            paths.append(newpath)
            newpath = Vector(path.x, path.y, Direction.LEFT)
            newpath.move()
            paths.append(newpath)
        elif grid[y][x] == '|':
            newpath = Vector(path.x, path.y, Direction.UP)
            newpath.move()
            paths.append(newpath)
            newpath = Vector(path.x, path.y, Direction.DOWN)
            newpath.move()
            paths.append(newpath)
        elif grid[y][x] == '.':
            path.move()
            paths.append(path)
        else:
            continue

        # print(f'{paths}')
    seen = {}
    for x in energized:
        seen[(x[0], x[1])] = True

    # print()
    # for y in range(len(grid)):
    #     for x in range(len(grid[0])):
    #         print(grid[y][x], end='')
    #     print("  ", end='')
    #     for x in range(len(grid[0])):
    #         print('x' if (x, y) not in seen else grid[y][x], end='')
    #     print()

                
    # print(seen)

    return len(seen)

           

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

def part2(data):
    m = 0
    for y in range(len(data)):
        if y == 0:
            for x in range(len(data[0])):
                m = max(m, shootLaser(data, Vector(x, y, Direction.DOWN)))
        elif y == len(data) - 1:
            for x in range(len(data[0])):
                m = max(m, shootLaser(data, Vector(x, y, Direction.UP)))
        m = max(m, shootLaser(data, Vector(0, y, Direction.RIGHT)))
        m = max(m, shootLaser(data, Vector(len(data[0])-1, y, Direction.LEFT)))
    return m

def test():
    testdata = r'''
.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....'''
    data = parseData(testdata)
    energized = shootLaser(data, Vector(0, 0, Direction.RIGHT))
    print(f'Test Part 1: {energized}  Expects: 46')

    m = part2(data)
    print(f'Test Part 2: {m}, Expects: 51')


def main():
    data = open("input.txt", "r").read()
    data = parseData(data)
    energized = shootLaser(data, Vector(0, 0, Direction.RIGHT))
    print(f'Part 1: {energized}')
    m = part2(data)
    print(f'Part 2: {m}')


if __name__ == "__main__":
    test()
    main()
