# The engine schematic (your puzzle input) consists of a visual representation of the engine. T
# here are lots of numbers and symbols you don't really understand, but apparently any number adjacent to a symbol, even diagonally, is a "part number" and should be included in your sum. 
# (Periods (.) do not count as a symbol.)

# Here is an example engine schematic:

# 467..114..
# ...*......
# ..35..633.
# ......#...
# 617*......
# .....+.58.
# ..592.....
# ......755.
# ...$.*....
# .664.598..
# In this schematic, two numbers are not part numbers because they are not adjacent to a symbol: 114 (top right) and 58 (middle right). Every other number is adjacent to a symbol and so is a part number; their sum is 4361.

#method checks to see if a number has a adjacent gear
def has_adjacent_gear(x, y):

    directions = [(0,1), (0,-1), (1,0), (-1,0), (1,-1), (1,1), (-1,1), (-1,-1)]
    for dir in directions:
        new_x = x + dir[0]
        new_y = y + dir[1]
        if (new_x,new_y) in gears:
            return True

    return False

symbols = ['*', '#', '+', '$', '/', '@', '=', '%', '-', '&']
grid = []

with open("input.txt") as file:
    for line in file:
        grid.append(line.strip())
gears = {}
numbers = []

#putting every symbol into a dictionary with its position and every number into a array
for i in range(len(grid)):
    for j in range(len(grid[i])):
        if grid[i][j] in symbols:
            gears[(i, j)] = grid[i][j]
        elif grid[i][j].isdigit():
            numbers.append((i,j))

gear_nums = []

for num in numbers:
    y,x = num
    if has_adjacent_gear(y,x):
        gear_nums.append((y,x))
        continue

    #check to see if there are adjacent numbers and whether they are next to a gear
    if (
        (((y, x+1) in numbers) and has_adjacent_gear(y,x+1))  or 
        ((((y, x+2) in numbers) and ((y, x+1) in numbers)) and has_adjacent_gear(y, x+2)) or
        (((y, x-1) in numbers) and has_adjacent_gear(y,x-1)) or 
        ((((y, x-2) in numbers) and ((y, x-1) in numbers)) and has_adjacent_gear(y, x-2)) 
        ):
        gear_nums.append((y,x))
        continue
    else:
        continue

entire_lonely_nums = []

skip_counter = 0

for pos in gear_nums:
    if skip_counter < 0:
        skip_counter = 0
    if skip_counter > 0:
        skip_counter -= 1
        continue
    y,x = pos
    if((y, x+2) in gear_nums) and ((y,x+1) in gear_nums):
        current_number = grid[y][x] +  grid[y][x+1] + grid[y][x+2]
        entire_lonely_nums.append(current_number)
        skip_counter += 2
    elif((y, x+1) in gear_nums):
        current_number = grid[y][x] + grid[y][x+1]
        entire_lonely_nums.append(current_number)
        skip_counter +=1
    else:
        entire_lonely_nums.append(grid[y][x])
        skip_counter -= 1

total = 0
for num in entire_lonely_nums:
    total += int(num)

print('Part 1: ', total)

## Part 2

def get_total_num(input):
    y, x  = input
    
    total_num = []

    if((y, x+2) in gear_nums) and ((y,x+1) in gear_nums):
        current_number = grid[y][x] +  grid[y][x+1] + grid[y][x+2]
        total_num.append(current_number)
    elif((y, x-2) in gear_nums) and ((y,x-1) in gear_nums):
        current_number = grid[y][x-2] +  grid[y][x-1] + grid[y][x]
        total_num.append(current_number)
    elif((y, x-1) in gear_nums) and ((y, x+1) in gear_nums):
        current_number = grid[y][x-1] + grid [y][x] + grid[y][x+1]
        total_num.append(current_number)
    elif((y, x+1) in gear_nums):
        current_number = grid[y][x] + grid[y][x+1]
        total_num.append(current_number)
    elif((y, x-1) in gear_nums):
        current_number = grid[y][x-1] + grid[y][x]
        total_num.append(current_number)
    else:
        total_num.append(grid[y][x])

    return total_num[0]

gears = {}
numbers = []

#putting every symbol into a dictionary with its position and every number into a array
for i in range(len(grid)):
    for j in range(len(grid[i])):
        if grid[i][j] == '*':
            gears[(i, j)] = grid[i][j]
        elif grid[i][j].isdigit():
            numbers.append((i,j))


directions = [(0,1), (0,-1), (1,0), (-1,0), (1,-1), (1,1), (-1,1), (-1,-1)]

part_two_total = 0

for gear in gears:
    adjacent_nums = set()
    for dir in directions:
        y = gear[0]
        x = gear[1]
        new_x = x + dir[1]
        new_y = y + dir[0]
        if (new_y,new_x) in numbers:
            adjacent_nums.add(get_total_num((new_y, new_x)))

    if len(adjacent_nums) == 2:
        adjacent_nums = list(adjacent_nums)
        part_two_total += (int(adjacent_nums[0]) * int(adjacent_nums[1]))

print('part2:', part_two_total)



