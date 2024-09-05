def next_position(values, instructions):

    for instruction in instructions:

        if len(instruction) < 4:
            return instruction

        letter = instruction[0]
        comparison_op = instruction[1]
        number, destination = instruction[2:].split(":")

        if comparison_op == "<":
            if values[letter] < int(number):
                return destination
        else:
            if values[letter] > int(number):
                return destination


def part1(instructions, input_vals):

    instructions_dict = {}
    for instruction in instructions:
        key, rest = instruction.split("{")
        rest = rest[:-1].split(",")
        instructions_dict[key] = rest

    total = 0
    input_dict = {}
    for val in input_vals:
        val = val[1:-1].split(",")
        for x in val:
            key, value = x.split("=")
            input_dict[key] = int(value)

        position = "in"
        while position != "A" and position != "R":
            position = next_position(input_dict, instructions_dict[position])

        if position == "A":
            for value in input_dict.values():
                total += value

    return total


with open("input.txt") as file:
    lines = file.read().split("\n")

instructions = lines[: lines.index("")]
input_vals = lines[lines.index("") + 1 :]

part1_total = part1(instructions, input_vals)
print("part 1 total is: ", part1_total)
