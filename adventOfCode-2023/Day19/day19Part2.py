import time, copy

start_time = time.time()

with open("input.txt") as f:
    content_list = f.readlines()

content_list = [x.strip() for x in content_list]

def new_group_lowside(old_group, variable, limit):
    
    group = copy.deepcopy(old_group)

    if group[variable][1] >= limit:
        group[variable][1] = limit - 1

    return group

def new_group_highside(old_group, variable, limit):
    
    group = copy.deepcopy(old_group)

    if group[variable][0] <= limit:
        group[variable][0] = limit + 1

    return group

def flow(workflow, enter_group):

    total = 0

    comparisons = workflow.split(",")

    for comp in comparisons:
        if "<" in comp or ">" in comp:
            result = comp.split(":")[1]
            ltgt = comp.split(":")[0]

            if "<" in comp:
                var = ltgt.split("<")[0]
                tar = int(ltgt.split("<")[1])

                if enter_group[var][0] < tar:

                    new_group = new_group_lowside(enter_group, var, tar)
                    enter_group = new_group_highside(enter_group, var, tar - 1)
                    
                    if result == "A":
                        to_add = 1
                        for key in new_group:
                            to_add *= (new_group[key][1] - new_group[key][0] + 1)
                        total += to_add
                    elif result != "R":
                        total += flow(all_flows[result], new_group)

            if ">" in comp:
                var = ltgt.split(">")[0]
                tar = int(ltgt.split(">")[1])

                if enter_group[var][1] > tar:

                    new_group = new_group_highside(enter_group, var, tar)
                    enter_group = new_group_lowside(enter_group, var, tar + 1)
                    
                    if result == "A":
                        to_add = 1
                        for key in new_group:
                            to_add *= (new_group[key][1] - new_group[key][0] + 1)
                        total += to_add
                    elif result != "R":
                        total += flow(all_flows[result], new_group)

        else:
            if comp == "A":
                to_add = 1
                for key in enter_group:
                    to_add *= (enter_group[key][1] - enter_group[key][0] + 1)
                total += to_add
            elif comp != "R":
                total += flow(all_flows[comp], enter_group)

    return total

all_flows = {}

current_line = 0
while content_list[current_line] != "":

    all_flows[content_list[current_line].split("{")[0]] = content_list[current_line].split("{")[1][:-1]

    current_line += 1

group = {"x":[1, 4000], "m":[1, 4000], "a": [1, 4000], "s": [1, 4000]}

total = flow(all_flows["in"], group)

print(total)

print(4000*4000*4000*4000)

print(time.time() - start_time)