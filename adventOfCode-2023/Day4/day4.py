def part1(lines):
    total = 0

    for line in lines:
        winningNums = (
            line.split(":")[1].split("|")[0].strip().replace("  ", " ").split(" ")
        )
        playingNums = line.split("|")[1].strip().replace("  ", " ").split(" ")

        score = 0
        for num in playingNums:
            if num in winningNums:
                if score == 0:
                    score = 1
                else:
                    score += score

        total += score

    print("total score:", total)


def part2(lines):

    cards = [1 for _ in range(len(lines))]

    for i, line in enumerate(lines):
        winningNums = (
            line.split(":")[1].split("|")[0].strip().replace("  ", " ").split(" ")
        )
        playingNums = line.split("|")[1].strip().replace("  ", " ").split(" ")

        score = 0
        for num in playingNums:
            if num in winningNums:
                score += 1

        for x in range(score):
            cards[i + x + 1] += cards[i]

    print("total score =", sum(cards))


with open("input.txt") as file:
    lines = file.read().split("\n")

part1(lines)
part2(lines)
