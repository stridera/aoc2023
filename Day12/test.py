def parse_data(data):
    maints = []
    for line in data.split('\n'):
        if line:
            parts = line.split(' ')
            maint = {'springs': list(parts[0]), 'groups': []}
            for s in parts[1].split(','):
                if s:
                    maint['groups'].append(int(s))
            maints.append(maint)
    return maints


def create_perms(springs):
    # print("Create Perms: ", ''.join(springs))

    perms = []
    for spring in springs:
        new_perms = []
        if spring == '?':
            if not perms:
                new_perms.extend(['.', '#'])
            else:
                for p in perms:
                    new_perms.extend([p + '#', p + '.'])
        else:
            if not perms:
                new_perms.append(spring)
            else:
                for p in perms:
                    new_perms.append(p + spring)
        perms = new_perms

    seen = set()
    new_perms = []
    for perm in perms:
        if perm not in seen:
            seen.add(perm)
            new_perms.append(perm)

    return new_perms


def is_valid(springs, groups):
    while True:
        if not springs and not groups:
            return True
        elif not springs:
            return False

        if springs[0] == '.':
            springs = springs[1:]
        elif springs[0] == '#':
            if not groups:
                return False
            group = groups[0]
            if len(springs) < group:
                return False
            for i in range(group):
                if springs[i] != '#':
                    return False
            springs = springs[group:]
            groups = groups[1:]
            if len(springs) >= 1 and springs[0] == '#':
                return False


def part1(maints):
    sum_valid = 0
    for maint in maints:
        perms = create_perms(maint['springs'])
        for perm in perms:
            if is_valid(perm, maint['groups']):
                sum_valid += 1
    return sum_valid


def part2(maints) -> int:
    sum_valid = 0
    for maint in maints:
        springs = "?".join(maint['springs'] * 5)
        groups = [maint['groups']] * 5
        perms = create_perms(springs)
        for perm in perms:
            if is_valid(perm, groups):
                sum_valid += 1


def test():
    test_data = """???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1"""

    maints = parse_data(test_data)
    p1 = part1(maints)
    print(f"TEST Part 1: {p1} Expects: 21 Pass? {p1 == 21}")
    p2 = part2(maints)
    print(f"TEST Part 2: {p2} Expects: 525152 Pass? {p2 == 525152}")


if __name__ == "__main__":
    test()
    with open("input.txt") as f:
        data = f.read()
    maints = parse_data(data)
    p1 = part1(maints)
    print(f"Part 1: {p1}")
